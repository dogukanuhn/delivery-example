package main

import (
	"os"

	"github.com/dogukanuhn/delivery-system/cmd"
)

// @title Delivery System Fleet Management Project
// @version 1.0
// @description This project can handle delivery system for Branch, Distribution and Transfer centers.
// @termsOfService http://swagger.io/terms/

// @contact.name Berkay Dogukan Urhan
// @contact.url https://www.linkedin.com/in/berkay-dogukan-urhan/
// @contact.email b.dogukanurhan@gmail.com

// @host localhost:3000
// @BasePath /
// @schemes http
func main() {
	cmd.Execute(os.Getenv("APP_ENV"))
}

//<3 END
