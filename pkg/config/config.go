package config

type Config struct {
	database *DBCreds
}

func (c *Config) Database() *DBCreds {
	return c.database
}

type DBCreds struct {
	Engine       string `json:"engine"`
	Host         string `json:"host"`
	Password     string `json:"password"`
	Username     string `json:"username"`
	Port         int    `json:"port"`
	DatabaseName string `json:"databaseName"`
}
