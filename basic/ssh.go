package main

import (
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	//goMysql "github.com/go-sql-driver/mysql"
	//"golang.org/x/crypto/ssh"
	"gorm.io/driver/mysql"
)

func main() {
	fmt.Println('Hello, World!')
	//sshConfig := &ssh.ClientConfig{
	//	User: "root",
	//	Auth: []ssh.AuthMethod{
	//		ssh.Password("p6jYjyoOEsJXhOBp"),
	//	},
	//	HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	//}
	//
	//sshConn, err := ssh.Dial("tcp", "47.93.236.71:22", sshConfig)
	//if err != nil {
	//	panic(err)
	//}
	//
	//goMysql.RegisterDialContext("mysql+tcp", func(ctx context.Context, addr string) (net.Conn, error) {
	//	return sshConn.Dial("tcp", "remote_database_host:remote_database_port")
	//})

	//db, err := gorm.Open(mysql(dsn), &gorm.Config{})
	//if err != nil {
	//	panic(err)
	//}
	dsn:="zb_rw:pwd@tcp(localhost:3307)/zb?loc=PRC&charset=utf8mb4&parseTime=True"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: logger.Discard}) //
	if err != nil {
		panic(err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		panic("failed to get *sql.DB")
	}

	err = sqlDB.Ping()
	if err != nil {
		fmt.Println("failed to connect to database:", err)
	} else {
		fmt.Println("successfully connected to database")
	}
}