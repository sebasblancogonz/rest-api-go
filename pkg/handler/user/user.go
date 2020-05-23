package user

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sebasblancogonz/rest-api-go/config"
	model_user "github.com/sebasblancogonz/rest-api-go/pkg/models/user"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// UserCollection static collection
const UserCollection = "user"

// MongoConfig returns config
func MongoConfig() *mgo.Database {
	db, err := config.GetMongoDB()
	if err != nil {
		fmt.Println(err)
	}
	return db
}

// GetAllUsers will return all users
func GetAllUsers(c *gin.Context) {
	db := *MongoConfig()

	users := model_user.Users{}
	err := db.C(UserCollection).Find(bson.M{}).All(&users)

	if err != nil {
		c.JSON(500, gin.H{
			"message": "Error getting all users",
		})
		return
	}

	c.JSON(200, gin.H{
		"users": &users,
	})
}

// GetUser will return a user given an id
func GetUser(c *gin.Context) {
	db := *MongoConfig()

	id := c.Param("id")
	idParse, errParse := strconv.Atoi(id)
	if errParse != nil {
		c.JSON(500, gin.H{
			"message": "Error parsing param",
		})
		return
	}

	user := model_user.User{}
	err := db.C(UserCollection).Find(bson.M{"id": &idParse}).One(&user)

	if err != nil {
		c.JSON(500, gin.H{
			"message": "Error getting user",
		})
		return
	}

	c.JSON(200, gin.H{
		"user": &user,
	})
}

// CreateUser will create new users
func CreateUser(c *gin.Context) {
	db := *MongoConfig()

	user := model_user.User{}
	err := c.Bind(&user)

	if err != nil {
		c.JSON(500, gin.H{
			"message": "Error getting body",
		})
		return
	}

	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	err = db.C(UserCollection).Insert(user)

	if err != nil {
		c.JSON(500, gin.H{
			"message": "Error saving user",
		})

		return
	}

	c.JSON(200, gin.H{
		"message": "User saved",
		"user":    &user,
	})
}
