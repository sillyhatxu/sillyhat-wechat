package service

import (
	"encoding/json"
	log "github.com/xushikuan/microlog"
	"io/ioutil"
	"net/http"
	"sillyhat-wechat/common"
	"sillyhat-wechat/model"
)

func GetToken() (*model.Accesstoken, error) {
	client := &http.Client{}
	url := common.WECHAT_URL + common.METHOD_TOKEN + "&appid=" + common.APPID + "&secret=" + common.SECRET
	log.Infof("url : %v", url)
	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var accesstoken *model.Accesstoken
	err = json.Unmarshal(body, &accesstoken)
	if err != nil {
		return nil, err
	}
	return accesstoken, nil
}
