package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-docker/route"
	"github.com/joho/godotenv"
)

func main() {

	// ENV
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	r := gin.Default()
	route.InitRoute(r)

	var PORT string = os.Getenv("PORT")
	fmt.Println("PORT: " + PORT)
	r.Run(PORT)
}
