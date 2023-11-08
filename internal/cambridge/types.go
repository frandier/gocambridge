package cambridge

type UnitsResult struct {
	Toc Toc `json:"toc"`
}

type Toc struct {
	Result []Result `json:"result"`
}

type Result struct {
	Name     string       `json:"name"`
	ItemType string       `json:"item-type"`
	ItemCode string       `json:"item-code"`
	SubType  string       `json:"sub-type"`
	Items    []ResultItem `json:"items"`
}

type ResultItem struct {
	Name     string     `json:"name"`
	ItemType string     `json:"item-type"`
	ItemCode string     `json:"item-code"`
	SubType  string     `json:"sub-type"`
	Items    []ItemItem `json:"items"`
}

type ItemItem struct {
	Name                string              `json:"name"`
	ItemCode            string              `json:"item-code"`
	Resource            string              `json:"resource"`
	EXTCupXapiscoreable EXTCupXapiscoreable `json:"ext-cup-xapiscoreable"`
}

type EXTCupXapiscoreable struct {
	URL        string     `json:"url"`
	Filename   string     `json:"filename"`
	Filesize   int64      `json:"filesize"`
	Container  string     `json:"container"`
	CupOptions CupOptions `json:"cup-options"`
}

type CupOptions struct {
	Contentid string `json:"contentid"`
	Title     string `json:"title"`
	Engine    string `json:"engine"`
	DP        string `json:"dp"`
	LmsApis   string `json:"lmsApis"`
}
