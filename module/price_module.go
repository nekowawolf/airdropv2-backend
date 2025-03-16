package module

import (
	"encoding/json"
	"fmt"
	"net/http"
	"io"

	"github.com/nekowawolf/airdropv2/models"
)

func GetPrice(url string) (*models.CryptoData, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var data []models.CryptoData
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	if len(data) > 0 {
		return &data[0], nil
	}

	return nil, fmt.Errorf("data not found")
}
