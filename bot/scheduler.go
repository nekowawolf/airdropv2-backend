package bot

import (
	"context"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/nekowawolf/airdropv2/config"
	"github.com/nekowawolf/airdropv2/features/message"
	"github.com/nekowawolf/airdropv2/features/github"
	"github.com/robfig/cron/v3"
	"go.mongodb.org/mongo-driver/bson"
	tele "gopkg.in/telebot.v3"
)

var cronScheduler *cron.Cron

func InitScheduler() {
	loc, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		log.Printf("Failed to load timezone Asia/Jakarta, falling back to UTC: %v", err)
		cronScheduler = cron.New()
	} else {
		cronScheduler = cron.New(cron.WithLocation(loc))
	}

	_, err = cronScheduler.AddFunc("0 3 * * 1", func() {
		log.Println("Running scheduled backup...")
		SendBackupArchive()
	})
	if err != nil {
		log.Fatalf("Failed to add cron job: %v", err)
	}

	_, err = cronScheduler.AddFunc("0 20 * * *", func() {
		log.Println("Running journal alert check (20:00)...")
		CheckAndSendJournalAlert(1)
	})
	if err != nil {
		log.Printf("Failed to add cron job (20:00): %v", err)
	}

	_, err = cronScheduler.AddFunc("0 23 * * *", func() {
		log.Println("Running journal alert check (23:00)...")
		CheckAndSendJournalAlert(2)
	})
	if err != nil {
		log.Printf("Failed to add cron job (23:00): %v", err)
	}

	_, err = cronScheduler.AddFunc("0 0,12 * * *", func() {
		log.Println("Running Github Repo Stats Sync (00:00 & 12:00)...")
		github.SyncAllGithubRepoStats()
	})
	if err != nil {
		log.Printf("Failed to add cron job for Github sync: %v", err)
	}

	cronScheduler.Start()
	log.Println("Scheduler initialized")
}

func GetNextBackupTime() string {
	if cronScheduler == nil {
		return "Unknown"
	}
	
	entries := cronScheduler.Entries()
	if len(entries) > 0 {
		return entries[0].Next.Format("Monday 15:04 WIB")
	}
	return "None"
}

func CheckAndSendJournalAlert(textIndex int) {
	if TelegramBot == nil {
		return
	}
	
	chatIDStr := os.Getenv("TELEGRAM_CHAT_ID")
	if chatIDStr == "" { return }
	chatID, err := strconv.ParseInt(chatIDStr, 10, 64)
	if err != nil { return }
	
	wib := time.FixedZone("WIB", 7*3600)
	now := time.Now().In(wib)
	startDate := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, wib)
	endDate := startDate.AddDate(0, 0, 1)
	
	filter := bson.M{
		"type": "journal",
		"created_at": bson.M{"$gte": startDate, "$lt": endDate},
	}
	
	count, err := config.Database.Collection("notes").CountDocuments(context.Background(), filter)
	if err != nil {
		log.Printf("Failed to count journal notes: %v", err)
		return
	}
	
	if count > 0 {
		return
	}
	
	var msgConfig message.Message
	err = config.Database.Collection("messages").FindOne(context.Background(), bson.M{}).Decode(&msgConfig)
	if err != nil {
		log.Printf("Failed to fetch message config for alert: %v", err)
	}
	
	var alertText string
	if textIndex == 1 { alertText = msgConfig.Text1 }
	if textIndex == 2 { alertText = msgConfig.Text2 }
	
	if alertText == "" {
		return
	}
	
	chat := &tele.Chat{ID: chatID}
	_, err = TelegramBot.Send(chat, alertText)
	if err != nil {
		log.Printf("Failed to send journal alert: %v", err)
	}
}