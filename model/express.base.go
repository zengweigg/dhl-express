package model

type ErrorResponse struct {
	Instance string `json:"instance"` // 示例: "/expressapi/pickups/PRG211126000382"
	Detail   string `json:"detail"`   // 示例: "Missing mandatory parameters: requestorName"
	Title    string `json:"title"`    // 示例: "Missing parameters"
	Message  string `json:"message"`  // 示例: "Bad request"
	Status   string `json:"status"`   // 示例: "400"
}

type Account struct {
	Number   string `json:"number"`   // DHL账号 最长12位
	TypeCode string `json:"typeCode"` // DHL账号类别（账号用途）"该字段的可选值为：
	// - shipper - 发件人
	// - payer - 运费支付方
	// - duties-taxes - 目的地税金支付方
	// 如果accounts节点只添加typeCode=shipper并录入相应的number(账号），则表示使用该发件人账号支付运费；
	// 如果如果accounts节点同时添加typeCode=shipper和typeCode=payer，并相应录入两个number(账号），则typeCode=payer所对应的number会作为支付运费的账号；
	// 如果accounts节点添加了typeCode=duties-taxes并录入相应的number(账号）,那么该number将被用于支付目的地税金。如果该number是发件人账号或第三国账号（发件人、收件人以外的国家），需要同时在valueAddedServices节点录入特殊服务代码DD，从而启用DTP服务。
	// 注：启用DTP会产生额外服务费用。
}

type SpecialInstruction struct {
	Value    string `json:"value"` // 预约取件特殊说明 预约取件的备注
	TypeCode string `json:"typeCode"`
}

type ContactInfo struct {
	FullName    string `json:"fullName"`              // 联系人姓名
	CompanyName string `json:"companyName"`           // 公司名称
	Phone       string `json:"phone"`                 // 电话号码
	Email       string `json:"email,omitempty"`       // 电子邮箱
	MobilePhone string `json:"mobilePhone,omitempty"` // 手机号码
}

type PostalAddress struct {
	AddressLine1 string `json:"addressLine1"`           // 发件人地址栏1
	AddressLine2 string `json:"addressLine2,omitempty"` // 发件人地址栏2
	AddressLine3 string `json:"addressLine3,omitempty"` // 发件人地址栏3
	CityName     string `json:"cityName"`               // 发件人城市名称
	ProvinceCode string `json:"provinceCode,omitempty"` // 发件人州代码
	ProvinceName string `json:"provinceName,omitempty"` // 发件人州名/省名
	PostalCode   string `json:"postalCode"`             // 发件人邮编
	CountryCode  string `json:"countryCode"`            // 发件人国家代码
	CountyName   string `json:"countyName,omitempty"`   // 发件人所在城市的郊区名
	CountryName  string `json:"countryName,omitempty"`  // 发件人所在城市的郊区名
}

// 快件长宽高信息节点
type Dimensions struct {
	Length float64 `json:"length"` // 单件的长度
	Width  float64 `json:"width"`  // 单件的宽度
	Height float64 `json:"height"` // 单件的高度
}

// 空的结构体
type RequestBody struct{}

type DocumentResponse struct {
	Documents []Document `json:"documents"`
}

type ImageDocument struct {
	ShipmentTrackingNumber string `json:"shipmentTrackingNumber"` // 示例: "1234567890"
	TypeCode               string `json:"typeCode"`               // 示例: "waybill"
	Function               string `json:"function"`               // 示例: "import"
	EncodingFormat         string `json:"encodingFormat"`         // 示例: "PDF"
	Content                string `json:"content"`                // Base64编码的文档内容
}
