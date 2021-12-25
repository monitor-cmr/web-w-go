package main

import "github.com/monitor-cmr/web-w-go/router"

func main() {
    // Get gin.Engine from router/route.go
    route := router.Router()

    // Run App
    route.Run(":9080")
}
