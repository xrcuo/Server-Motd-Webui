package api

import (
	"fmt"
	"net"
	"net/url"
)

func Ipx() {

	s := "//aqco.top:5432"

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
	var dc = ips[0]
	fmt.Println(dc)
}
