package fake

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

func init()  {
	fmt.Println("thi is init")
	FackIP()
}

var ips = make([]string,200)

var userAgentSlice = [...]string{
	"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.132 Safari/537.36",
	"Mozilla/5.0 (Windows NT 6.3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/31.0.1650.63 Safari/537.36",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/38.0.2125.111 Safari/537.36",
	"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.1",
	"Mozilla/5.0 (Windows NT 6.2) AppleWebKit/536.3",
	"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/535.24",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_8_0) AppleWebKit/536.3",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/65.0.3325.181 Safari/537.36",
	"Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko)",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/60.0.3112.113 Safari/537.36",
	"Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.11 (KHTML, like Gecko) Chrome/23.0.1271.64 Safari/537.11",
	"Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/55.0.2883.103 Safari/537.36",
	"Mozilla/5.0 (Windows NT 6.3; WOW64; Trident/7.0; rv:11.0) like Gecko",
}

func GetUserAgent() string {
	n := rand.Intn(len(userAgentSlice))
	return userAgentSlice[n]
}

func GetIP() string {

	n := rand.Intn(len(ips))
	if n == 0 {
		return GetIP()
	}
	ip := ips[n]
	ips = append(ips[:n-1],ips[n:]...)
	return ip
}

func FackIP() {
	resp, err := http.Get("http://www.superfastip.com/api/ip?tid=99961bd7c1da5b60517f321b0787f464&num=200")
	if err != nil {

		log.Printf("http.NewRequest err: %v", err)
	}
	defer resp.Body.Close()
	body,err := ioutil.ReadAll(resp.Body)
	body2 := fmt.Sprintf("%s",body)
	body3 := strings.Replace(strings.Split(body2,"0,success")[1],",",":",-1)
	ips = strings.Split(body3,"<br />")[1:]
	//fmt.Print(ips)
}


func FackRequest(request *http.Request)(*http.Client) {
	if len(ips) < 10 {
		FackIP()
	}
	request.Header.Add("Accept", "application/json, text/javascript, */*; q=0.01")
	request.Header.Add("Accept-Encoding", "gzip, deflate, br")
	request.Header.Add("Accept-Languag", "zh-CN,zh;q=0.9")
	request.Header.Add("Connection", "keep-alive")
	request.Header.Add("Content-Length", "25")
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	request.Header.Add("Host", "www.lagou.com")
	request.Header.Add("Origin", "https://www.lagou.com")
	request.Header.Add("Referer", "https://www.lagou.com/jobs/list_golang?labelWords=&fromSearch=true&suginput=")
	//request.Header.Add("User-Agent", GetUserAgent())
	request.Header.Add("User-Agent", GetUserAgent())
	request.Header.Add("Cookie", "JSESSIONID=ABAAABAAAGGABCB6A0C5F80A6786A202852379B0B7B8F2A; _ga=GA1.2.1039849811.1544434721; user_trace_token=20181210173843-627abacf-fc5f-11e8-8ced-5254005c3644; LGUID=20181210173843-627abe2e-fc5f-11e8-8ced-5254005c3644; X_HTTP_TOKEN=1fb1fd7be3f42a53bda27d710e0c00ac; Hm_lvt_4233e74dff0ae5bd0a3d81c6ccf756e6=1543829061,1544434724,1544435109; index_location_city=%E5%8C%97%E4%BA%AC; sensorsdata2015jssdkcross=%7B%22distinct_id%22%3A%221679781cb1ec8b-02b8ef60a729f4-35677607-1764000-1679781cb1f71d%22%2C%22%24device_id%22%3A%221679781cb1ec8b-02b8ef60a729f4-35677607-1764000-1679781cb1f71d%22%7D; TG-TRACK-CODE=index_navigation; _gid=GA1.2.170535933.1544761188; LGSID=20181214121948-7ea9da5a-ff57-11e8-917e-525400f775ce; PRE_UTM=; PRE_HOST=; PRE_SITE=; PRE_LAND=https%3A%2F%2Fwww.lagou.com%2F; _gat=1; ab_test_random_num=0; sm_auth_id=028voxp4xmig8lhk; _putrc=01B4B2F45814A039123F89F2B170EADC; login=true; unick=%E6%8B%89%E5%8B%BE%E7%94%A8%E6%88%B75584; hasDeliver=0; gate_login_token=6279d7edddcd50a31e13f1fe4cedc566f4d0e8e3034e5755769c3d1149a1f227; LGRID=20181214122510-3e9e4980-ff58-11e8-8cef-5254005c3644; Hm_lpvt_4233e74dff0ae5bd0a3d81c6ccf756e6=1544761511")

	//fmt.Println(fmt.Sprintf("first=false&pn=%d",calculate.PageNo))

	//proxy := func(_ *http.Request) (*url.URL, error) {
	//	return url.Parse(fmt.Sprintf("http://%s",GetIP()))//根据定义Proxy func(*Request) (*url.URL, error)这里要返回url.URL
	//}
	//fmt.Println(fmt.Sprintf("當前代理IP: http://%s",GetIP()))
	// timeout
	//transport := &http.Transport{Proxy: proxy}

	//return &http.Client{Transport: transport,Timeout:6*time.Second}
	return &http.Client{Timeout:10*time.Second}


}