package migrate

import (
	"example.com/m/initializers"
	"example.com/m/models"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.User{})
}
