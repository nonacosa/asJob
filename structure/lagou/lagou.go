package lagou


type Calculate struct {
	TotalCount int
	PageSize   int
	PageNo     int
}

/**
 解析 JSON：https://blog.golang.org/json-and-go
 */
type Result struct {
	City              string
	BusinessZones     []string
	CompanyFullName   string
	CompanyLabelList  []string
	CompanyShortName  string
	CompanySize       string
	CreateTime        string
	District          string
	Education         string
	FinanceStage      string
	FirstType         string
	IndustryField     string
	IndustryLables    []string
	JobNature         string
	Latitude          string
	Longitude         string
	PositionAdvantage string
	PositionId        int32
	PositionLables    []string
	PositionName      string
	Salary            string
	SecondType        string
	Stationname       string
	Subwayline        string
	Linestaion        string
	WorkYear          string
}

type ListResult struct {
	Code    int
	Success bool
	Msg     string
	Content Content
}

type Content struct {
	PositionResult PositionResult
	PageNo         int
	PageSize       int
}

type PositionResult struct {
	Result     []Result
	TotalCount int
}

type jobService struct {
	City string
}
