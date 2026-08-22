package utils

import (
	"bytes"
	"context"
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"os"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/mayahiro/go-webp"
)

var r2Client *s3.Client
var r2Bucket string
var r2PublicURL string

func InitR2Client() {
	accountID := os.Getenv("R2_ACCOUNT_ID")
	accessKeyID := os.Getenv("R2_ACCESS_KEY_ID")
	secretAccessKey := os.Getenv("R2_SECRET_ACCESS_KEY")
	r2Bucket = os.Getenv("R2_BUCKET_NAME")
	endpoint := os.Getenv("R2_ENDPOINT")
	r2PublicURL = os.Getenv("R2_PUBLIC_URL")

	if accountID == "" || accessKeyID == "" || secretAccessKey == "" || r2Bucket == "" || endpoint == "" || r2PublicURL == "" {
		fmt.Println("WARNING: R2 environment variables not fully set")
		return
	}

	r2PublicURL = strings.TrimRight(r2PublicURL, "/")

	cfg, err := config.LoadDefaultConfig(context.Background(),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(accessKeyID, secretAccessKey, "")),
		config.WithRegion("auto"),
	)
	if err != nil {
		fmt.Printf("Failed to load R2 config: %v\n", err)
		return
	}

	r2Client = s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.BaseEndpoint = aws.String(endpoint)
	})

	fmt.Println("R2 client initialized successfully")
}

func UploadToR2(fileBytes []byte, key string, contentType string) (string, error) {
	if r2Client == nil {
		return "", fmt.Errorf("R2 client not initialized")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	_, err := r2Client.PutObject(ctx, &s3.PutObjectInput{
		Bucket:      aws.String(r2Bucket),
		Key:         aws.String(key),
		Body:        bytes.NewReader(fileBytes),
		ContentType: aws.String(contentType),
	})
	if err != nil {
		return "", fmt.Errorf("failed to upload to R2: %v", err)
	}

	url := fmt.Sprintf("%s/%s", r2PublicURL, key)
	return url, nil
}

func DeleteFromR2(key string) error {
	if r2Client == nil {
		return fmt.Errorf("R2 client not initialized")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	_, err := r2Client.DeleteObject(ctx, &s3.DeleteObjectInput{
		Bucket: aws.String(r2Bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		return fmt.Errorf("failed to delete from R2: %v", err)
	}

	return nil
}

func ConvertToWebP(imageBytes []byte) ([]byte, error) {
	img, _, err := image.Decode(bytes.NewReader(imageBytes))
	if err != nil {
		return nil, fmt.Errorf("failed to decode image: %v", err)
	}

	var buf bytes.Buffer
	if err := webp.Encode(&buf, img, &webp.Options{
		Compression: webp.CompressionLossy,
		Quality:     80,
	}); err != nil {
		return nil, fmt.Errorf("failed to encode to webp: %v", err)
	}

	return buf.Bytes(), nil
}

func IsImageContentType(contentType string) bool {
	return strings.HasPrefix(contentType, "image/")
}

func IsVideoContentType(contentType string) bool {
	return strings.HasPrefix(contentType, "video/")
}

func GetMediaType(contentType string) string {
	if IsImageContentType(contentType) {
		return "image"
	}
	if IsVideoContentType(contentType) {
		return "video"
	}
	return "unknown"
}

func GenerateR2Key(mediaType, filename string) string {
	now := time.Now()
	return fmt.Sprintf("%s/%d/%d_%s", mediaType, now.Year(), now.Unix(), filename)
}

func StreamToBytes(stream io.Reader) ([]byte, error) {
	return io.ReadAll(stream)
}