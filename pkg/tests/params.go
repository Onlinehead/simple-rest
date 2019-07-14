package tests

import "os"

func GetPostgresTestParams() (username string, password string, db string, host string, port string) {
	username = os.Getenv("PostgresTestUsername")
	if username == "" {
		username = "user"
	}
	password = os.Getenv("PostgresTestPassword")
	if password == "" {
		password = "zzz"
	}
	db = os.Getenv("PostgresTestDB")
	if db == "" {
		db = "simple_rest"
	}
	port = os.Getenv("PostgresTestPort")
	if port == "" {
		port = "5432"
	}
	host = os.Getenv("PostgresTestHost")
	if host == "" {
		host = "192.168.244.130"
	}
	return
}