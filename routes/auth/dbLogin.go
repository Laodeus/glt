package routesAuth

type DbLogin struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type DbLoginRow struct {
	Id int `json:"id"`
	DbLogin
}
