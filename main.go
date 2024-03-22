package main

import (
	"final-assignment/database"
	"final-assignment/routers"
	"os"
	"fmt"
)

//	@title			MyGram api
//	@version		1.0
//	@description	This is a API to reach out my final assignment
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		localhost:8080
//	@BasePath	/

// @securityDefinitions.apikey	ApiKeyAuth
// @in							header
// @name						Authorization
// @description				Description for what is this security definition being used
func main() {
	fmt.Println("This is port of postgres", os.Getenv("PGPORT")
	database.PostGresDB()
	routers.Router()
}
