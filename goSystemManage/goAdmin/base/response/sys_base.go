package response

type PageResult struct {
	List     interface{} `json:"list"`
	Total    int         `json:"total"`
	PageSize int         `json:"pageSize"`
	PageNum  int         `json:"pageNum"`
}
