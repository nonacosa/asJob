package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
	"github.com/pkwenda/asJob/src/structure/lagou"
)

const PageSize = 15

type lagouCalculate struct {
  lagou.Calculate
}

func (calculate *lagouCalculate) NextPageNo() (int) {
	if calculate.PageNo >= calculate.MaxPageNo() {

	}
	return calculate.PageNo + 1
}

func (calculate *lagouCalculate) MaxPageNo() (int) {
	return calculate.TotalCount / calculate.PageSize
}

func (calculate *lagouCalculate) setCurrentPageNo(pageNo int) {
	calculate.PageNo = pageNo
}


//func NewJobService(city string) *jobService {
//	return &jobService{City: city}
//}

func main() {
	calculate := lagouCalculate{Calculate:lagou.Calculate{TotalCount:0,PageSize:PageSize,PageNo:0}}
	for i:=0; i<2; i++ {


		//os.Create("job.log")
		logFile, err := os.OpenFile("job.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
		defer logFile.Close()
		if err != nil {
			log.Fatalln("open file error!")
		}
		debugLog := log.New(logFile, "--[Debug]--", log.Ltime)
		debugLog.SetFlags(debugLog.Flags() | log.LstdFlags)
		debugLog.Println("A debug message here")

		time.Sleep(time.Millisecond * 1000)

		url := "https://www.lagou.com/jobs/positionAjax.json?px=new&city=%E5%8C%97%E4%BA%AC&needAddtionalResult=false"
		client := http.Client{}

		postReader := strings.NewReader(fmt.Sprintf("first=false&pn=%d",calculate.PageNo))
		fmt.Println(fmt.Sprintf("first=false&pn=%d",calculate.PageNo))
		req, err := http.NewRequest("POST", url, postReader)
		if err != nil {
			log.Printf("http.NewRequest err: %v", err)
		}

		//req.Header.Set("Proxy-Switch-Ip", "yes")

		req.Header.Add("Accept", "application/json, text/javascript, */*; q=0.01")
		req.Header.Add("Accept-Encoding", "gzip, deflate, br")
		req.Header.Add("Accept-Languag", "zh-CN,zh;q=0.9")
		req.Header.Add("Connection", "keep-alive")
		req.Header.Add("Content-Length", "25")
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
		req.Header.Add("Host", "www.lagou.com")
		req.Header.Add("Origin", "https://www.lagou.com")
		req.Header.Add("Referer", "https://www.lagou.com/jobs/list_golang?labelWords=&fromSearch=true&suginput=")
		req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.110 Safari/537.36")
		req.Header.Add("Cookie", "_ga=GA1.2.161331334.1522592243; ")

		resp, err := client.Do(req)

		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()
		if resp.StatusCode == http.StatusOK {
			all, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				panic(err)
			}

			fmt.Printf("%s \n", all)

			var results lagou.ListResult
			json.Unmarshal([]byte(all), &results)
			calculate.setCurrentPageNo(calculate.NextPageNo())
			for _, v := range results.Content.PositionResult.Result {
				//fmt.Println(v)
				debugLog.Println(v)
			}

		} else {
			fmt.Printf("Error-code: %d \n", resp.StatusCode)
		}
	}

}
