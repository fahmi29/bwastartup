package main

import (
	"bwastartup2/auth"
	"bwastartup2/handler"
	"bwastartup2/user"
	"fmt"
	"log"

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

	userRepository := user.NewRepository(db)
	userServices := user.NewService(userRepository)
	authService := auth.NewService()

	// token, err := authService.ValidateToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxfQ.JJjGYA9PExsu83H51Q1NcF_fuNpXUMWu0KbkrcHW1Ac")
	// if err != nil {
	// 	fmt.Println("ERROR")
	// 	fmt.Println("ERROR")
	// 	fmt.Println("ERROR")
	// }

	// if token.Valid {
	// 	fmt.Println("VALID")
	// 	fmt.Println("VALID")
	// 	fmt.Println("VALID")
	// }else {
	// 	fmt.Println("INVALID")
	// 	fmt.Println("INVALID")
	// 	fmt.Println("INVALID")
	// }

	userHandler := handler.NewUserHandler(userServices, authService)

	router := gin.Default()
	api := router.Group("/api/v1")

	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)
	api.POST("/email_checkers", userHandler.CheckEmailAvailability)
	api.POST("/avatars", userHandler.UploadAvatar)

	router.Run()

	// manual update avatar user
	// userServices.SaveAvatar(1, "images/1-profile.png")

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

	// manual login
	// input := user.LoginInput{
	// 	Email: "prieze29@driveme.id",
	// 	Password: "passwords",
	// }

	// user, err := userServices.Login(input)
	// if err != nil {
	// 	fmt.Println("Salah nih boss")
	// 	fmt.Println(err.Error())
	// }

	// fmt.Println(user.Email)
	// fmt.Println(user.Name)

	// login user
	// userByEmail, err := userRepository.FindByEmail("prieze29@driveme.id")
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }

	// if userByEmail.ID == 0 {
	// 	fmt.Println("User tidak di temukan")
	// } else {
	// 	fmt.Println(userByEmail.Name)
	// }

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
