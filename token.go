package main

import (
	"encoding/json"

	"io/ioutil"
	
	"net/http"
	"strings"

//	"github.com/levigross/grequests"
)

const AccessTokenAPI = "https://api.weixin.qq.com/cgi-bin/token"

type Token struct {
	AccessToken string `json:"access_token"`
	Expire      int    `json:"expires_in"`
}

// 获取AppID的access_token
func (t *Token) Get(appid string, secret string) string {
/*	var params = map[string]string{
		"appid":      appid,
		"secret":     secret,
		"grant_type": "client_credential",
	}

	ro := &grequests.RequestOptions{
		Params: params,
	}
*/
    requestLine := strings.Join([]string{AccessTokenAPI,"?grant_type=client_credential&appid=",appid,"&secret=",secret}, "")

	resp, err := http.Get(requestLine)
//	res, _ := grequests.Get(AccessTokenAPI, ro)

	if err != nil || resp.StatusCode != http.StatusOK {
		return ""
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ""
	}

	if err := json.Unmarshal(body, t); err != nil {
		return ""
	}

	return string(body)
}
