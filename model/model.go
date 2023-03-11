package model

type Data struct {
	Id          int    `json:"id"`
	CpuInfo     string `json:"cpu_info"`
	ProcessInfo string `json:"process_info"`
	UsersInfo   string `json:"users_info"`
	OsInfo      string `json:"os_info"`
}
