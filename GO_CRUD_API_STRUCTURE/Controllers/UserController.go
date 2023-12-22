package controllers

import (
	models "api/Models"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Clinet *mongo.Client

func init(){

	client_option := options.Client().ApplyURI("mongodb://localhost:27017")
	Clinet,_ = mongo.Connect(context.Background(),client_option)

}

////////////////////////////////////////////////////////////////////////////////

func Welcome(c*gin.Context){

	data:=models.Welcome{
		Message: "WELCOME TO API.....",
	}

	c.JSON(http.StatusOK,data)

}

////////////////////////////////////////////////////////////////////////////////

func CreateNewUser(c *gin.Context){

	var user models.User

	err:= c.BindJSON(&user)
	if err != nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"error":err.Error(),
		})
		return
	}else{
        createUser(user)
	}
}

func createUser(user models.User){

	collection := Clinet.Database("test_demo").Collection("users")

	
	_,_ = collection.InsertOne(context.Background(),user)

}


////////////////////////////////////////////////////////////////////////////////

func GetAllUsers(c *gin.Context){

	collection := Clinet.Database("test_demo").Collection("users")

	cursor ,err := collection.Find(context.Background(),bson.M{})

	if err != nil{
		c.JSON(http.StatusInternalServerError,gin.H{
			"error":err.Error(),
		})
		return
	}

	defer cursor.Close(context.Background())

	var users []models.User

	for cursor.Next(context.Background()){
		var user models.User
		if err := cursor.Decode(&user);err!=nil{
			c.JSON(http.StatusInternalServerError,gin.H{
				"error":err.Error(),
			})
			return
		}
		users = append(users, user)
	}

	if err:= cursor.Err();err!=nil{
		c.JSON(http.StatusConflict,gin.H{
			"error":"error while itering",
		})
	}

	c.JSON(http.StatusOK,users)

}

////////////////////////////////////////////////////////////////////////////////

func UpdateUser(c *gin.Context){

	var User models.User

	UserId := c.Param("id")
	objectId,_ := primitive.ObjectIDFromHex(UserId)

	if err:= c.BindJSON(&User);err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{
			"error":err.Error(),
		})
		return
	}

	userupdate(objectId,User)

	c.JSON(http.StatusOK,gin.H{
		"message":"SUCCESS",
	})
}

func userupdate(userId primitive.ObjectID,User models.User){

	collection := Clinet.Database("test_demo").Collection("users")

	filter := bson.M{"_id":userId}
	update := bson.M{"$set":User}

	_,_ = collection.UpdateOne(context.Background(),filter,update)
}

////////////////////////////////////////////////////////////////////////////////

func DeleteUser(c *gin.Context){

	userId := c.Param("id")
	objectId,err := primitive.ObjectIDFromHex(userId)

	if err != nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"error":err.Error(),
		})
	}

    deleteuser(objectId)

	c.JSON(http.StatusOK,gin.H{
		"message":"SUCCESS DELETE",
	})

}

func deleteuser(userid primitive.ObjectID){

	collection:= Clinet.Database("test_demo").Collection("users")

	filter:= bson.M{"_id":userid}

   _,_ = collection.DeleteOne(context.Background(),filter)


}