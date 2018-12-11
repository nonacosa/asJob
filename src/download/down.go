package download

import (
	"encoding/json"
	"fmt"
	"github.com/pkwenda/asJob/src/fake"
	"github.com/pkwenda/asJob/src/structure/lagou"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
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

var calculate lagouCalculate

//func NewJobService(city string) *jobService {
//	return &jobService{City: city}
//}

func init() {
	calculate = lagouCalculate{Calculate: lagou.Calculate{TotalCount: 0, PageSize: PageSize, PageNo: 0}}
}

func Worker(chanCount int)   {
	for i:=0; i<chanCount; i++ {
		if err := Spider(); err != nil {
			fmt.Println(err)
		}
	}
}
func Spider() (error) {

		//time.Sleep(time.Second * 10)
		//os.Create("job.log")
		logFile, err := os.OpenFile("job.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
		defer logFile.Close()
		if err != nil {
			log.Fatalln("open file error!")
		}
		debugLog := log.New(logFile, "--[Debug]--", log.Ltime)
		debugLog.SetFlags(debugLog.Flags() | log.LstdFlags)
		debugLog.Println("A debug message here")

		lgUrl := "https://www.lagou.com/jobs/positionAjax.json?px=new&city=%E5%8C%97%E4%BA%AC&needAddtionalResult=false"

		postReader := strings.NewReader(fmt.Sprintf("first=false&pn=%d", calculate.PageNo))

		req, err := http.NewRequest("POST", lgUrl, postReader)
		client := fake.FackRequest(req)
		if err != nil {
			return err
			log.Printf("http.NewRequest err: %v", err)
		}

		//req.Header.Set("Proxy-Switch-Ip", "yes")
		fake.FackRequest(req)

		resp, err := client.Do(req)
		if err != nil {
			return err
			log.Printf("client.Do: %v", err)
		}

		if resp != nil && resp.StatusCode == http.StatusOK {
			all, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return err
				log.Printf("ioutil.ReadAll err: %v", err)
			}

			fmt.Printf("%s \n", all)

			var results lagou.ListResult
			json.Unmarshal([]byte(all), &results)
			if results.Success {
				calculate.setCurrentPageNo(calculate.NextPageNo())

			}

			for _, v := range results.Content.PositionResult.Result {
				//fmt.Println(v)
				debugLog.Println(v)
			}

		} else {
			fmt.Printf("Error-code: %d \n", "")
		}

		defer func() {
			p := recover()
			if p != nil {
				debugLog.Print(p)
			}
		}()

		defer resp.Body.Close()
		return nil

}
