package module

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"time"
	"net/url"
	"strings"

	"github.com/nekowawolf/airdropv2/models"
)

func UploadToGitHub(file multipart.File, fileHeader *multipart.FileHeader) (string, string, string, error) {
	token := os.Getenv("GITHUB_TOKEN")
	username := os.Getenv("GITHUB_USERNAME")
	repo := os.Getenv("GITHUB_REPO")
	baseDir := os.Getenv("GITHUB_UPLOAD_DIR")

	if token == "" || username == "" || repo == "" {
		return "", "", "", fmt.Errorf("GitHub environment variables not set")
	}

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		return "", "", "", err
	}

	encoded := base64.StdEncoding.EncodeToString(fileBytes)

	now := time.Now()
	folderPath := fmt.Sprintf("%s/%d", baseDir, now.Year())

	filename := fmt.Sprintf("%d_%s", time.Now().Unix(), fileHeader.Filename)
	fullPath := fmt.Sprintf("%s/%s", folderPath, filename)

	uploadURL := fmt.Sprintf(
		"https://api.github.com/repos/%s/%s/contents/%s",
		username,
		repo,
		fullPath,
	)

	body := models.GitHubUploadRequest{
		Message: "Upload image via API",
		Content: encoded,
	}

	jsonBody, _ := json.Marshal(body)

	req, err := http.NewRequest("PUT", uploadURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return "", "", "", err
	}

	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return "", "", "", err
	}
	defer res.Body.Close()

	if res.StatusCode >= 300 {
		responseText, _ := io.ReadAll(res.Body)
		return "", "", "", fmt.Errorf("GitHub upload failed: %s", responseText)
	}

	var ghResp models.GitHubUploadResponse
	json.NewDecoder(res.Body).Decode(&ghResp)

	parts := strings.Split(ghResp.Content.Path, "/")
	for i, p := range parts {
		parts[i] = url.PathEscape(p)
	}
	escapedPath := strings.Join(parts, "/")

	finalURL := fmt.Sprintf(
		"https://%s.github.io/%s/%s",
		username,
		repo,
		escapedPath,
	)

	return finalURL, ghResp.Content.Sha, ghResp.Content.Path, nil
}

func DeleteFromGitHub(path, sha string) error {
	token := os.Getenv("GITHUB_TOKEN")
	username := os.Getenv("GITHUB_USERNAME")
	repo := os.Getenv("GITHUB_REPO")

	if token == "" || username == "" || repo == "" {
		return fmt.Errorf("GitHub environment variables not set")
	}

	deleteURL := fmt.Sprintf(
		"https://api.github.com/repos/%s/%s/contents/%s",
		username,
		repo,
		path,
	)

	body := map[string]string{
		"message": "Delete image via API",
		"sha":     sha,
	}

	jsonBody, _ := json.Marshal(body)

	req, err := http.NewRequest("DELETE", deleteURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode >= 300 {
		return fmt.Errorf("GitHub deletion failed: %s", res.Status)
	}

	return nil
}