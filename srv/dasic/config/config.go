package config

type AppConfig struct {
	Mysql
	Redis
}
type Mysql struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
}
type Redis struct {
	Host     string
	Port     int
	Password string
	Database int
}

type Mysqls struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
}
type Redies struct {
	Host     string
	Port     int
	Password string
	Database int
}
type Mysqlss struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
}
type Rediess struct {
	Host     string
	Port     int
	Password string
	Database int
}
