package configs

// sqlite3
const (
	SQLITE3    = "sqlite3"
	SQLITEFILE = "./pkg/database/emoji.db"
)

// postgres
const (
	HOST     = "127.0.0.1"
	POSTGRES = "postgres"
	DATABASE = "emoji_data"
	SSLMODE  = "disable"
	USER     = "postgres"
	PASSWORD = ""
	PORT     = "5432"
)
