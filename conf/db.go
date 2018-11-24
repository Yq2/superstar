package conf

const DriverName = "mysql"

type DbConf struct {
	Host   string
	Port   int
	User   string
	Pwd    string
	DbName string
}

//主连接
var MasterDbConfig DbConf = DbConf{
	Host:   "39.107.77.94",
	Port:   3306,
	User:   "root",
	Pwd:    "123456",
	DbName: "superstar",
}
//从连接
var SlaveDbConfig DbConf = DbConf{
	Host:   "39.107.77.94",
	Port:   3306,
	User:   "root",
	Pwd:    "123456",
	DbName: "superstar",
}
