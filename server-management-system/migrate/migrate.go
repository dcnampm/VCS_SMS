package main

import (
	"fmt"
	"log"

	"github.com/dcnampm/VCS_SMS.git/initializers"
	"github.com/dcnampm/VCS_SMS.git/models"
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal(" Could not load environment variables", err)
	}

	initializers.ConnectDB(&config)
}

func main() {
	initializers.DB.AutoMigrate(&models.User{})
	fmt.Println(" Migration complete")
}
