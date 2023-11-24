package config

type DatabaseConfig struct {
	Driver                   string
	Host                     string
	Port                     int
	User                     string
	Password                 string
	DBName                   string
	SSLMode                  string
	Url                      string
	ConnMaxLifetimeInMinutes int
	MaxOppenConns            int
	MaxIdleConns             int
}

type HttpServerConfig struct {
	Port int
}
