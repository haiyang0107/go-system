package request

type GetById struct {
	Id float64 `json:"id" form:"id"`
}

type GetIds struct {
	Ids []int `json:"ids" form:"ids"`
}

type PageStrut struct {
	PageSize int `json:"pageSize"`
	PageNum  int `json:"pageNum"`
}
