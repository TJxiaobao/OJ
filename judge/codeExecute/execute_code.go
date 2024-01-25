package codeExecute

import (
	"bytes"
	"encoding/json"
	"github.com/TJxiaobao/OJ/constant"
	"github.com/TJxiaobao/OJ/judge/models"
	"log"
	"net/http"
)

func CodeExecute(language, code, input string) {
	// 构建参数
	request := models.Request{
		Language: language,
		Code:     code,
		Input:    input,
	}
	jsonBody, err := json.Marshal(request)
	if err != nil {
		log.Fatal("json error :", err)
		return
	}

	// 创建请求
	req, err := http.NewRequest("POST", constant.SandBoXUrl, bytes.NewBuffer(jsonBody))

	// 设置请求头
	req.Header.Set("Content-Type", "application/json")

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Failed to send request:", err)
	}
	defer resp.Body.Close()

	// 处理响应
	if resp.StatusCode != http.StatusOK {
		log.Fatal("Request failed with status:", resp.StatusCode)
		return
	}

	var responseBody models.Response
	err = json.NewDecoder(resp.Body).Decode(&responseBody)
	if err != nil {
		log.Fatal("Failed to decode response body:", err)
	}

	// 可以访问解析后的响应数据
	output := responseBody.Output
	message := responseBody.Message
	status := responseBody.Status
	judgeMessage := responseBody.JudgeInfo.Message
	judgeMemory := responseBody.JudgeInfo.Memory
	judgeTime := responseBody.JudgeInfo.Time

	// 业务处理
	print(output)
	print(message)
	print(status)
	print(judgeTime)
	print(judgeMemory)
	print(judgeMessage)
}
