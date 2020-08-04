package request

type GetById struct {
	Id float64 `json:"id" form:"id"`
}

type GetIds struct {
	Ids []int `json:"ids" form:"ids"`
}
