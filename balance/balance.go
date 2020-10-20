package balance

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"weixin_shop_pay"
	"weixin_shop_pay/tools"
)

// Balance 余额
type Balance struct {
	Config *weixin_shop_pay.Config
}

// SubMch 二级商户余额查询
func (c *Balance) SubMch(p *weixin_shop_pay.BalanceSubMch) (*SubMchResp, error) {

	// 请求参数
	dataJsonByte, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}

	// 发起请求
	urlPath := "v3/ecommerce/fund/balance/" + p.SubMchid
	resp, err := tools.GetRequest(c.Config, urlPath, dataJsonByte, c.Config.KeyPath)
	if err != nil {
		return nil, err
	}

	// 解析返回参数
	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	log.Println(string(respData))
	var output SubMchResp
	err = json.Unmarshal(respData, &output)
	if err != nil {
		return nil, err
	}
	return &output, nil
}

// SubMchResp .
type SubMchResp struct {
	SubMchid        string `json:"sub_mchid"`        // 二级商户号
	AvailableAmount int64  `json:"available_amount"` // 可用余额
	PendingAmount   int64  `json:"pending_amount"`   // 不可用余额
}