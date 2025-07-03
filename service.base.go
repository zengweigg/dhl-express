package dhl_express

import (
	"errors"
	"fmt"
	"github.com/bytedance/sonic"
	"github.com/zengweigg/dhl-express/model"
	"net/http"
)

type baseService service

// 创建运单+预约取件
func (s baseService) ShipmentCreate(postData model.CreateShipmentData) (model.ShipmentResp, error) {
	respData := model.ShipmentResp{}
	//请求数据
	resp, err := s.httpClient.R().
		SetBody(postData).
		Post("shipments")
	fmt.Println(string(resp.Body()))
	if err != nil {
		return respData, err
	}
	// 检查HTTP状态码
	if resp.StatusCode() != http.StatusCreated {
		// 处理错误响应
		var errorResp struct {
			Message         string   `json:"message"`
			AdditionalDetails []string `json:"additionalDetails"`
		}
		if err := sonic.Unmarshal(resp.Body(), &errorResp); err != nil {
			return respData, fmt.Errorf("failed to parse error response: %v", err)
		}
		return respData, fmt.Errorf("API返回错误: %s, 详情: %v", errorResp.Message, errorResp.AdditionalDetails)
	}
	// 解析数据
	err = sonic.Unmarshal(resp.Body(), &respData)
	if err != nil {
		return respData, err
	}
	return respData, nil
}

// 单独预约取件
func (s baseService) PickupsCreate(postData model.PickupsRequest) (model.PickupsResp, error) {
	respData := model.PickupsResp{}
	// 请求数据
	resp, err := s.httpClient.R().
		SetBody(postData).
		Post("pickups")
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

// 取消下单
func (s baseService) PickupsCancel(number string) (bool, error) {
	respData := model.ErrorResponse{}
	// 请求数据
	url := "pickups/" + number
	resp, err := s.httpClient.R().
		SetBody(model.RequestBody{}).
		Delete(url)
	if err != nil {
		return false, err
	}
	statusCode := resp.StatusCode()
	if statusCode == http.StatusOK {
		return true, nil
	}
	// 解析数据
	err = sonic.Unmarshal(resp.Body(), &respData)
	if err != nil {
		return false, err
	}
	return false, errors.New(respData.Detail)
}

// 获取面单文件 shipments/9356579890/get-image?shipperAccountNumber=SOME_STRING_VALUE&typeCode=SOME_STRING_VALUE&pickupYearAndMonth=YYYY-MM"
func (s baseService) ShipmentGetImage(number string) (model.ImageDocument, error) {
	respData := model.ImageDocument{}
	// 请求数据
	url := "shipments/" + number + "/get-image"
	resp, err := s.httpClient.R().
		Get(url)
	if err != nil {
		return respData, err
	}
	// 解析数据
	err = sonic.Unmarshal(resp.Body(), &respData)
	if err != nil {
		return respData, err
	}
	return respData, nil
}

// 查询面单详情 没有找到
