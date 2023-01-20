package main

import (
	"bwastartup2/handler"
	"bwastartup2/user"
	"log"

	// "github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=postgres password=Admin1234% dbname=bwastartup port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	// fmt.Println("Connection to database is good")

	// var users []user.User
	// db.Find(&users)

	// for _, user := range users {
	// 	fmt.Println(user.Name)
	// 	fmt.Println(user.Email)
	// 	fmt.Println("========")
	// }
	
	// router := gin.Default()
	// router.GET("/", handler)
	// router.Run()

	// input user
		userRepository := user.NewRepository(db)
		userServices := user.NewService(userRepository)

		userHandler := handler.NewUserHandler(userServices)

		router := gin.Default()
		api := router.Group("/api/v1")

		api.POST("/users", userHandler.RegisterUser)

		router.Run()

		// old input ya ini kan input by code bukan by user makanya harus di ganti

		// userInput := user.RegisterUserInput{}
		// userInput.Name = "Tes Simpan dari service"
		// userInput.Email = "anjay@gmail.com"
		// userInput.Occupation = "anjay apaan tuh"
		// userInput.Password = "password"

		// userServices.RegisterUser(userInput)

		// user := user.User{
		// 	Name: "Test Simpan",
		// }
	
		// userRepository.Save(user)

		// ini ring road atau alur apinya
		// input dari user
		// handler, mapping dari input user -> struct input
		// service => melakukan mapping dari struct input ke struct User
		// repository
		// db
	
}

// func handler(c *gin.Context)  {
// 	dsn := "host=localhost user=postgres password=Admin1234% dbname=bwastartup port=5432 sslmode=disable TimeZone=Asia/Jakarta"
// 	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

// 	if err != nil {
// 		log.Fatal(err.Error())
// 	}
// }