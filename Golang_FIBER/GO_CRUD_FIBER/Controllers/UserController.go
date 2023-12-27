package controllers

import (
	models "api/Models"
	"context"
	"log"
	"sync"
	// "time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	 
	usermutex sync.Mutex
)

var Client *mongo.Client

func init(){
	
	Clinet_Option:= options.Client().ApplyURI("mongodb://localhost:27017")
	// Clinet_Option.SetTimeout(5*time.Second)
	Client,_ = mongo.Connect(context.Background(),Clinet_Option)

	if err:=Client.Ping(context.Background(),nil);err!=nil{
		log.Println("ERROR WHILE CONNECTING TO DB.....")
		log.Fatal(err)
	}

	log.Println("SUCCESS : DB CONNECTED")

}

func Welcome(c *fiber.Ctx)error{

	return c.JSON(fiber.Map{
		"message":"SHREE GANESH....",
	})

}

func RegisterUser(c *fiber.Ctx)error{

	var user models.User

	if err:= c.BodyParser(&user);err!=nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":err.Error(),
		})
	}

	err:= InsertUser(user)

	if err!=nil{
		return c.JSON(fiber.Map{
			"error":err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message":"SUCESS",
	})

}

func InsertUser(user models.User)error{

	collection := Client.Database("test_demo").Collection("users")

	_,err:= collection.InsertOne(context.Background(),user)

	return err

}


func GetAllUsers(c *fiber.Ctx)error{

	collection:= Client.Database("test_demo").Collection("users")

	cursor,err := collection.Find(context.Background(),bson.M{})

	if err!= nil{
		return c.JSON(fiber.Map{
			"error":err.Error(),
		})
	}

	defer cursor.Close(context.Background())
	
	var users [] models.User

	if err:= cursor.All(context.Background(),&users);err!=nil{
		return c.JSON(fiber.Map{
			"error":err.Error(),
		})
	}
	return c.JSON(users)

}

func UpdateUser(c *fiber.Ctx)error{

	User_Id := c.Params("id")

	ObjectId,_:= primitive.ObjectIDFromHex(User_Id)

	var user models.User

	if err:= c.BodyParser(&user);err!=nil{
		return c.JSON(fiber.Map{
			"error":err.Error(),
		})
	}

	userupdate(ObjectId,user)


	return c.JSON("SUCCESS")

}

func userupdate(userid primitive.ObjectID,user models.User){

	collection := Client.Database("test_demo").Collection("users")

	filter := bson.M{"_id":userid}
	update := bson.M{"$set":bson.M{}}

	if user.Name != ""{
		update["$set"].(bson.M)["name"] = user.Name 
	}
	if user.Email !=""{
		update["$set"].(bson.M)["email"] = user.Email
	}
	if user.Location !=""{
		update["$set"].(bson.M)["location"] = user.Location
	}


	_,_=collection.UpdateOne(context.Background(),filter,update)

	 

}


func Deleteuser(c* fiber.Ctx)error{

	userId:= c.Params("id")

	id,_ := primitive.ObjectIDFromHex(userId)

	deleteuser(id)

	return c.JSON(fiber.Map{
		"message":"success",
	})
}

func deleteuser(objectid primitive.ObjectID){
	collection:= Client.Database("test_demo").Collection("users")

	filter:= bson.M{"_id":objectid}
	_,_= collection.DeleteOne(context.Background(),filter)
}