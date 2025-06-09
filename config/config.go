package config

type Config struct {
	APIKey   string
	ApiToken string
	Debug    bool // 是否启用调试模式
	Sandbox  bool // 是否为沙箱环境
	Timeout  int  // HTTP 超时设定（单位：秒）
}

func LoadConfig() *Config {
	return &Config{
		APIKey:   "changshatiaCN",
		ApiToken: "M!0cY!0dF$0oA!0o",
		Debug:    true,
		Sandbox:  true,
		Timeout:  360,
	}
}
