package ipinfo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const TAOBAO_API = "http://ip.taobao.com/service/getIpInfo2.php?ip="

type TaobaoIpData struct {
	Country    string
	Country_id string
	Area       string
	Area_id    string
	Region     string
	Region_id  string
	City       string
	City_id    string
	County     string
	County_id  string
	Isp        string
	Isp_id     string
	Ip         string
}

type TaobaoIpInfo struct {
	Code int
	Data TaobaoIpData
}

func Taobao(ip string) (*IpInfo, error) {
	/*每个ip访问频率需小于10s*/
	if ip == "" {
		ip = "myip"
	}
	info := new(IpInfo)
	resp, err := http.Get(TAOBAO_API + ip)
	if err != nil {
		return info, err
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return info, err
	}
	t := new(TaobaoIpInfo)
	if json.Unmarshal(data, t) != nil {
		return info, err
	}
	if t.Code != 0 {
		return info, fmt.Errorf("code error:%d", t.Code)
	}
	i := info
	i.Ip = t.Data.Ip
	i.Isp = t.Data.Isp
	i.Country = t.Data.Country
	i.Area = t.Data.Area
	i.Region = t.Data.Region
	i.City = t.Data.City
	i.County = t.Data.County
	return i, err
}
