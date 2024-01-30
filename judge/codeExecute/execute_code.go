package codeExecute

import (
	"bytes"
	"encoding/json"
	"github.com/TJxiaobao/OJ/constant"
	"github.com/TJxiaobao/OJ/judge/models"
	"log"
	"net/http"
)

func CodeExecute(language, code, input string) *models.Response {
	// 构建参数
	request := models.Request{
		Language: language,
		Code:     code,
		Input:    input,
	}
	jsonBody, err := json.Marshal(request)
	if err != nil {
		log.Fatal("json error :", err)
		return nil
	}

	// 创建请求
	req, err := http.NewRequest("POST", constant.SandBoXUrl, bytes.NewBuffer(jsonBody))

	// 添加鉴权请求头
	req.Header.Set(constant.AUTH_REQUEST_HEADER, constant.AUTH_REQUEST_SECRET)

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
		return nil
	}

	// 判断是否存在 body
	var responseBody models.Response
	if resp.Body != nil {
		defer resp.Body.Close() // 确保在函数返回前关闭响应体
		err := json.NewDecoder(resp.Body).Decode(&responseBody)
		if err != nil {
			log.Println("Failed to decode response body:", err)
			// 其他错误处理逻辑
			return nil
		} else {
			return nil
		}
	}

	// 可以访问解析后的响应数据
	output := responseBody.Output
	message := responseBody.Message
	status := responseBody.Status
	judgeMessage := responseBody.JudgeInfo.Message
	judgeMemory := responseBody.JudgeInfo.Memory
	judgeTime := responseBody.JudgeInfo.Time
	judgeInfo := &models.JudgeInfo{
		Message: judgeMessage,
		Memory:  judgeMemory,
		Time:    judgeTime,
	}
	res := &models.Response{
		Message:   message,
		Output:    output,
		Status:    status,
		JudgeInfo: *judgeInfo,
	}
	return res
}
