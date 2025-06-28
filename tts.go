package tts

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

var client = &http.Client{}

const VOLCENGINE_HOST_URL = "https://translate.volcengine.com/web/tts/v1"

type GenerateResponse struct {
	Audio struct {
		Duration int    `json:"duration"`
		Data     string `json:"data"`
	} `json:"audio"`
	BaseResp struct {
		StatusCode    int    `json:"status_code"`
		StatusMessage string `json:"status_message"`
	} `json:"base_resp"`
}

// Generate 生成的 WAV 格式的 BASE64 字符串
func Generate(speaker Speaker, text string) (string, error) {
	data := map[string]string{
		"language": "zh",
		"speaker":  string(speaker),
		"text":     text,
	}
	payload, err := json.Marshal(data)
	if err != nil {
		return "", fmt.Errorf("payload json marshal error: %v", err)
	}

	req, _ := http.NewRequest(http.MethodPost, VOLCENGINE_HOST_URL, bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/137.0.0.0 Safari/537.36")

	response, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("generate request error: %v", err)
	}
	defer response.Body.Close()

	var respData GenerateResponse
	err = json.NewDecoder(response.Body).Decode(&respData)
	if err != nil {
		return "", fmt.Errorf("response json unmarshal error: %v", err)
	}

	if respData.BaseResp.StatusCode != 0 {
		return "", fmt.Errorf("response status code error[%d]: %v", respData.BaseResp.StatusCode, respData.BaseResp.StatusMessage)
	}

	return respData.Audio.Data, nil
}
