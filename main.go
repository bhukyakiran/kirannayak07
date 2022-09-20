package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type users1 struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
	City     string `json:"city"`
}

func getAllUsers1(c *gin.Context) {
	var allusers = []users1{}
	err := DB.Find(&allusers).Error
	if err != nil {
		c.IndentedJSON(http.StatusNoContent, err.Error())
		return
	}
	c.IndentedJSON(http.StatusFound, allusers)

}
func getUser1ByID(c *gin.Context) {
	var users1 users1
	id := c.Param("id")
	if id == "" {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "id cannot be empty"})
		return
	}

	if err := DB.Where("id = ?", id).First(&users1).Error; err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "could not find the requested user"})
		return
	}
	c.IndentedJSON(http.StatusFound, gin.H{"User found :)": users1})

}

func createNewUser(c *gin.Context) {
	var newuser users1

	if err := c.ShouldBindJSON(&newuser); err != nil {
		c.IndentedJSON(http.StatusUnprocessableEntity, gin.H{"error-BindJSON": err.Error()})
		return
	}

	validator := validator.New()
	if err := validator.Struct(&newuser); err != nil {
		c.IndentedJSON(http.StatusUnprocessableEntity, gin.H{"Validation-error": err.Error()})
		return
	}

	err := DB.Create(&newuser).Error
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"Error creating User :( ": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusCreated, gin.H{"New user created :)": newuser})

}

func deleteUser(c *gin.Context) {
	var userToDelete users1

	id := c.Param("id")
	if id == "" {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "id cannot be empty"})
		return
	}

	err := DB.Delete(userToDelete, id).Error
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "could not delete user", "error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Delete successful."})

}
func updateUser(c *gin.Context) {
	var updateuser users1
	id := c.Param("id")
	if id == "" {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "id cannot be empty"})
		return
	}

	if err := c.ShouldBindJSON(&updateuser); err != nil {
		c.IndentedJSON(http.StatusUnprocessableEntity, gin.H{"not valid": err.Error()})
		return
	}
	err := DB.Where("id=?", id).Updates(&updateuser).Error
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"not valid": err.Error()})
	}
}

var DB *gorm.DB

func setupDB() {
	connectionString := "host=localhost user=postgres password=Tinku@9515302551 dbname=users1 port=5432 sslmode=disable TimeZone=Asia/Kolkata"
	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})

	if err != nil {
		log.Fatal("could not connect to database")
		return
	}
	DB = db

	err = DB.AutoMigrate(&users1{})
	if err != nil {
		log.Fatal("could not migrate to db")
		return
	}
}
func main() {
	setupDB()
	router := gin.Default()
	router.GET("/allusers1", getAllUsers1)
	router.GET("/users:id", getUser1ByID)
	router.POST("/newusers", createNewUser)
	router.DELETE("/delete/:id", deleteUser)
	router.PUT("/Updateusers", updateUser)
	router.Run("localhost:8080")

}
