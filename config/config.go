package config

type Config struct {
	Username     string // 账号
	Password     string // 密码
	CustomerCode string // APP 客户编码
	Platform     string // 平台名称 英文
	Version      string // 版本
	Debug        bool   // 是否启用调试模式
	Sandbox      bool   // 是否为沙箱环境
	Timeout      int    // HTTP 超时设定（单位：秒）
}

func LoadConfig() *Config {
	return &Config{
		Username:     "changshatiaCN",
		Password:     "M!0cY!0dF$0oA!0o",
		CustomerCode: "123456789",
		Platform:     "nebula",
		Version:      "1.0",
		Debug:        true,
		Sandbox:      true,
		Timeout:      360,
	}
}
