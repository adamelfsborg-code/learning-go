package util

import "fmt"

type Database interface {
	connect() string
}

type PostgreSQLAdapter struct {
	Host     string
	Port     int16
	Database string
	User     string
	Password string
}

type MongoDBAdapter struct {
	ConnectionString string
}

func (p PostgreSQLAdapter) connect() string {
	connectStr := fmt.Sprintf("http://%v.%v/name=%vdatabase=%vpassword=%v", p.Host, p.Port, p.User, p.Database, p.Password)
	return connectStr
}

func (m MongoDBAdapter) connect() string {
	connectStr := fmt.Sprintf("%v", m.ConnectionString)
	return connectStr
}

func ConnectToDb(d Database) {
	fmt.Printf("Connected to database: %v \n", d.connect())
}
