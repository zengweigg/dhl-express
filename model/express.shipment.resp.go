package model

type ShipmentResp struct {
	ShipmentTrackingNumber string            `json:"shipmentTrackingNumber"` // 运单号
	TrackingUrl            string            `json:"trackingUrl"`            // 跟踪链接
	Packages               []PackageResp     `json:"packages"`               // 包裹信息
	Documents              []Document        `json:"documents"`              // 运单文件
	ShipmentDetails        []ShipmentDetail  `json:"shipmentDetails"`        // 运单详情
	EstimatedDeliveryDate  EstimatedDelivery `json:"estimatedDeliveryDate"`  // 预计送达时间
	OnDemandDeliveryURL    string            `json:"onDemandDeliveryURL"`    // 按需交付(ODD)链接
}

type PackageResp struct {
	ReferenceNumber  int     `json:"referenceNumber"`            // 包裹编号
	TrackingNumber   string  `json:"trackingNumber"`             // 跟踪号
	TrackingUrl      string  `json:"trackingUrl"`                // 跟踪链接
	VolumetricWeight float64 `json:"volumetricWeight,omitempty"` // 体积重量
}

type Document struct {
	ImageFormat            string `json:"imageFormat"`                      // 图像格式
	Content                string `json:"content"`                          // 图像内容
	TypeCode               string `json:"typeCode"`                         // 类型
	PackageReferenceNumber int    `json:"packageReferenceNumber,omitempty"` // 包裹编号
}

type ShipmentDetail struct {
	PickupDetails PickupDetails `json:"pickupDetails"` // 运单详情
}

type PickupDetails struct {
	LocalCutoffDateAndTime                string `json:"localCutoffDateAndTime"`                // 取货预订截止时间
	CutoffTimeOffset                      string `json:"cutoffTimeOffset"`                      // 取货预订截止时间（格林威治标准时间偏移量）如 PT30M
	PickupEarliest                        string `json:"pickupEarliest"`                        // DHL最早取件时间
	PickupLatest                          string `json:"pickupLatest"`                          // DHL最晚取件时间
	PickupCutoffSameDayOutboundProcessing string `json:"pickupCutoffSameDayOutboundProcessing"` // 本地取件截止时间，允许在要求的日期转发货件。在此取件截止时间之后的任何取件请求都可能影响运输时间
	TotalTransitDays                      string `json:"totalTransitDays"`                      // 运输天数
	PickupAdditionalDays                  string `json:"pickupAdditionalDays"`                  // 这是从上述城市或邮政区提货到到达服务区期间的额外运输延误（以天为单位）
	DeliveryAdditionalDays                string `json:"deliveryAdditionalDays"`                // 这是货物到达服务区后，送达指定城市或邮政区的额外运输延误（以天为单位）
	PickupDayOfWeek                       string `json:"pickupDayOfWeek"`                       // 星期几取货
	DeliveryDayOfWeek                     string `json:"deliveryDayOfWeek"`                     // 目的地星期几
}

type EstimatedDelivery struct {
	EstimatedDeliveryDate string `json:"estimatedDeliveryDate"` // 预计交货日期
	EstimatedDeliveryType string `json:"estimatedDeliveryType"` // 预计交付类型
}
