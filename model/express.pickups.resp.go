package model

type PickupsResp struct {
	DispatchConfirmationNumbers []string `json:"dispatchConfirmationNumbers"` // 发货确认号码列表，用于识别预定的取货
	ReadyByTime                 string   `json:"readyByTime"`                 // 按时间准备  12:00
	NextPickupDate              string   `json:"nextPickupDate"`              // 下次取货日期 2020-06-01
	Warnings                    []string `json:"warnings"`                    // 警告 创建了 Pickup 但出现了问题
}
