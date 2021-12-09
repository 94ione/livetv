package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"github.com/94ione/livetv/database"
	"github.com/94ione/livetv/handlers"
	"github.com/94ione/livetv/services/parser"
	"github.com/94ione/livetv/services/proxy"
)

func main() {
	fmt.Println("Under Development")
	err := database.Init()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	err = proxy.Init()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	parser.Init()
	mux := gin.Default()
	mux.GET("/ts", handlers.TsProxyHandler)
	mux.GET("/m3u8", handlers.M3U8ProxyHandler)
	err = mux.Run(os.Getenv("LIVETV_LISTEN"))
	if err != nil {
		fmt.Println(err.Error())
	}
}
