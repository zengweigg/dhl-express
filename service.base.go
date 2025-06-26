package dhl_express

import (
	"fmt"
	"github.com/bytedance/sonic"
	"github.com/zengweigg/dhl-express/model"
)

type baseService service

func (s baseService) ShipmentCreate(postData model.CreateShipmentData) (model.ShipmentRequest, error) {
	respData := model.ShipmentRequest{}
	// 请求数据
	resp, err := s.httpClient.R().
		SetBody(postData).
		Post("shipments")
	fmt.Println(string(resp.Body()))
	if err != nil {
		return respData, err
	}
	// 解析数据
	err = sonic.Unmarshal(resp.Body(), &respData)
	if err != nil {
		return respData, nil
	}
	return respData, nil
}
