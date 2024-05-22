package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
)

type Yiyan struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		Text     string `json:"text"`
		Location struct {
			Ip   string `json:"ip"`
			Area string `json:"area"`
		} `json:"location"`
	} `json:"data"`
}

var urx = "https://iuxcn.cn"

func YiYan() (H ZW, err error) {
	req, _ := http.NewRequest("GET", urx+"/Api/YiYan?format=json", nil)
	res, _ := http.DefaultClient.Do(req)
	body, _ := ioutil.ReadAll(res.Body)
	var M Yiyan
	fmt.Print(M)
	json.Unmarshal([]byte(body), &M)
	H.Code = M.Code
	H.Msg = M.Msg
	H.Text = M.Data.Text
	return H, err
}

func Ip(ipx string) (H ZW, err error) {
	s := "//" + ipx

	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}
	host, port, _ := net.SplitHostPort(u.Host)
	fmt.Println(host)
	fmt.Println(port)

	// 域名解析
	ips, err := net.LookupIP(host)
	if err != nil {
		fmt.Println("域名解析失败:", err)
		return
	}
	var dcip = ips[0]
	req, _ := http.NewRequest("GET", urx+"/Api/IP?format=json&ip="+dcip.String(), nil)
	res, _ := http.DefaultClient.Do(req)
	body, _ := ioutil.ReadAll(res.Body)
	var M Yiyan
	fmt.Print(M)
	json.Unmarshal([]byte(body), &M)
	H.Code = M.Code
	H.Msg = M.Msg
	H.Location = M.Data.Location.Area
	H.Ip = M.Data.Location.Ip
	return H, err
}

type ZW struct {
	Code string
	Msg  string

	Ip       string
	Location string
	Text     string
}
