package main

import (
	"bwastartup2/auth"
	"bwastartup2/campaign"
	"bwastartup2/handler"
	"bwastartup2/helper"
	"bwastartup2/user"
	"log"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=postgres password=Admin1234% dbname=bwastartup port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	// online db
	// dsn := "host=tiny.db.elephantsql.com user=bkmocuis password=QodwpnqegQNRyRlarmOrNkCl3ArMdfc5 dbname=bkmocuis port=5432 sslmode=disable TimeZone=Asia/Jakarta"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := user.NewRepository(db)
	campaignRepository := campaign.NewRepository(db)

	// campaigns, err := campaignRepository.FindByUserID(1)

	// fmt.Println("debug")
	// fmt.Println("debug")
	// fmt.Println("debug")
	// fmt.Println(len(campaigns))
	// for _, campaign := range campaigns {
	// 	fmt.Println(campaign.Name)
	// 	if len(campaign.CampaignImages) > 0 {
	// 		fmt.Println(campaign.CampaignImages[0].FileName)
	// 	}
	// }

	userServices := user.NewService(userRepository)
	campaignService := campaign.NewService(campaignRepository)
	authService := auth.NewService()

	userHandler := handler.NewUserHandler(userServices, authService)
	campaignHandler := handler.NewCampaignHandler(campaignService)

	router := gin.Default()
	api := router.Group("/api/v1")

	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)
	api.POST("/email_checkers", userHandler.CheckEmailAvailability)
	api.POST("/avatars", authMiddleware(authService, userServices), userHandler.UploadAvatar)

	api.GET("/campaigns", campaignHandler.GetCampaigns)

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

func authMiddleware(authService auth.Service, userService user.Service) gin.HandlerFunc {

	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if !strings.Contains(authHeader, "Bearer") {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		// authtoken => Bearer tokentokentokentoken
		tokenString := ""
		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		token, err := authService.ValidateToken(tokenString)
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)

		if !ok || !token.Valid {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		userID := int(claim["user_id"].(float64))

		user, err := userService.GetUserByID(userID)
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		c.Set("currentUser", user)

	}

}

// func handler(c *gin.Context)  {
// 	dsn := "host=localhost user=postgres password=Admin1234% dbname=bwastartup port=5432 sslmode=disable TimeZone=Asia/Jakarta"
// 	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

// 	if err != nil {
// 		log.Fatal(err.Error())
// 	}
// }

// didalam midleware ada apa saja
// ambil nilai header Authorization (harus mengirim auth di header => Bearer tokentokentoken)
// dari header Authorization, kita ambil nilai tokennya saja
// kita validasi token => menggukana auth/service.go
// dari token bisa ambil user_id
// ambil user dari db berdasarkan user_id lewat service
// kalau usernya ada kita set context isinya user
