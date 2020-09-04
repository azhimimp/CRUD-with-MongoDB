package main

import (
	"context"

	"github.com/azhimimp/mongo-crud/database"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func main() {
	router := gin.Default()

	user := router.Group("/user")
	{
		user.GET("/", Get)
		user.GET("/:id", GetByID)
		user.POST("/create", Create)
		user.PUT("/:id/", Update)
		user.DELETE("/:id", Delete)
	}

	router.Run(":8080")
}

func Get(c *gin.Context) {
	client, err := database.Mongodb()
	if err != nil {
		c.String(500, err.Error())
		return
	}

	result, err := client.Database("user").Collection("user").Find(context.Background(), bson.M{})
	if err != nil {
		c.String(500, err.Error())
		return
	}

	var data []map[string]interface{}
	result.All(context.Background(), &data)

	c.JSON(200, data)
}

func GetByID(c *gin.Context) {
	id := c.Param("id")

	client, err := database.Mongodb()
	if err != nil {
		c.String(500, err.Error())
		return
	}

	result, err := client.Database("user").Collection("user").Find(context.Background(), bson.M{"username": id})
	if err != nil {
		c.String(500, err.Error())
		return
	}

	var data []map[string]interface{}
	result.All(context.Background(), &data)

	c.JSON(200, data)
}

func Create(c *gin.Context) {
	var user User
	c.BindJSON(&user)

	client, err := database.Mongodb()
	if err != nil {
		c.String(500, err.Error())
		return
	}

	_, err = client.Database("user").Collection("user").InsertOne(context.Background(), user)
	if err != nil {
		c.String(500, err.Error())
		return
	}

	c.String(200, "success")
}

func Update(c *gin.Context) {
	var user User
	c.BindJSON(&user)

	client, err := database.Mongodb()
	if err != nil {
		c.String(500, err.Error())
		return
	}

	_, err = client.Database("user").Collection("user").UpdateOne(context.Background(), bson.M{"username": user.Username}, bson.M{"$set": user})
	if err != nil {
		c.String(500, err.Error())
		return
	}

	c.String(200, "success")
}

func Delete(c *gin.Context) {
	id := c.Param("id")

	client, err := database.Mongodb()
	if err != nil {
		c.String(500, err.Error())
		return
	}

	_, err = client.Database("user").Collection("user").DeleteOne(context.Background(), bson.M{"username": id})
	if err != nil {
		c.String(500, err.Error())
		return
	}

	c.String(200, "success")
}
