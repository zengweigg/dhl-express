package model

type PickupsRequest struct {
	PlannedPickupDateAndTime string                  `json:"plannedPickupDateAndTime"` // 标识包裹可供取件的日期和时间。字符串的日期和时间部分应同时使用。日期不得为过去日期或未来 10 天以上的日期。时间是基于托运人时区的货件当地时间。日期部分必须采用以下格式：YYYY-MM-DD；时间部分必须采用以下格式：HH:MM:SS，并使用 24 小时制。日期和时间部分以字母 T 分隔（例如，2006-06-26T17:00:00 GMT+01:00）
	CloseTime                string                  `json:"closeTime"`                // 该地点可以派送 DHL Express 货件的最晚时间
	Location                 string                  `json:"location"`                 // 提供 DHL 快递员应在何处领取包裹的地点
	LocationType             string                  `json:"locationType"`             // 位置类型
	Accounts                 []Account               `json:"accounts"`                 // 账户
	SpecialInstructions      []SpecialInstruction    `json:"specialInstructions"`      // 详细说明您可能希望发送给 DHL 快递的特殊取件说明
	Remark                   string                  `json:"remark"`                   // 请提供额外的取货备注
	CustomerDetails          PickupsCustomerDetails  `json:"customerDetails"`          // 客户详情
	ShipmentDetails          []PickupsShipmentDetail `json:"shipmentDetails"`          // 发货详情
}

type PickupsCustomerDetails struct {
	ShipperDetails  *PickupsPartyDetails `json:"shipperDetails"`  // 发件人信息
	ReceiverDetails *PickupsPartyDetails `json:"receiverDetails"` // 收件件人信息
}

type PickupsPartyDetails struct {
	ContactInformation *ContactInfo   `json:"contactInformation"` // 发件人联系方式信息节点
	PostalAddress      *PostalAddress `json:"postalAddress"`      // 发件人地址信息节点
}

type PickupsShipmentDetail struct {
	Accounts              []Account                  `json:"accounts"`
	Packages              []PickupsPackage           `json:"packages"`
	ProductCode           string                     `json:"productCode"`           // 请提供货件的DHL Express Global产品代码
	DeclaredValue         float64                    `json:"declaredValue"`         // 请提供货件的DHL Express本地产品代码
	UnitOfMeasurement     string                     `json:"unitOfMeasurement"`     // 测量单位代码 ["metric","imperial"]
	ValueAddedServices    []PickupsValueAddedService `json:"valueAddedServices"`    // 增值服务
	IsCustomsDeclarable   bool                       `json:"isCustomsDeclarable"`   // 是否需要进行海关申报
	DeclaredValueCurrency string                     `json:"declaredValueCurrency"` // 货件申报货币代码
}

type PickupsPackage struct {
	Weight     float64     `json:"weight"`     // 快件单件的重量（即单箱重量）
	TypeCode   string      `json:"typeCode"`   // 包装类型代码
	Dimensions *Dimensions `json:"dimensions"` // 快件长宽高信息节点
}

type PickupsValueAddedService struct {
	ServiceCode string  `json:"serviceCode"`        // 特殊/增值服务代码
	Value       float64 `json:"value,omitempty"`    // 特殊/增值服务价值（当前主要用于填写保险价值）
	Currency    string  `json:"currency,omitempty"` // 特殊/增值服务的货币单位 CNY
}
