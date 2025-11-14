package models

type GitHubUploadRequest struct {
	Message string `json:"message"`
	Content string `json:"content"`
}

type GitHubUploadResponse struct {
	Content struct {
		Path string `json:"path"`
		Sha  string `json:"sha"`
	} `json:"content"`
}