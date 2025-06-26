package model


package main
type ShipmentRequest struct {
	PlannedShippingDateAndTime string            `json:"plannedShippingDateAndTime"` // 标识包裹交付的日期和时间。字符串的日期和时间部分均需使用。日期不应为过去日期或未来 10 天以后的日期。时间表示货件发货地的当地时间加上相应的时区。日期部分的格式必须为：YYYY-MM-DD；时间部分的格式必须为：HH:MM:SS，使用 24 小时制。日期和时间部分以字母 T 分隔（例如，2025-01-18T17:00:00 GMT+01:00）
	Pickup                     *Pickup           `json:"pickup,omitempty"`
	ProductCode                string            `json:"productCode"`
	LocalProductCode           string            `json:"localProductCode,omitempty"`
	GetRateEstimates           bool              `json:"getRateEstimates,omitempty"`
	Accounts                   []Account         `json:"accounts,omitempty"`
	ValueAddedServices         []Service         `json:"valueAddedServices,omitempty"`
	DangerousGoods             []DangerousGood   `json:"dangerousGoods,omitempty"`
	OutputImageProperties      *OutputImage      `json:"outputImageProperties,omitempty"`
	CustomerReferences         []Reference       `json:"customerReferences,omitempty"`
	Identifiers                []Identifier      `json:"identifiers,omitempty"`
	CustomerDetails            *Customer         `json:"customerDetails"`
	Content                    *Content          `json:"content"`
	OnDemandDelivery           *OnDemandDelivery `json:"onDemandDelivery,omitempty"`
	ShipmentNotification       []Notification    `json:"shipmentNotification,omitempty"`
	PrepaidCharges             []PrepaidCharge   `json:"prepaidCharges,omitempty"`
	GetTransliteratedResponse  bool              `json:"getTransliteratedResponse,omitempty"`
	EstimatedDeliveryDate      *EDDRequest       `json:"estimatedDeliveryDate,omitempty"`
	GetAdditionalInformation   []InfoRequest     `json:"getAdditionalInformation,omitempty"`
	ParentShipment             *ParentShipment   `json:"parentShipment,omitempty"`
}

type Pickup struct {
	IsRequested            bool                 `json:"isRequested"`
	CloseTime              string               `json:"closeTime,omitempty"`
	Location               string               `json:"location,omitempty"`
	SpecialInstructions    []SpecialInstruction `json:"specialInstructions,omitempty"`
	PickupDetails          *PickupDetails       `json:"pickupDetails,omitempty"`
	PickupRequestorDetails *ContactDetails      `json:"pickupRequestorDetails,omitempty"`
}

type SpecialInstruction struct {
	Value    string `json:"value"`
	TypeCode string `json:"typeCode,omitempty"`
}

type PickupDetails struct {
	PostalAddress       *PostalAddress `json:"postalAddress"`
	ContactInformation  *ContactInfo   `json:"contactInformation"`
	RegistrationNumbers []Registration `json:"registrationNumbers,omitempty"`
	BankDetails         []BankDetail   `json:"bankDetails,omitempty"`
	TypeCode            string         `json:"typeCode,omitempty"`
}

type Account struct {
	TypeCode string `json:"typeCode"`
	Number   string `json:"number"`
}

type Service struct {
	ServiceCode string  `json:"serviceCode"`
	Value       float64 `json:"value,omitempty"`
	Currency    string  `json:"currency,omitempty"`
	Method      string  `json:"method,omitempty"`
}

type DangerousGood struct {
	ContentId            string   `json:"contentId"`
	DryIceTotalNetWeight float64  `json:"dryIceTotalNetWeight,omitempty"`
	CustomDescription    string   `json:"customDescription,omitempty"`
	UnCodes              []string `json:"unCodes,omitempty"`
}

type OutputImage struct {
	PrinterDPI       int               `json:"printerDPI,omitempty"`
	CustomerBarcodes []CustomerBarcode `json:"customerBarcodes,omitempty"`
	CustomerLogos    []CustomerLogo    `json:"customerLogos,omitempty"`
	EncodingFormat   string            `json:"encodingFormat,omitempty"`
	ImageOptions     []ImageOption     `json:"imageOptions,omitempty"`
}

type CustomerBarcode struct {
	Content          string `json:"content"`
	TextBelowBarcode string `json:"textBelowBarcode,omitempty"`
	SymbologyCode    string `json:"symbologyCode,omitempty"`
}

type CustomerLogo struct {
	FileFormat     string `json:"fileFormat"`
	Content        string `json:"content"`
	EncodingFormat string `json:"encodingFormat,omitempty"`
}

type ImageOption struct {
	TypeCode                          string `json:"typeCode"`
	TemplateName                      string `json:"templateName,omitempty"`
	IsRequested                       bool   `json:"isRequested,omitempty"`
	HideAccountNumber                 bool   `json:"hideAccountNumber,omitempty"`
	NumberOfCopies                    int    `json:"numberOfCopies,omitempty"`
	InvoiceType                       string `json:"invoiceType,omitempty"`
	LanguageCode                      string `json:"languageCode,omitempty"`
	LanguageCountryCode               string `json:"languageCountryCode,omitempty"`
	LanguageScriptCode                string `json:"languageScriptCode,omitempty"`
	RenderDHLLogo                     bool   `json:"renderDHLLogo,omitempty"`
	FitLabelsToA4                     bool   `json:"fitLabelsToA4,omitempty"`
	LabelFreeText                     string `json:"labelFreeText,omitempty"`
	LabelCustomerDataText             string `json:"labelCustomerDataText,omitempty"`
	ShipmentReceiptCustomerDataText   string `json:"shipmentReceiptCustomerDataText,omitempty"`
	SplitTransportAndWaybillDocLabels bool   `json:"splitTransportAndWaybillDocLabels,omitempty"`
	AllDocumentsInOneImage            bool   `json:"allDocumentsInOneImage,omitempty"`
	SplitDocumentsByPages             bool   `json:"splitDocumentsByPages,omitempty"`
	SplitInvoiceAndReceipt            bool   `json:"splitInvoiceAndReceipt,omitempty"`
	ReceiptAndLabelsInOneImage        bool   `json:"receiptAndLabelsInOneImage,omitempty"`
}

type Reference struct {
	Value    string `json:"value"`
	TypeCode string `json:"typeCode"`
}

type Identifier struct {
	TypeCode       string `json:"typeCode"`
	Value          string `json:"value"`
	DataIdentifier string `json:"dataIdentifier,omitempty"`
}

type Customer struct {
	ShipperDetails           *ContactDetails `json:"shipperDetails"`
	ReceiverDetails          *ContactDetails `json:"receiverDetails"`
	BuyerDetails             *ContactDetails `json:"buyerDetails,omitempty"`
	ImporterDetails          *ContactDetails `json:"importerDetails,omitempty"`
	ExporterDetails          *ContactDetails `json:"exporterDetails,omitempty"`
	SellerDetails            *ContactDetails `json:"sellerDetails,omitempty"`
	PayerDetails             *ContactDetails `json:"payerDetails,omitempty"`
	ManufacturerDetails      *ContactDetails `json:"manufacturerDetails,omitempty"`
	UltimateConsigneeDetails *ContactDetails `json:"ultimateConsigneeDetails,omitempty"`
	BrokerDetails            *ContactDetails `json:"brokerDetails,omitempty"`
}

type ContactDetails struct {
	PostalAddress       *PostalAddress `json:"postalAddress"`
	ContactInformation  *ContactInfo   `json:"contactInformation"`
	RegistrationNumbers []Registration `json:"registrationNumbers,omitempty"`
	BankDetails         []BankDetail   `json:"bankDetails,omitempty"`
	TypeCode            string         `json:"typeCode,omitempty"`
}

type PostalAddress struct {
	PostalCode   string `json:"postalCode,omitempty"`
	CityName     string `json:"cityName"`
	CountryCode  string `json:"countryCode"`
	ProvinceCode string `json:"provinceCode,omitempty"`
	AddressLine1 string `json:"addressLine1"`
	AddressLine2 string `json:"addressLine2,omitempty"`
	AddressLine3 string `json:"addressLine3,omitempty"`
	CountyName   string `json:"countyName,omitempty"`
	ProvinceName string `json:"provinceName,omitempty"`
	CountryName  string `json:"countryName,omitempty"`
}

type ContactInfo struct {
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	MobilePhone string `json:"mobilePhone,omitempty"`
	CompanyName string `json:"companyName"`
	FullName    string `json:"fullName"`
}

type Registration struct {
	TypeCode          string `json:"typeCode"`
	Number            string `json:"number"`
	IssuerCountryCode string `json:"issuerCountryCode,omitempty"`
}

type BankDetail struct {
	Name                      string `json:"name,omitempty"`
	SettlementLocalCurrency   string `json:"settlementLocalCurrency,omitempty"`
	SettlementForeignCurrency string `json:"settlementForeignCurrency,omitempty"`
}

type Content struct {
	Packages          []Package          `json:"packages"`
	Description       string             `json:"description,omitempty"`
	USFilingTypeValue string             `json:"USFilingTypeValue,omitempty"`
	Incoterm          string             `json:"incoterm,omitempty"`
	UnitOfMeasurement string             `json:"unitOfMeasurement,omitempty"`
	DocumentImages    []DocumentImage    `json:"documentImages,omitempty"`
	ExportDeclaration *ExportDeclaration `json:"exportDeclaration,omitempty"`
}

type Package struct {
	TypeCode              string         `json:"typeCode,omitempty"`
	Weight                float64        `json:"weight"`
	Dimensions            *Dimensions    `json:"dimensions,omitempty"`
	CustomerReferences    []Reference    `json:"customerReferences,omitempty"`
	Identifiers           []Identifier   `json:"identifiers,omitempty"`
	Description           string         `json:"description,omitempty"`
	LabelBarcodes         []LabelBarcode `json:"labelBarcodes,omitempty"`
	LabelText             []LabelText    `json:"labelText,omitempty"`
	LabelDescription      string         `json:"labelDescription,omitempty"`
	ReferenceNumber       int            `json:"referenceNumber,omitempty"`
	IsCustomsDeclarable   bool           `json:"isCustomsDeclarable,omitempty"`
	DeclaredValue         float64        `json:"declaredValue,omitempty"`
	DeclaredValueCurrency string         `json:"declaredValueCurrency,omitempty"`
}

type Dimensions struct {
	Length float64 `json:"length"`
	Width  float64 `json:"width"`
	Height float64 `json:"height"`
}

type LabelBarcode struct {
	Position         string `json:"position"`
	SymbologyCode    string `json:"symbologyCode"`
	Content          string `json:"content"`
	TextBelowBarcode string `json:"textBelowBarcode,omitempty"`
}

type LabelText struct {
	Position string `json:"position"`
	Caption  string `json:"caption,omitempty"`
	Value    string `json:"value,omitempty"`
}

type ExportDeclaration struct {
	LineItems           []LineItem         `json:"lineItems"`
	Invoice             *Invoice           `json:"invoice"`
	Remarks             []Remark           `json:"remarks,omitempty"`
	AdditionalCharges   []AdditionalCharge `json:"additionalCharges,omitempty"`
	DestinationPortName string             `json:"destinationPortName,omitempty"`
	PlaceOfIncoterm     string             `json:"placeOfIncoterm,omitempty"`
	PayerVATNumber      string             `json:"payerVATNumber,omitempty"`
	RecipientReference  string             `json:"recipientReference,omitempty"`
	Exporter            *Exporter          `json:"exporter,omitempty"`
	PackageMarks        string             `json:"packageMarks,omitempty"`
	DeclarationNotes    []DeclarationNote  `json:"declarationNotes,omitempty"`
	ExportReference     string             `json:"exportReference,omitempty"`
	ExportReason        string             `json:"exportReason,omitempty"`
	ExportReasonType    string             `json:"exportReasonType,omitempty"`
	Licenses            []License          `json:"licenses,omitempty"`
	ShipmentType        string             `json:"shipmentType,omitempty"`
	CustomsDocuments    []CustomsDocument  `json:"customsDocuments,omitempty"`
}

type LineItem struct {
	Number                          int               `json:"number"`
	Description                     string            `json:"description"`
	Price                           float64           `json:"price"`
	Quantity                        *Quantity         `json:"quantity"`
	CommodityCodes                  []CommodityCode   `json:"commodityCodes,omitempty"`
	ExportReasonType                string            `json:"exportReasonType,omitempty"`
	ManufacturerCountry             string            `json:"manufacturerCountry,omitempty"`
	Weight                          *ItemWeight       `json:"weight,omitempty"`
	IsTaxesPaid                     bool              `json:"isTaxesPaid,omitempty"`
	AdditionalInformation           []string          `json:"additionalInformation,omitempty"`
	CustomerReferences              []Reference       `json:"customerReferences,omitempty"`
	CustomsDocuments                []CustomsDocument `json:"customsDocuments,omitempty"`
	PreCalculatedLineItemTotalValue float64           `json:"preCalculatedLineItemTotalValue,omitempty"`
}

type Quantity struct {
	Value             int    `json:"value"`
	UnitOfMeasurement string `json:"unitOfMeasurement"`
}

type CommodityCode struct {
	TypeCode string `json:"typeCode"`
	Value    string `json:"value"`
}

type ItemWeight struct {
	NetValue   float64 `json:"netValue,omitempty"`
	GrossValue float64 `json:"grossValue,omitempty"`
}

type Invoice struct {
	Number                   string         `json:"number"`
	Date                     string         `json:"date"`
	SignatureName            string         `json:"signatureName,omitempty"`
	SignatureTitle           string         `json:"signatureTitle,omitempty"`
	SignatureImage           string         `json:"signatureImage,omitempty"`
	Instructions             []string       `json:"instructions,omitempty"`
	CustomerDataTextEntries  []string       `json:"customerDataTextEntries,omitempty"`
	TotalNetWeight           float64        `json:"totalNetWeight,omitempty"`
	TotalGrossWeight         float64        `json:"totalGrossWeight,omitempty"`
	CustomerReferences       []Reference    `json:"customerReferences,omitempty"`
	TermsOfPayment           string         `json:"termsOfPayment,omitempty"`
	IndicativeCustomsValues  *CustomsValues `json:"indicativeCustomsValues,omitempty"`
	PreCalculatedTotalValues *TotalValues   `json:"preCalculatedTotalValues,omitempty"`
}

type CustomsValues struct {
	ImportCustomsDutyValue        float64 `json:"importCustomsDutyValue,omitempty"`
	ImportTaxesValue              float64 `json:"importTaxesValue,omitempty"`
	TotalWithImportDutiesAndTaxes float64 `json:"totalWithImportDutiesAndTaxes,omitempty"`
}

type TotalValues struct {
	PreCalculatedTotalGoodsValue   float64 `json:"preCalculatedTotalGoodsValue,omitempty"`
	PreCalculatedTotalInvoiceValue float64 `json:"preCalculatedTotalInvoiceValue,omitempty"`
}

type Remark struct {
	Value string `json:"value"`
}

type AdditionalCharge struct {
	Value    float64 `json:"value"`
	Caption  string  `json:"caption,omitempty"`
	TypeCode string  `json:"typeCode"`
}

type Exporter struct {
	Id   string `json:"id,omitempty"`
	Code string `json:"code,omitempty"`
}

type DeclarationNote struct {
	Value string `json:"value"`
}

type License struct {
	TypeCode string `json:"typeCode"`
	Value    string `json:"value"`
}

type CustomsDocument struct {
	TypeCode string `json:"typeCode"`
	Value    string `json:"value"`
}

type DocumentImage struct {
	TypeCode    string `json:"typeCode"`
	ImageFormat string `json:"imageFormat"`
	Content     string `json:"content"`
}

type OnDemandDelivery struct {
	DeliveryOption             string `json:"deliveryOption"`
	Location                   string `json:"location,omitempty"`
	SpecialInstructions        string `json:"specialInstructions,omitempty"`
	GateCode                   string `json:"gateCode,omitempty"`
	WhereToLeave               string `json:"whereToLeave,omitempty"`
	NeighbourName              string `json:"neighbourName,omitempty"`
	NeighbourHouseNumber       string `json:"neighbourHouseNumber,omitempty"`
	AuthorizerName             string `json:"authorizerName,omitempty"`
	ServicePointId             string `json:"servicePointId,omitempty"`
	RequestedDeliveryDate      string `json:"requestedDeliveryDate,omitempty"`
	RequestOndemandDeliveryURL bool   `json:"requestOndemandDeliveryURL,omitempty"`
}

type Notification struct {
	TypeCode            string `json:"typeCode"`
	ReceiverId          string `json:"receiverId"`
	LanguageCode        string `json:"languageCode,omitempty"`
	LanguageCountryCode string `json:"languageCountryCode,omitempty"`
	BespokeMessage      string `json:"bespokeMessage,omitempty"`
}

type PrepaidCharge struct {
	TypeCode string  `json:"typeCode"`
	Currency string  `json:"currency"`
	Value    float64 `json:"value"`
	Method   string  `json:"method"`
}

type EDDRequest struct {
	IsRequested bool   `json:"isRequested"`
	TypeCode    string `json:"typeCode,omitempty"`
}

type InfoRequest struct {
	TypeCode    string `json:"typeCode"`
	IsRequested bool   `json:"isRequested"`
}

type ParentShipment struct {
	ProductCode   string `json:"productCode"`
	PackagesCount int    `json:"packagesCount"`
}
