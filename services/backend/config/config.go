package config

import "fmt"

var (
	Host     = "-"
	Username = "-"
	Password = "-"
	Database = "-"
	Port     = "-"
)

var DB_CONFIG = fmt.Sprintf(
	"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
	Host, Username, Password, Database, Port,
)
