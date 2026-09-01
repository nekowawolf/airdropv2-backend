package bot

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/nekowawolf/airdropv2/config"
	"github.com/nekowawolf/airdropv2/features/media"
	"github.com/nekowawolf/airdropv2/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"

	tele "gopkg.in/telebot.v3"
)

var userUploadState = make(map[int64]bool)

type ProjectEndpoint struct {
	ID    string
	Label string
	Path  string
	Icon  string
}

var projectEndpoints = []ProjectEndpoint{
	{"airdrop", "Airdrop", "/airdrops", "🪂"},
	{"cryptocommunity", "Crypto Community", "/cryptocommunity", "🪙"},
	{"aitools", "AI Tools", "/aitools", "🤖"},
	{"web3tools", "Web3 Tools", "/web3tools", "🌐"},
	{"githubrepo", "Github Repo", "/githubrepo", "🐙"},
}

func checkAuth(c tele.Context) bool {
	chatIDStr := os.Getenv("TELEGRAM_CHAT_ID")
	if chatIDStr == "" {
		return false
	}
	expectedID, _ := strconv.ParseInt(chatIDStr, 10, 64)
	return c.Chat().ID == expectedID
}

var apiBaseURL string

func getBaseURL() (string, error) {
	if apiBaseURL != "" {
		return apiBaseURL, nil
	}
	baseURL := os.Getenv("API_BASE_URL")
	if baseURL == "" {
		return "", fmt.Errorf("API_BASE_URL is not set in .env")
	}
	apiBaseURL = baseURL
	return apiBaseURL, nil
}

func handleSpeedTest(c tele.Context) error {
	c.Respond()
	if !checkAuth(c) {
		return c.Send("❌ Unauthorized access.")
	}

	c.Send("⏳ Testing API speed...")
	baseURL, err := getBaseURL()
	if err != nil {
		return c.Send(fmt.Sprintf("❌ Configuration Error: %v", err))
	}
	endpoints := []string{"/airdrops", "/profilelink", "/postslink", "/cryptocommunity", "/price", "/portfolio", "/aitools", "/web3tools", "/githubrepo"}

	results := "⚡ API Speed Test Results\n\n"
	allNormal := true

	for _, ep := range endpoints {
		start := time.Now()
		resp, err := http.Get(baseURL + ep)
		if err != nil {
			allNormal = false
			results += fmt.Sprintf("🔗 %s : Error (%v)\n", ep, err)
			continue
		}
		resp.Body.Close()
		duration := time.Since(start).Milliseconds()
		
		if resp.StatusCode == 200 {
			results += fmt.Sprintf("🔗 %s : %d ms\n", ep, duration)
		} else {
			allNormal = false
			results += fmt.Sprintf("🔗 %s : %d ms (Status %d)\n", ep, duration, resp.StatusCode)
		}
	}

	results += "\nStatus: "
	if allNormal {
		results += "All endpoints responded normally."
	} else {
		results += "Some endpoints experienced issues."
	}

	return c.Send(results)
}

func handleCheckMissingImages(c tele.Context) error {
	c.Respond()
	if !checkAuth(c) {
		return c.Send("❌ Unauthorized access.")
	}

	menu := &tele.ReplyMarkup{}
	var rows []tele.Row

	for _, ep := range projectEndpoints {
		if ep.ID == "githubrepo" { continue }
		btn := menu.Data(ep.Icon+" "+ep.Label, "exe_img_chk", ep.ID)
		if len(rows) > 0 && len(rows[len(rows)-1]) < 2 {
			rows[len(rows)-1] = append(rows[len(rows)-1], btn)
		} else {
			rows = append(rows, menu.Row(btn))
		}
	}

	btnCancel := menu.Data("❌ Cancel", "btn_cancel_chk")
	rows = append(rows, menu.Row(btnCancel))
	
	menu.Inline(rows...)
	return c.Edit("🔍 *Check Missing Images*\nPilih project mana yang mau di-scan:", menu, tele.ModeMarkdown)
}

func handleExecuteImageCheck(c tele.Context) error {
	c.Respond()
	if !checkAuth(c) {
		return c.Send("❌ Unauthorized access.")
	}

	projectID := c.Callback().Data
	var selectedEp *ProjectEndpoint
	for _, ep := range projectEndpoints {
		if ep.ID == projectID {
			selectedEp = &ep
			break
		}
	}

	if selectedEp == nil {
		return c.Edit("❌ Invalid project selected.")
	}

	c.Send(fmt.Sprintf("🔍 Checking missing images for *%s*...", selectedEp.Label), tele.ModeMarkdown)

	baseURL, err := getBaseURL()
	if err != nil {
		return c.Send(fmt.Sprintf("❌ Configuration Error: %v", err))
	}
	urlStr := baseURL + selectedEp.Path

	go func() {
		totalMissing := 0
		var detailsBlocks []string

		resp, err := http.Get(urlStr)
		if err != nil {
			log.Printf("Error fetching %s: %v\n", urlStr, err)
			c.Send(fmt.Sprintf("❌ Error fetching %s: %v", selectedEp.Label, err))
			return
		}

		var data struct {
			Data []struct {
				Name          string `json:"name"`
				ImageURL      string `json:"image_url"`
			} `json:"data"`
		}

		err = json.NewDecoder(resp.Body).Decode(&data)
		resp.Body.Close()
		if err != nil {
			log.Printf("Error decoding JSON from %s: %v\n", urlStr, err)
			c.Send(fmt.Sprintf("❌ Error decoding JSON from %s", selectedEp.Label))
			return
		}

		var blockDetails string
		for _, item := range data.Data {
			imgURL := item.ImageURL
			if imgURL == "" { continue }

			req, err := http.NewRequest("GET", imgURL, nil)
			if err != nil { continue }
			req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36")
			
			client := &http.Client{Timeout: 10 * time.Second}
			imgResp, err := client.Do(req)

			if err != nil || imgResp.StatusCode != 200 {
				totalMissing++
				blockDetails += fmt.Sprintf("- Name: \"%s\"\n", item.Name)
			}
			if imgResp != nil {
				imgResp.Body.Close()
			}
		}

		if blockDetails != "" {
			detailsBlocks = append(detailsBlocks, fmt.Sprintf("[%s]\n%s", selectedEp.Label, blockDetails))
		}

		msg := fmt.Sprintf("🔍 Image Check Complete for *%s*!\n\nTotal Broken Images: %d\n", selectedEp.Label, totalMissing)
		if len(detailsBlocks) > 0 {
			msg += "\nDetails:\n" + strings.Join(detailsBlocks, "\n")
		} else {
			msg += "\nDetails: All images are safe!"
		}

		if len(msg) > 4000 { msg = msg[:4000] + "\n... (truncated)" }
		c.Send(msg)
	}()

	return nil
}

func handleCheckInvalidLink(c tele.Context) error {
	c.Respond()
	if !checkAuth(c) {
		return c.Send("❌ Unauthorized access.")
	}

	menu := &tele.ReplyMarkup{}
	var rows []tele.Row

	for _, ep := range projectEndpoints {
		btn := menu.Data(ep.Icon+" "+ep.Label, "exe_lnk_chk", ep.ID)
		if len(rows) > 0 && len(rows[len(rows)-1]) < 2 {
			rows[len(rows)-1] = append(rows[len(rows)-1], btn)
		} else {
			rows = append(rows, menu.Row(btn))
		}
	}

	btnCancel := menu.Data("❌ Cancel", "btn_cancel_chk")
	rows = append(rows, menu.Row(btnCancel))
	
	menu.Inline(rows...)
	return c.Edit("🔗 *Check Invalid Links*\nPilih project mana yang mau di-scan:", menu, tele.ModeMarkdown)
}

func handleExecuteLinkCheck(c tele.Context) error {
	c.Respond()
	if !checkAuth(c) {
		return c.Send("❌ Unauthorized access.")
	}

	projectID := c.Callback().Data
	var selectedEp *ProjectEndpoint
	for _, ep := range projectEndpoints {
		if ep.ID == projectID {
			selectedEp = &ep
			break
		}
	}

	if selectedEp == nil {
		return c.Edit("❌ Invalid project selected.")
	}

	c.Send(fmt.Sprintf("🔗 Checking invalid links for *%s*...", selectedEp.Label), tele.ModeMarkdown)

	baseURL, err := getBaseURL()
	if err != nil {
		return c.Send(fmt.Sprintf("❌ Configuration Error: %v", err))
	}
	urlStr := baseURL + selectedEp.Path

	go func() {
		totalInvalid := 0
		var detailsBlocks []string

		resp, err := http.Get(urlStr)
		if err != nil {
			log.Printf("Error fetching %s: %v\n", urlStr, err)
			c.Send(fmt.Sprintf("❌ Error fetching %s: %v", selectedEp.Label, err))
			return
		}

		var data struct {
			Data []struct {
				Name         string `json:"name"`
				Link         string `json:"link"`
				Website      string `json:"website"`
				RepoURL      string `json:"repo_url"`
				Twitter      string `json:"twitter"`
				Discord      string `json:"discord"`
				Telegram     string `json:"telegram"`
				ClaimURL     string `json:"claim_url"`
				GuideURL     string `json:"guide_url"`
				VideoURL     string `json:"video_url"`
				Instagram    string `json:"instagram"`
				Youtube      string `json:"youtube"`
			} `json:"data"`
		}

		err = json.NewDecoder(resp.Body).Decode(&data)
		resp.Body.Close()
		if err != nil {
			log.Printf("Error decoding JSON from %s: %v\n", urlStr, err)
			c.Send(fmt.Sprintf("❌ Error decoding JSON from %s", selectedEp.Label))
			return
		}

		var blockDetails string
		for _, item := range data.Data {
			primaryLink := item.Link
			if primaryLink == "" { primaryLink = item.Website }

			linksToCheck := []struct {
				Name string
				URL  string
			}{
				{"Primary", primaryLink},
				{"Repo", item.RepoURL},
				{"Twitter", item.Twitter},
				{"Discord", item.Discord},
				{"Telegram", item.Telegram},
				{"Claim", item.ClaimURL},
				{"Guide", item.GuideURL},
				{"Video", item.VideoURL},
				{"Instagram", item.Instagram},
				{"Youtube", item.Youtube},
			}

			for _, l := range linksToCheck {
				link := l.URL
				if link == "" { continue }

				parsedURL, err := url.ParseRequestURI(link)
				if err != nil || parsedURL.Scheme == "" || parsedURL.Host == "" {
					totalInvalid++
					blockDetails += fmt.Sprintf("- Name: \"%s\" (Invalid %s format: %s)\n", item.Name, l.Name, link)
					continue
				}

				req, err := http.NewRequest("GET", link, nil)
				if err != nil { continue }
				req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36")
				
				client := &http.Client{Timeout: 10 * time.Second}
				linkResp, err := client.Do(req)

				if err != nil || linkResp.StatusCode >= 400 {
					totalInvalid++
					blockDetails += fmt.Sprintf("- Name: \"%s\" (%s Link: %s)\n", item.Name, l.Name, link)
				}
				if linkResp != nil {
					linkResp.Body.Close()
				}
			}
		}

		if blockDetails != "" {
			detailsBlocks = append(detailsBlocks, fmt.Sprintf("[%s]\n%s", selectedEp.Label, blockDetails))
		}

		msg := fmt.Sprintf("🔗 Invalid Link Check Complete for *%s*!\n\nTotal Invalid Links: %d\n", selectedEp.Label, totalInvalid)
		if len(detailsBlocks) > 0 {
			msg += "\nDetails:\n" + strings.Join(detailsBlocks, "\n")
		} else {
			msg += "\nDetails: All links are valid!"
		}

		if len(msg) > 4000 { msg = msg[:4000] + "\n... (truncated)" }
		c.Send(msg)
	}()

	return nil
}

func handleCDNInit(c tele.Context) error {
	if !checkAuth(c) {
		return nil
	}
	userUploadState[c.Chat().ID] = true
	return c.Send("🖼️ Cloudflare R2 CDN Upload\n\nPlease send me the photo you want to upload. (It will be uploaded to Cloudflare R2 and auto-converted to WebP).")
}

func handlePhotoUpload(c tele.Context) error {
	if !checkAuth(c) {
		return nil
	}

	if !userUploadState[c.Chat().ID] {
		return nil
	}

	userUploadState[c.Chat().ID] = false
	c.Send("⏳ Processing and uploading photo to Cloudflare R2...")

	photo := c.Message().Photo
	if photo == nil {
		return c.Send("❌ No photo found in the message.")
	}

	file, err := TelegramBot.FileByID(photo.FileID)
	if err != nil {
		return c.Send(fmt.Sprintf("❌ Failed to get photo: %v", err))
	}

	rc, err := TelegramBot.File(&file)
	if err != nil {
		return c.Send(fmt.Sprintf("❌ Failed to download photo: %v", err))
	}
	defer rc.Close()

	buf, err := io.ReadAll(rc)
	if err != nil {
		return c.Send(fmt.Sprintf("❌ Failed to read photo data: %v", err))
	}

	// Convert to WebP
	filename := fmt.Sprintf("%d_upload.webp", time.Now().Unix())
	uploadBytes := buf
	uploadContentType := "image/jpeg"

	webpBytes, err := utils.ConvertToWebP(buf)
	if err != nil {
		fmt.Printf("WebP conversion failed in bot, uploading original: %v\n", err)
		filename = fmt.Sprintf("%d_upload.jpg", time.Now().Unix())
	} else {
		uploadBytes = webpBytes
		uploadContentType = "image/webp"
	}

	r2Key := utils.GenerateR2Key("image", filename)

	finalURL, err := utils.UploadToR2(uploadBytes, r2Key, uploadContentType)
	if err != nil {
		return c.Send(fmt.Sprintf("❌ R2 upload failed: %v", err))
	}

	mediaObj := media.Media{
		ID:          primitive.NewObjectID(),
		Filename:    filename,
		URL:         finalURL,
		Size:        int64(len(uploadBytes)),
		ContentType: uploadContentType,
		MediaType:   "image",
		R2Key:       r2Key,
		CreatedAt:   time.Now(),
	}

	_, err = config.Database.Collection("media").InsertOne(context.Background(), mediaObj)
	if err != nil {
		return c.Send(fmt.Sprintf("❌ Upload successful to R2, but failed to save to database: %v", err))
	}

	return c.Send(fmt.Sprintf("✅ Upload Successful!\n\nURL: %s", finalURL))
}