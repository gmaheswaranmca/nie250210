package main

import (
	"context" //**
	"fmt"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"           //**
	"go.mongodb.org/mongo-driver/bson/primitive" //**
	"go.mongodb.org/mongo-driver/mongo"          //**
	"go.mongodb.org/mongo-driver/mongo/options"  //**
)

// Config
var mongoUri string = "mongodb://localhost:27017"
var mongoDbName string = "hms_app_db" //1

// Database variables
var mongoClient *mongo.Client
var departmentCollection *mongo.Collection //2

type Department struct { //3
	Id          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name        string             `json:"name" bson:"name"`
	Description string             `json:"description" bson:"description"`
}

// mongo connect
func connectToMongo() {
	var err error
	mongoClient, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoUri))
	if err != nil {
		fmt.Println("Mongo DB Connection Error!" + err.Error())
		return
	}
	name := "departments" //4
	departmentCollection = mongoClient.Database(mongoDbName).Collection(name)
}

// apis
func createDepartment(c *gin.Context) { //4.1
	var department Department //4.2
	if err := c.BindJSON(&department); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server Error." + err.Error()})
		return
	}
	//
	department.Id = primitive.NewObjectID()                              //5
	_, err := departmentCollection.InsertOne(context.TODO(), department) //6
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server Error.\n" + err.Error()})
		return
	}
	//
	c.JSON(http.StatusCreated,
		gin.H{"message": "Department Created Successfully", "department": department}) //7
}

func readAllDepartments(c *gin.Context) { //7
	//
	var departments []Department                                       //8
	cursor, err := departmentCollection.Find(context.TODO(), bson.M{}) //9
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server Error.\n" + err.Error()})
		return
	}
	defer cursor.Close(context.TODO())
	//
	departments = []Department{}                   //10
	err = cursor.All(context.TODO(), &departments) //11
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server Error.\n" + err.Error()})
		return
	}
	//
	c.JSON(http.StatusOK, departments) //12
}

func readDepartmentById(c *gin.Context) { //13
	id := c.Param("id")
	//
	departmentId, err := primitive.ObjectIDFromHex(id) //14
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID.\n" + err.Error()})
		return
	}
	//
	var department Department                                                                           //15
	err = departmentCollection.FindOne(context.TODO(), bson.M{"_id": departmentId}).Decode(&department) //16
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Department Not Found."}) //17
		return
	}
	//
	c.JSON(http.StatusOK, department) //18
}

func updateDepartment(c *gin.Context) { //19
	id := c.Param("id")
	departmentId, err := primitive.ObjectIDFromHex(id) //20
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID.\n" + err.Error()})
		return
	}
	//
	var oldDepartment Department                                                                           //21
	err = departmentCollection.FindOne(context.TODO(), bson.M{"_id": departmentId}).Decode(&oldDepartment) //22
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Department Not Found."}) //23
		return
	}
	//
	var jbodyDepartment Department     //24
	err = c.BindJSON(&jbodyDepartment) //25
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server Error." + err.Error()})
		return
	}
	oldDepartment.Name = jbodyDepartment.Name//26
	oldDepartment.Description = jbodyDepartment.Description
	//
	_, err = departmentCollection.UpdateOne(context.TODO(),//27
		bson.M{"_id": departmentId},
		bson.M{"$set": bson.M{"name": jbodyDepartment.Name,
			"description": jbodyDepartment.Description}})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server Error." + err.Error()})
		return
	}
	//response
	c.JSON(http.StatusOK, gin.H{"message": "Department Updated Successfully", "department": oldDepartment})//28
}

func deleteDepartment(c *gin.Context) {//29
	id := c.Param("id")
	departmentId, err := primitive.ObjectIDFromHex(id)//30
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID.\n" + err.Error()})
		return
	}
	//
	var oldDepartment Department//31
	err = departmentCollection.FindOne(context.TODO(), bson.M{"_id": departmentId}).Decode(&oldDepartment)//32
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Department Not Found."})//33
		return
	}
	//delete
	_, err = departmentCollection.DeleteOne(context.TODO(), bson.M{"_id": departmentId})//34
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server Error." + err.Error()})
		return
	}
	//response
	c.JSON(http.StatusOK, gin.H{"message": "Department deleted successfully."})//35
}

func main() {
	//
	connectToMongo()
	//router
	r := gin.Default()
	//cors
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // React frontend URL
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	//routes
	r.POST("/departments", createDepartment)//36
	r.GET("/departments", readAllDepartments)
	r.GET("/departments/:id", readDepartmentById)
	r.PUT("/departments/:id", updateDepartment)
	r.DELETE("/departments/:id", deleteDepartment)
	//server
	r.Run(":8080")
}
