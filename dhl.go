package dhl_express

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/zengweigg/dhl-express/config"
	"time"
)

const (
	Version   = "1.0.0"
	userAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/129.0.0.0 Safari/537.36"
)

type DhlExpClient struct {
	config     *config.Config
	httpClient *resty.Client
	logger     Logger   // Logger
	Services   services // API Services
}

func NewDHLService(cfg config.Config) *DhlExpClient {
	loc, _ := time.LoadLocation("Asia/Shanghai")                     // 加载中国标准时间的时区信息
	currentTime := time.Now().In(loc)                                // 获取当前时间并转换为中国时区
	formattedTime := currentTime.Format("2006-01-02T15:04:05+08:00") // 按照指定格式输出，其中 +08:00 表示 UTC+8 时区的时间偏移
	DhlClient := &DhlExpClient{
		config: &cfg,
		logger: createLogger(),
	}
	// OnBeforeRequest：设置请求发送前的钩子函数，允许在请求发送之前对请求进行修改或添加逻辑。
	// OnAfterResponse：设置响应接收后的钩子函数，允许在接收到响应后处理响应数据或执行其他逻辑。
	// SetRetryCount：设置请求失败时的最大重试次数。
	// SetRetryWaitTime：设置每次重试之间的等待时间（最小等待时间）。
	// SetRetryMaxWaitTime：设置每次重试之间的最大等待时间，实际等待时间会在最小和最大等待时间之间随机选取。
	// AddRetryCondition：添加自定义的重试条件，当满足该条件时触发重试机制。
	// | 状态码 | 说明 |
	// |--------|------|
	// | **200** | 请求成功。"成功"的具体含义取决于使用的 HTTP 方法：<br>• **GET**：资源已成功获取，并在响应体中传输。<br>• **POST / PATCH**：操作结果对应的资源在响应体中返回。 |
	// | **201** | 请求成功，且服务器创建了一个新资源作为结果。通常用于 `POST` 请求后返回新资源的创建状态。 |
	// | **400** | 由于客户端错误（如请求语法错误、无效的消息结构或欺骗性请求路由），服务器无法处理该请求。 |
	// | **401** | 请求需要用户身份验证。虽然状态码字面意思是“未授权”，但实际语义是“未认证”。客户端需提供凭证才能获得响应。 |
	// | **403** | 客户端没有访问所请求资源的权限。与 401 不同的是，服务器已经知道客户端的身份，但拒绝授予访问权限。 |
	// | **404** | 服务器找不到请求的资源。在浏览器中表示 URL 无法识别；在 API 中可能表示端点有效但资源不存在。服务器也可能用此状态码代替 403 来隐藏资源是否存在。这是最广为人知的状态码之一，常见于网页访问失败时。 |
	// | **422** | 请求格式正确，但由于语义错误（如字段值不符合业务规则）无法执行。例如参数校验失败。 |
	// | **500** | 服务器遇到未知错误，无法处理请求。通常是服务器内部异常导致。 |
	httpClient := resty.
		New().
		SetDebug(DhlClient.config.Debug).
		SetHeaders(map[string]string{
			"Content-Type":                     "application/json",
			"Accept":                           "application/json",
			"User-Agent":                       userAgent,
			"Message-Reference":                GenerateID(),
			"Message-Reference-Date":           formattedTime,
			"Plugin-Name":                      cfg.Platform,
			"Plugin-Version":                   cfg.Version,
			"Shipping-System-Platform-Name":    cfg.Platform,
			"Shipping-System-Platform-Version": cfg.Version,
			"Webstore-Platform-Name":           cfg.Platform,
			"Webstore-Platform-Version":        cfg.Version,
		})
	if cfg.Sandbox {
		httpClient.SetBaseURL("https://express.api.dhl.com/mydhlapi/test")
	} else {
		httpClient.SetBaseURL("https://express.api.dhl.com/mydhlapi")
	}
	httpClient.
		SetTimeout(time.Duration(cfg.Timeout) * time.Second).
		OnBeforeRequest(func(client *resty.Client, request *resty.Request) error {
			// 获取token算法加密后
			request.SetHeaders(map[string]string{
				"Authorization": GetTestToken(cfg.Username, cfg.Password),
			})
			return nil
		}).
		SetRetryCount(1).
		SetRetryWaitTime(5 * time.Second).
		SetRetryMaxWaitTime(10 * time.Second).
		AddRetryCondition(func(r *resty.Response, err error) bool {
			text := r.Request.URL
			if r == nil {
				return false
			}
			if err != nil {
				text += fmt.Sprintf(", error: %s", err.Error())
				DhlClient.logger.Debugf("Retry request: %s", text)
				return true // 如果有错误则重试
			}
			// 检查响应状态码是否不是200
			// if r.StatusCode() != http.StatusOK && r.StatusCode() != http.StatusCreated {
			// 	text += fmt.Sprintf(", error: %d", r.StatusCode())
			// 	DhlClient.logger.Debugf("Retry request: %s", text)
			// 	return true
			// }
			return false
		})
	DhlClient.httpClient = httpClient
	xService := service{
		config:     &cfg,
		logger:     DhlClient.logger,
		httpClient: DhlClient.httpClient,
	}
	DhlClient.Services = services{
		Base:  (baseService)(xService),
		Track: (trackService)(xService),
	}
	return DhlClient
}
