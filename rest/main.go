package main

import (
    "database/sql"
    "fmt"

    "github.com/monitor-cmr/web-w-go/router"

    _ "github.com/go-sql-driver/mysql"
)

func main() {
    // Mysql
    db := MysqlConn("172.17.0.2", 3306, "alochym", "alochym@123", "MYSQLTEST")
    defer db.Close()

    // MySQL test connection
    err := db.Ping()

    if err != nil {
        panic(err)
    }

    // Get gin.Engine from router/route.go
    // route := router.Router()
    route := router.Router(db)

    // Run App
    route.Run(":8080")
}

// MysqlConn ...
func MysqlConn(host string, port int64, username, password, dbname string) *sql.DB {
    dbSource := fmt.Sprintf(
        "%s:%s@tcp(%s:%d)/%s?charset=utf8",
        username,
        password,
        host,
        port,
        dbname,
    )

    dbConn, err := sql.Open("mysql", dbSource)

    if err != nil {
        fmt.Println(err)
        panic(err)
    }

    return dbConn
}

