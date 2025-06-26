package dhl_express

import (
	"fmt"
	"github.com/zengweigg/dhl-express/config"
	"github.com/zengweigg/dhl-express/model"
	"testing"
	"time"
)

func Test001(m *testing.T) {
	// 初始化
	cfg := config.LoadConfig()
	DhlClient := NewDHLService(*cfg)
	// 单独预约取件
	// 构造测试请求数据
	loc, _ := time.LoadLocation("Asia/Shanghai")
	currentTime := time.Now().Add(50 * time.Hour) // 加上 24 小时，即明天的此刻
	formattedTime := currentTime.In(loc).Format("2006-01-02T15:04:05+08:00")
	accounts := []model.Account{
		{
			Number:   cfg.CustomerCode,
			TypeCode: "shipper",
		},
	}
	specialinstructions := []model.SpecialInstruction{
		{
			TypeCode: "TBD",
			Value:    "please ring front desk",
		},
	}

	shipmentdetails := []model.PickupsShipmentDetail{
		{
			Accounts:              accounts,
			ProductCode:           "P",
			DeclaredValue:         30.0,
			DeclaredValueCurrency: "EUR",
			IsCustomsDeclarable:   false,
			UnitOfMeasurement:     "metric",
			ValueAddedServices: []model.PickupsValueAddedService{
				{
					ServiceCode: "II",
					Value:       30.0,
					Currency:    "GBP",
				},
			},
			Packages: []model.PickupsPackage{
				{
					TypeCode: "3BX",
					Weight:   2.5,
					Dimensions: &model.Dimensions{
						Length: 20,
						Width:  25,
						Height: 15,
					},
				},
			},
		},
	}
	postData := model.PickupsRequest{
		PlannedPickupDateAndTime: formattedTime,
		CloseTime:                "20:30",
		Location:                 "nebula公司",
		LocationType:             "business",
		Remark:                   "two parcels required pickup",
		Accounts:                 accounts,
		SpecialInstructions:      specialinstructions,
		ShipmentDetails:          shipmentdetails,
		CustomerDetails: model.PickupsCustomerDetails{
			ReceiverDetails: &model.PickupsPartyDetails{
				ContactInformation: &model.ContactInfo{
					FullName:    "Adam Spencer",
					Email:       "that@before.gb",
					Phone:       "+1123456789",
					MobilePhone: "+60112345678",
					CompanyName: "Company Name",
				},
				PostalAddress: &model.PostalAddress{
					AddressLine1: "498 Bromford Gate",
					AddressLine2: "Bromford Lane",
					AddressLine3: "Erdington",
					CityName:     "BIRMINGHAM",
					CountryCode:  "GB",
					PostalCode:   "B24 8DW",
				},
			},
			ShipperDetails: &model.PickupsPartyDetails{
				ContactInformation: &model.ContactInfo{
					FullName:    "张三",
					Email:       "zhangsan@example.com",
					Phone:       "+8613645678890",
					CompanyName: cfg.Platform,
				},
				PostalAddress: &model.PostalAddress{
					AddressLine1: "yue lu qu lu gu ren min dong lu 123",
					CityName:     "Chang Sha",
					CountryCode:  "CN",
					PostalCode:   "410000",
				},
			},
		},
	}
	resp, err := DhlClient.Services.Base.PickupsCreate(postData)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Println("666", resp.DispatchConfirmationNumbers, resp.ReadyByTime, resp.NextPickupDate)
}
