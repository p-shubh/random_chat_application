// /config/config.go
package config

import (
	"os"
)

func GetDatabaseDSN() string {
	return os.Getenv("SUPABASE_DSN")
}
