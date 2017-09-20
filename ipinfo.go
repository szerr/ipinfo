package ipinfo

func GetIpInfo(ip string) (*IpInfo, error) {
	return Taobao(ip)
}
