package dhl_express

import (
	"fmt"
	"github.com/bytedance/sonic"
	"github.com/go-resty/resty/v2"
	"net/http"
	"time"
)

// 物流轨迹查询

type trackService service

type ExtraProperties struct {
	FlightNumber string `json:"FlightNumber"` // 航班号
}

type Checkpoint struct {
	TimeStamp            time.Time        `json:"time_stamp"`              // 操作时间
	TimeZone             string           `json:"time_zone"`               // 时区
	TrackingStatus       string           `json:"tracking_status"`         // 操作状态码
	Message              string           `json:"message"`                 // 轨迹描述
	Location             string           `json:"location"`                // 操作地点
	IsLastMileCheckpoint int              `json:"is_last_mile_checkpoint"` // 是否尾程服务商轨迹
	ExtraProperties      *ExtraProperties `json:"extraProperties"`         // 扩展字段
}

type TrackingStatusWaybill struct {
	Level1 string `json:"level1"` // 一级状态
	Level2 string `json:"level2"` // 二级状态
	Level3 string `json:"level3"` // 最后操作状态码
}

type Result struct {
	TrackingNumber           string                `json:"tracking_number"`             // 请求单号
	WaybillNumber            string                `json:"waybill_number"`              // 燕文单号
	ExchangeNumber           string                `json:"exchange_number"`             // 尾程单号
	Checkpoints              []Checkpoint          `json:"checkpoints"`                 // 轨迹明细
	TrackingStatus           string                `json:"tracking_status"`             // 最后操作状态码
	TrackingStatusWaybill    TrackingStatusWaybill `json:"tracking_status_waybill"`     // 包裹状态
	LastMileTrackingExpected bool                  `json:"last_mile_tracking_expected"` // 是否追踪派送环节轨迹
	OriginCountry            string                `json:"origin_country"`              // 始发国
	DestinationCountry       string                `json:"destination_country"`         // 目的国
}

type ElapsedMilliseconds struct {
	Total int `json:"total"` // 响应时长
}

type TrackingResponse struct {
	Code                int                 `json:"code"`
	Message             string              `json:"message"`
	Result              []Result            `json:"result"`
	RequestTime         time.Time           `json:"requestTime"`
	ElapsedMilliseconds ElapsedMilliseconds `json:"elapsedMilliseconds"`
}

// 物流轨迹查询
func (s trackService) Tracking(apikey, numbers string) (TrackingResponse, error) {
	respData := TrackingResponse{}
	url := "http://api.track.yw56.com.cn/api/tracking?nums=" + numbers
	// 创建一个新的 resty 客户端
	client := resty.New()
	client.SetTimeout(time.Duration(120) * time.Second).
		SetRetryCount(2).
		SetRetryWaitTime(5 * time.Second).
		SetRetryMaxWaitTime(10 * time.Second).
		AddRetryCondition(func(r *resty.Response, err error) bool {
			if r == nil {
				return false
			}
			if err != nil {
				return true // 如果有错误则重试
			}
			// 检查响应状态码是否不是200
			if r.StatusCode() != http.StatusOK {
				return true
			}
			type ResponseBody struct {
				Code int `json:"apiResultCode"`
			}
			// 解析响应体JSON
			var responseBody ResponseBody
			if err := sonic.Unmarshal(r.Body(), &responseBody); err != nil {
				return true // 如果解析错误则重试
			}
			// 检查apiResultCode字段是否不是0
			if responseBody.Code != 0 {
				return true
			}
			return false
		})
	// 请求数据
	resp, err := client.R().
		SetDebug(true).
		SetHeaders(map[string]string{
			"Content-Type":  "application/json",
			"Accept":        "application/json",
			"User-Agent":    userAgent,
			"Authorization": apikey,
		}).
		Get(url)
	fmt.Println(string(resp.Body()))
	if err != nil {
		return respData, err
	}

	return respData, nil
}
