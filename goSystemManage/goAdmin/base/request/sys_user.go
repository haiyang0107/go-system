package request

//注册用户结构体
type RegisterStruct struct {
	LoginName string `json:"loginName"`
	Name      string `json:"name"`
	Image     string `json:"image"`
	Password  string `json:"password"`
}
