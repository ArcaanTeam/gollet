package main

import (
	"gollet/config"
	"gollet/db"
)

func main() {
	config.LoadConfig()
	db.InitDB()
	db.Migrate()
}
