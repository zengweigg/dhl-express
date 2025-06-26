package dhl_express

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
	"github.com/zengweigg/dhl-express/config"
)

type service struct {
	config     *config.Config // Config
	logger     Logger         // Logger
	httpClient *resty.Client  // HTTP client
}

type services struct {
	Base  baseService
	Track trackService
}

// GetSign MD5加密
func GetSign(text string) string {
	// 创建一个 md5 哈希对象
	hash := md5.New()
	// 将输入字符串写入哈希对象
	hash.Write([]byte(text))
	// 计算哈希值
	hashBytes := hash.Sum(nil)
	// 将哈希值转换为十六进制字符串
	hashStr := hex.EncodeToString(hashBytes)
	return hashStr
}

func GenerateID() string {
	return uuid.New().String()
}

// Base64加密字符串
func Base64Encode(input string) string {
	return base64.StdEncoding.EncodeToString([]byte(input))
}

func GetTestToken(username, password string) (token string) {
	if username == "" || password == "" {
		return ""
	}
	// 使用测试账号和密码用BASE64进行加密
	bs64 := Base64Encode(username + ":" + password)
	token = "Basic " + bs64
	return token
}
