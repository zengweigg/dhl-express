package dhl_express

import (
	"fmt"
	"github.com/bytedance/sonic"
	"github.com/go-resty/resty/v2"
	"github.com/hiscaler/gox/bytex"
	"github.com/hiscaler/gox/filex"
	"os"
	"path"
	"time"
)

type Token struct {
	AccessToken     string `json:"access_token"`
	TokenType       string `json:"token_type"`
	ExpiresIn       int    `json:"expireIn"`
	ExpiresDatetime int64  `json:"expires_datetime"`
}

// Valid 判断是否有token且没有过期
func (t Token) Valid() bool {
	return t.AccessToken != "" && t.ExpiresDatetime > time.Now().Unix()
}

type FileToken struct {
	Key   string
	Path  string
	Token Token
}

var tokenFileCache = make(map[string]string)

type TokenWriterReader interface {
	Read(key string) (Token, error)
	Write(key string, token Token) (bool, error)
	newHttpClient(debug, sandbox bool) *resty.Client
	GetToken(apikey, apisecret string, debug, sandbox bool) (Token, error)
	GetValidAccessToken(apikey, apisecret string, debug, sandbox bool) (Token, error)
}

// tokenFilePath 生成并返回存储Token的文件路径。
// 参数 key 用于区分不同的Token文件，通常是与Token相关联的唯一标识。
// 返回的路径是操作系统临时目录下的一个文件路径，文件名为 "sf_token_<key>.json"。
func tokenFilePath(key string) string {
	if tokenPath, ok := tokenFileCache[key]; ok {
		return tokenPath
	}
	str := path.Join(os.TempDir(), fmt.Sprintf("sf_token_%s.json", key))
	tokenFileCache[key] = str
	return str
}

// Read 从文件中读取与指定 key 关联的 Token 信息。
// 参数 key 用于标识唯一的 Token 文件。
// 返回 Token 结构体以及可能的错误信息。
func (ft FileToken) Read(key string) (Token, error) {
	var token Token
	file := tokenFilePath(key)
	if !filex.Exists(file) {
		return token, os.ErrNotExist
	}
	b, err := os.ReadFile(file)
	if err != nil {
		return token, err
	}
	if err := sonic.Unmarshal(b, &token); err != nil {
		return token, err
	}
	return token, nil
}

// Write 将 Token 结构体序列化为 JSON 并写入与指定 key 关联的文件中。
// 参数 key 用于标识唯一的 Token 文件，token 是需要存储的 Token 数据。
// 返回一个布尔值表示写入是否成功，以及可能的错误信息。
func (ft FileToken) Write(key string, token Token) (bool, error) {
	// 将 Token 结构体序列化为 JSON 格式的字节数组
	b, err := sonic.Marshal(token)
	if err != nil {
		return false, err // 如果序列化失败，返回错误
	}

	// 将 JSON 数据写入文件，文件路径由 tokenFilePath(key) 生成，权限设置为 0777（可读可写可执行）
	err = os.WriteFile(tokenFilePath(key), b, 0775)
	if err != nil {
		return false, err // 如果写入文件失败，返回错误
	}

	// 写入成功，返回 true 和 nil 错误
	return true, nil
}

// GetValidAccessToken 获取有效的 AccessToken
func (ft FileToken) GetValidAccessToken(apikey, apisecret string, debug, sandbox bool) (Token, error) {
	// 尝试从本地读取 Token
	token, err := ft.Read(apikey)
	if err == nil && token.Valid() {
		return token, nil
	}

	// 如果本地 Token 无效或不存在，则从远程获取新的 Token
	// newToken, err := s.GetToken()
	newToken, err := ft.GetToken(apikey, apisecret, debug, sandbox)
	if err != nil {
		return Token{}, err
	}

	// 将新的 Token 写入本地文件
	_, err = ft.Write(apikey, newToken)
	if err != nil {
		return Token{}, err
	}

	return newToken, nil
}

func (ft FileToken) newHttpClient(debug, sandbox bool) *resty.Client {
	httpClient := resty.
		New().
		SetDebug(debug).
		SetHeaders(map[string]string{
			"Content-Type": "application/x-www-form-urlencoded",
			"Accept":       "application/json",
			"User-Agent":   userAgent,
		})
	if sandbox {
		httpClient.SetBaseURL("https://api-sandbox.dhlecs.com")
	} else {
		httpClient.SetBaseURL("https://api.dhlecs.com")
	}
	return httpClient
}

func (ft FileToken) GetToken(apikey, apisecret string, debug, sandbox bool) (ar Token, err error) {
	// 定义一个结构体用于解析API返回的JSON数据
	result := struct {
		AccessToken string `json:"access_token"`
		ExpiresIn   int    `json:"expires_in"`
		TokenType   string `json:"token_type"`
		ClientID    string `json:"client_id"`
	}{}

	// 创建一个新的HTTP客户端，并设置日志记录器
	// 设置返回结果解析到result结构体
	resp, err := ft.newHttpClient(debug, sandbox).
		R().
		SetResult(&result).
		SetFormData(map[string]string{
			"client_id":     apikey,
			"client_secret": apisecret,
			"grant_type":    "client_credentials",
		}).
		Post("/auth/v4/accesstoken") // 发送请求获取Token

	// 如果请求过程中发生错误，返回错误信息
	if err != nil {
		err = fmt.Errorf("%s: %s", resp.Status(), bytex.ToString(resp.Body()))
		return
	}
	// 如果请求成功
	if resp.StatusCode() == 200 {
		// 将API返回的Token数据赋值给ar
		if result.AccessToken != "" && result.ExpiresIn > 0 {
			ar.AccessToken = result.AccessToken
			ar.TokenType = result.TokenType
			// 计算Token的过期时间，设置为当前时间加上Token有效期的90%（剩余10%时间需要更换Token）
			ar.ExpiresDatetime = time.Now().Unix() + int64(result.ExpiresIn*9/10)
		}
	} else {
		// 如果请求不成功，返回错误信息
		err = fmt.Errorf("%s: %s", resp.Status(), bytex.ToString(resp.Body()))
	}
	return
}
