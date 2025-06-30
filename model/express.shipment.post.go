package model

// CreateShipmentData 制作运单+预约取件 格式参考文档https://developer.dhl.com/api-reference/dhl-express-mydhl-api#get-started-section/
type CreateShipmentData struct {
	PlannedShippingDateAndTime string                  `json:"plannedShippingDateAndTime"`            // 发件日期 启用预约取件服务时也会调用该字段的值作为快件备妥时间 格式 2025-04-05T10:30:00 GMT+08:00
	ProductCode                string                  `json:"productCode"`                           // Global产品代码 "普通包裹：P，正午特派包裹：Y 更多详情，请查看开发包中Reference_Data中的 Global Product Codes页面
	LocalProductCode           string                  `json:"localProductCode,omitempty"`            // 本地产品代码 对于CN来说，localProductCode通常与productCode一致。但对于某些国家/某些产品，两者可能存在差异，这取决于当地DHL的设定。 虽然localProductCode是选填字段，但出于数据完整传输考虑，建议在Request中保留localProductCode，一起传输。"
	GetRateEstimates           bool                    `json:"getRateEstimates,omitempty"`            // 是否返回预估运费"---true  返回预估运费； ---false 不返回预估运费 Request中未添加该字段时，将默认为false
	RequestOndemandDeliveryURL bool                    `json:"requestOndemandDeliveryURLM,omitempty"` // 是否返回ODD网址链 该字段的可选值为： ---true  将该字段的值设为true并结合发件人账号的配置，可以在Response中返回ODD网址---false 将该字段的值设为false，则在Response中不返回ODD网址Request中未添加该字段时，将默认为false
	GetTransliteratedResponse  bool                    `json:"getTransliteratedResponse"`             // 是否返回多语言
	Pickup                     *Pickup                 `json:"pickup"`                                // 预约取件节点
	CustomerDetails            *CustomerDetails        `json:"customerDetails"`                       // 客户信息节点
	Content                    *Content                `json:"content"`                               // 快件详情节点
	OutputImageProperties      *OutputImageProperties  `json:"outputImageProperties,omitempty"`       // 运单/发票格式、样式
	Accounts                   []Account               `json:"accounts"`                              // 运费以及税金支付账号
	ValueAddedServices         []ValueAddedService     `json:"valueAddedServices,omitempty"`          // 特殊/增值服务节点
	CustomerReferences         []Reference             `json:"customerReferences,omitempty"`          // 快件参考信息节点
	DocumentImages             []DocumentImage         `json:"documentImages,omitempty"`              // 无纸化单据影像上传节点
	OnDemandDelivery           *OnDemandDelivery       `json:"onDemandDelivery,omitempty"`            // ODD服务节点
	ShipmentNotification       []Notification          `json:"shipmentNotification,omitempty"`        // 运单生成通知节点
	EstimatedDeliveryDate      *DeliveryDateRequest    `json:"estimatedDeliveryDate,omitempty"`       // 预计派送日期
	GetAdditionalInformation   []AdditionalInfoRequest `json:"getAdditionalInformation,omitempty"`    // 返回额外的快件信息信息
	// ParentShipment             *ParentShipment         `json:"parentShipment,omitempty"`
	// Identifiers                []Identifier            `json:"identifiers,omitempty"`
	// PrepaidCharges             []Charge                `json:"prepaidCharges,omitempty"`
}

type BankDetail struct {
	Name                      string `json:"name,omitempty"`                      // 银行名称（仅适用于俄国）
	SettlementLocalCurrency   string `json:"settlementLocalCurrency,omitempty"`   // 银行本币结算账号（仅适用于俄国）
	SettlementForeignCurrency string `json:"settlementForeignCurrency,omitempty"` // 银行外币结算账号（仅适用于俄国）
}

type Pickup struct {
	IsRequested            bool                 `json:"isRequested"`                      // 是否启用预约取件功能 "设为false：不预约取件，只制作运单 设为true：启用预约取件服务，同时制作运单和预约取件
	CloseTime              string               `json:"closeTime,omitempty"`              // 当日最晚取件时间 格式17:00 建议启用预约取件服务时添加该字段
	Location               string               `json:"location,omitempty"`               // 具体取件处 最终取件地点（如取件地址的楼号、楼层、房间号等，无需写完整取件地址，只需写最小取件地址）
	SpecialInstructions    []SpecialInstruction `json:"specialInstructions,omitempty"`    // 取件特殊说明节点
	PickupDetails          *PartyDetails        `json:"pickupDetails,omitempty"`          // 预约取件地址节点 "与shipperDetails节点基本相同 （除了registrationNumbers-typeCode字段的可选值与shipperDetails节点可能存在差异），请参考shipperDetails节点的组织架构、各下级字段的名称以及字符限制
	PickupRequestorDetails *PartyDetails        `json:"pickupRequestorDetails,omitempty"` // 远程取件预约请求方信息 "与shipperDetails节点基本相同 （除了registrationNumbers-typeCode字段的可选值与shipperDetails节点可能存在差异）， 请参考shipperDetails节点的组织架构、各下级字段的名称以及字符限制
}

type PartyDetails struct {
	ContactInformation  *ContactInfo         `json:"contactInformation"`            // 发件人联系方式信息节点
	PostalAddress       *PostalAddress       `json:"postalAddress"`                 // 发件人地址信息节点
	TypeCode            string               `json:"typeCode,omitempty"`            // 发件人类别
	RegistrationNumbers []RegistrationNumber `json:"registrationNumbers,omitempty"` // 发件人注册号/税号节点
	BankDetails         []BankDetail         `json:"bankDetails,omitempty"`         // 发件人银行信息节点（仅适用于俄国）
}

type RegistrationNumber struct {
	Number            string `json:"number"`            // 发件人注册号/税号
	TypeCode          string `json:"typeCode"`          // 发件人注册号/税号类别
	IssuerCountryCode string `json:"issuerCountryCode"` // 发件人注册号/税号所属国国家代码
}

type ValueAddedService struct {
	ServiceCode    string          `json:"serviceCode"`        // 特殊/增值服务代码
	Value          float64         `json:"value,omitempty"`    // 特殊/增值服务价值（当前主要用于填写保险价值）
	Currency       string          `json:"currency,omitempty"` // 特殊/增值服务的货币单位 CNY
	Method         string          `json:"method,omitempty"`
	DangerousGoods []DangerousGood `json:"dangerousGoods,omitempty"` // 危险品节点
}

type DangerousGood struct {
	ContentId            string  `json:"contentId"` // 危险品的代码
	DryIceTotalNetWeight float64 `json:"dryIceTotalNetWeight,omitempty"`
	CustomDescription    string  `json:"customDescription,omitempty"`
	UnCodes              []int   `json:"unCodes,omitempty"`
}

type OutputImageProperties struct {
	EncodingFormat                    string        `json:"encodingFormat,omitempty"`                    // 运单格式 有效值有四个，分别为：pdf, zpl, lp2, epl  不添加此项将默认为pdf
	PrinterDPI                        int           `json:"printerDPI,omitempty"`                        // 运单打印精度 可选值： 200，300
	SplitTransportAndWaybillDocLabels bool          `json:"splitTransportAndWaybillDocLabels,omitempty"` // 拆分成三份：运单转运联，底联，以及由收据+发票组成的PDF文件 false
	AllDocumentsInOneImage            bool          `json:"allDocumentsInOneImage,omitempty"`            // 不拆分，运单转运联，底联，Receipt以及发票集中在一个PDF文件下
	SplitDocumentsByPages             bool          `json:"splitDocumentsByPages,omitempty"`             // 适用于一票多件时，拆分为每个单独的Transport Label以及单独的运单底联
	SplitInvoiceAndReceipt            bool          `json:"splitInvoiceAndReceipt,omitempty"`            // 拆分成三份：运单（转运联+底联），收据以及发票
	CustomerBarcodes                  []Barcode     `json:"customerBarcodes,omitempty"`                  // 客户条形码节点
	CustomerLogos                     []Logo        `json:"customerLogos,omitempty"`                     // 客户商标节点
	ImageOptions                      []ImageOption `json:"imageOptions"`                                // 该节点可基于typeCode展开多个
	ReceiptAndLabelsInOneImage        bool          `json:"receiptAndLabelsInOneImage,omitempty"`
}

// Barcode 客户条形码
type Barcode struct {
	Position         string `json:"position,omitempty"`
	SymbologyCode    string `json:"symbologyCode,"`             // 条形码类型
	Content          string `json:"content"`                    // 条形码
	TextBelowBarcode string `json:"textBelowBarcode,omitempty"` // 条形码下方文字
}

type Logo struct {
	FileFormat string `json:"fileFormat"` // 客户商标影像的类型
	Content    string `json:"content"`    // 客户商标的base64编码流
}

type ImageOption struct {
	TypeCode                        string `json:"typeCode"`                        // 单据类型 "可选值为： - label 运单转运联- waybillDoc 运单底联- invoice 清关发票- receipt 收据（只适用于俄国）
	TemplateName                    string `json:"templateName,omitempty"`          // 单据样式
	IsRequested                     bool   `json:"isRequested,omitempty"`           // 是否生成此类单据
	HideAccountNumber               bool   `json:"hideAccountNumber,omitempty"`     // 是否在运单上隐藏DHL账号
	NumberOfCopies                  int    `json:"numberOfCopies,omitempty"`        // 运单底联的份数 1
	InvoiceType                     string `json:"invoiceType,omitempty"`           // 清关发票类型
	LanguageCode                    string `json:"languageCode,omitempty"`          // 清关发票的语言
	RenderDHLLogo                   bool   `json:"renderDHLLogo"`                   // 是否在运单上打印DHL Logo
	LabelFreeText                   string `json:"labelFreeText,omitempty"`         // 该字段可用于在运单转运联上备注快件信息
	LabelCustomerDataText           string `json:"labelCustomerDataText,omitempty"` // 该字段可用于在运单转运联上备注快件信息
	LanguageCountryCode             string `json:"languageCountryCode,omitempty"`
	LanguageScriptCode              string `json:"languageScriptCode,omitempty"`
	EncodingFormat                  string `json:"encodingFormat,omitempty"`
	ShipmentReceiptCustomerDataText string `json:"shipmentReceiptCustomerDataText,omitempty"`
	FitLabelsToA4                   bool   `json:"fitLabelsToA4,omitempty"`
}

type Reference struct {
	Value    string `json:"value"`              // 参考信息内容
	TypeCode string `json:"typeCode,omitempty"` // 参考信息类型
}

type Identifier struct {
	TypeCode       string `json:"typeCode"`
	Value          string `json:"value"`
	DataIdentifier string `json:"dataIdentifier"`
}

type CustomerDetails struct {
	ShipperDetails           *PartyDetails `json:"shipperDetails,omitempty"`  // 发件人信息节点
	ReceiverDetails          *PartyDetails `json:"receiverDetails,omitempty"` // 收件人信息节点
	BuyerDetails             *PartyDetails `json:"buyerDetails,omitempty"`    // 买家信息节点
	ImporterDetails          *PartyDetails `json:"importerDetails,omitempty"` // 进口商信息节点
	ExporterDetails          *PartyDetails `json:"exporterDetails,omitempty"` // 出口商信息节点
	SellerDetails            *PartyDetails `json:"sellerDetails,omitempty"`   // 卖家信息节点
	PayerDetails             *PartyDetails `json:"payerDetails,omitempty"`    // 付款方信息节点
	ManufacturerDetails      *PartyDetails `json:"manufacturerDetails,omitempty"`
	UltimateConsigneeDetails *PartyDetails `json:"ultimateConsigneeDetails,omitempty"`
	BrokerDetails            *PartyDetails `json:"brokerDetails,omitempty"`
}

type Content struct {
	IsCustomsDeclarable   bool               `json:"isCustomsDeclarable"`         // 所寄快件的类别：包裹T/文件F
	Description           string             `json:"description"`                 // 货物描述
	DeclaredValue         float64            `json:"declaredValue"`               // 当Request中未添加exportDeclaration字段时，该字段的值将作为包裹快件的申报价值
	DeclaredValueCurrency string             `json:"declaredValueCurrency"`       // 货币单位
	UnitOfMeasurement     string             `json:"unitOfMeasurement"`           // 重量单位
	Incoterm              string             `json:"incoterm"`                    // 贸易条款
	USFilingTypeValue     string             `json:"USFilingTypeValue,omitempty"` // 美国出口备案类型
	ExportDeclaration     *ExportDeclaration `json:"exportDeclaration,omitempty"` // 海关申报节点
	Packages              []Package          `json:"packages"`                    // 所寄快件每一件（即每一箱）的重量、尺寸节点
}

type Package struct {
	Weight             float64          `json:"weight"`                       // 快件单件的重量（即单箱重量）
	Description        string           `json:"description,omitempty"`        // 单件货物的内容描述
	TypeCode           string           `json:"typeCode,omitempty"`           // 包装类型代码
	LabelDescription   string           `json:"labelDescription,omitempty"`   // 客户的额外信息备注
	Dimensions         *Dimensions      `json:"dimensions,omitempty"`         // 快件长宽高信息节点
	LabelBarcodes      []Barcode        `json:"labelBarcodes,omitempty"`      // 客户条形码节点
	LabelText          []LabelText      `json:"labelText,omitempty"`          // 转运联客户备注信息节点
	CustomerReferences []ReferencesCode `json:"customerReferences,omitempty"` // 单件参考信息节点
}

type LabelText struct {
	Position string `json:"position"` // 备注信息的位置
	Caption  string `json:"caption"`  // 备注信息的标题
	Value    string `json:"value"`    // 备注信息的内容
}

type ExportDeclaration struct {
	DestinationPortName string             `json:"destinationPortName,omitempty"` // 目的地港名称
	PlaceOfIncoterm     string             `json:"placeOfIncoterm,omitempty"`     // 该字段用来录入具体的贸易条款所适用的港口名称（启运港、装运港、目的港等）
	ShipmentType        string             `json:"shipmentType,omitempty"`        // 发货目的
	PackageMarks        string             `json:"packageMarks,omitempty"`        // 货物标识
	PayerVATNumber      string             `json:"payerVATNumber,omitempty"`      // 支付方商品及服务税号/增值税号
	RecipientReference  string             `json:"recipientReference,omitempty"`  // 收件人参考信息
	ExportReasonType    string             `json:"exportReasonType,omitempty"`    // 出口类型 参考值permanent
	ExportReason        string             `json:"exportReason,omitempty"`        // 出口原因
	ExportReference     string             `json:"exportReference,omitempty"`     // 申报单号
	Exporter            *Exporter          `json:"exporter,omitempty"`            // 出口商ID和出口商代码
	Licenses            []License          `json:"licenses,omitempty"`            // 许可证信息录入节点
	Invoice             *Invoice           `json:"invoice,omitempty"`             // 发票基础信息节点
	CustomsDocuments    []CustomsDocument  `json:"customsDocuments,omitempty"`    // 清关单据清单节点
	Remarks             []Remark           `json:"remarks,omitempty"`             // 发票备注节点
	AdditionalCharges   []AdditionalCharge `json:"additionalCharges,omitempty"`   // 其他费用节点
	DeclarationNotes    []DeclarationNote  `json:"declarationNotes,omitempty"`    // 发票额外声明节点
	LineItems           []LineItem         `json:"lineItems"`                     // 单项商品的详细信息节点
}

type LineItem struct {
	Number                          int               `json:"number"`                                    // 顺序号，用以区分每项商品
	Description                     string            `json:"description"`                               // 单项商品的描述
	Price                           float64           `json:"price"`                                     // 单价
	Quantity                        *Quantity         `json:"quantity"`                                  // 单项商品的数量以及数量单位节点
	CommodityCodes                  []CommodityCode   `json:"commodityCodes,omitempty"`                  // 单项商品的海关编码（HS CODE)
	ExportReasonType                string            `json:"exportReasonType,omitempty"`                // 单项商品的出口类型
	ManufacturerCountry             string            `json:"manufacturerCountry"`                       // 原产国国家代码
	Weight                          *Weight           `json:"weight"`                                    // 净重和毛重节点
	IsTaxesPaid                     bool              `json:"isTaxesPaid,omitempty"`                     // 发往新西兰的商品GST支付信息
	AdditionalInformation           string            `json:"additionalInformation,omitempty"`           // 该项商品的包装信息补充说明
	CustomerReferences              []Reference       `json:"customerReferences,omitempty"`              // 单项商品参考信息节点
	CustomsDocuments                []CustomsDocument `json:"customsDocuments,omitempty"`                // 单项商品的清关单据信息节点
	PreCalculatedLineItemTotalValue int               `json:"preCalculatedLineItemTotalValue,omitempty"` // 该项商品行项的总金额
}

type Quantity struct {
	Value             string `json:"quantity"` // 单项商品的数量
	UnitOfMeasurement string `json:"quantity"` // 数量单位
}

type Weight struct {
	NetValue   float64 `json:"netValue"`   // 净重
	GrossValue float64 `json:"grossValue"` // 毛重
}

type Invoice struct {
	Date               string      `json:"date"`                         // 发票日期
	Number             string      `json:"number"`                       // 发票号码
	SignatureName      string      `json:"signatureName,omitempty"`      // 清关发票签名
	SignatureTitle     string      `json:"signatureTitle,omitempty"`     // 签字人的职位
	SignatureImage     string      `json:"signatureImage,omitempty"`     // 签名影像（电子签名）的base64编码
	Instructions       string      `json:"instructions,omitempty"`       // 快件指令
	TotalNetWeight     int         `json:"totalNetWeight,omitempty"`     // 所寄全部商品的总净重（即所有lineItems的总净重）
	TotalGrossWeight   int         `json:"totalGrossWeight,omitempty"`   // 所寄全部商品的总毛重（即所有lineItems的总毛重）
	TermsOfPayment     string      `json:"termsOfPayment,omitempty"`     // 支付条款
	CustomerReferences []Reference `json:"customerReferences,omitempty"` // 清关发票参考信息节点
}

type Remark struct {
	Value string `json:"value"` // 发票备注信息
}

type DeclarationNote struct {
	Value string `json:"value"` // 发票额外声明
}

type AdditionalCharge struct {
	Caption  string  `json:"caption,omitempty"` // 费用名称
	Value    float64 `json:"value"`             // 费用金额
	TypeCode string  `json:"typeCode"`          // 费用类型
}

type Exporter struct {
	Id   string `json:"id,omitempty"`   // 出口商ID
	Code string `json:"code,omitempty"` // 出口商代码
}

type DocumentImage struct {
	Content     string `json:"content"`               // 单据影像Base64编码
	TypeCode    string `json:"typeCode,omitempty"`    // 单据影像类型
	ImageFormat string `json:"imageFormat,omitempty"` // 单据影像格式
}

// ODD服务
type OnDemandDelivery struct {
	DeliveryOption        string `json:"deliveryOption"`                  // 基于ODD服务的派送选项
	Location              string `json:"location,omitempty"`              // 快件投送的最终位置
	SpecialInstructions   string `json:"specialInstructions,omitempty"`   // 对于派送员的额外备注信息（有助于派送）
	GateCode              string `json:"gateCode,omitempty"`              // 进入派送地址公寓的门禁码
	WhereToLeave          string `json:"whereToLeave,omitempty"`          // 派送给邻居的服务代码
	NeighbourName         string `json:"neighbourName,omitempty"`         // 邻居的姓名
	NeighbourHouseNumber  string `json:"neighbourHouseNumber,omitempty"`  // 邻居的门牌号
	AuthorizerName        string `json:"authorizerName,omitempty"`        // 被授权人的姓名
	ServicePointId        string `json:"servicePointId,omitempty"`        // 指定的DHL服务点ID
	RequestedDeliveryDate string `json:"requestedDeliveryDate,omitempty"` // 要求派送日期
}

type Notification struct {
	TypeCode            string `json:"typeCode"`                      // 通知方式
	ReceiverId          string `json:"receiverId"`                    // 电子邮件地址
	BespokeMessage      string `json:"bespokeMessage,omitempty"`      // 通知内容
	LanguageCode        string `json:"languageCode,omitempty"`        // 通知语言
	LanguageCountryCode string `json:"languageCountryCode,omitempty"` // 通知语言所属国家代码
}

type DeliveryDateRequest struct {
	IsRequested bool   `json:"isRequested"`        // 是否返回预计派送日期
	TypeCode    string `json:"typeCode,omitempty"` // 预计派送日期类型
}

type AdditionalInfoRequest struct {
	TypeCode    string `json:"typeCode"`    // 额外快件信息的类别 "该字段的可选值为： - pickupDetails- optionalShipmentData- barcodeInformation"
	IsRequested bool   `json:"isRequested"` // 是否返回此类信息
}

type CustomsDocument struct {
	TypeCode string `json:"typeCode"` // 清关单据类型
	Value    string `json:"value"`    // 清关单据的ID
}

type License struct {
	TypeCode string `json:"typeCode"` // 许可证类型
	Value    string `json:"value"`    // 许可证号`
}

type CommodityCode struct {
	TypeCode string `json:"typeCode"` // 单项商品的海关编码类别
	Value    string `json:"value"`    // 单项商品的海关编码
}

type ReferencesCode struct {
	TypeCode string `json:"typeCode"` // 单件参考信息类型
	Value    int    `json:"value"`    // 单件参考信息内容
}
