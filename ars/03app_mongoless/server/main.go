package main

import (
	//"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Flight struct {
	Id          string
	Number      string
	ArilineName string
	Source      string
	Destination string
	Capacity    int
	Price       float32
}

func readAllFlights(c *gin.Context) {
	flights := []Flight{
		{Id: "1001", Number: "AI 845",
			ArilineName: "Air India", Source: "Mumbai",
			Destination: "Abu Dhabi", Capacity: 180,
			Price: 15000.0},
		{Id: "1002", Number: "AI 846",
			ArilineName: "Air India", Source: "Abu Dhabi",
			Destination: "Mumbai", Capacity: 180,
			Price: 15000.0},
	}
	c.JSON(http.StatusOK, flights)
}

func readFlightById(c *gin.Context) {
	id := c.Param("id")
	flight := Flight {Id: id, Number: "AI 845",
			ArilineName: "Air India", Source: "Mumbai",
			Destination: "Abu Dhabi", Capacity: 180,
			Price: 15000.0}
	c.JSON(http.StatusOK, flight)
}

func createFlight(c *gin.Context) {
	var jbodyFlight Flight 
	err := c.BindJSON(&jbodyFlight)
	if err != nil {
		c.JSON(http.StatusInternalServerError, 
			gin.H{"error" : "Server Error." + err.Error()})
		return 
	}
	createdFlight := Flight {Id: "1001", Number: "AI 845",
				ArilineName: "Air India", Source: "Mumbai",
				Destination: "Abu Dhabi", Capacity: 180, Price: 15000.0}
	c.JSON(http.StatusCreated, 
				gin.H{"message" : "Flight Created Successfully.",
			         "flight" : createdFlight})
}

func updateFlight(c *gin.Context) {
	id := c.Param("id")//
	var jbodyFlight Flight 
	err := c.BindJSON(&jbodyFlight)
	if err != nil {
		c.JSON(http.StatusInternalServerError, 
			gin.H{"error" : "Server Error." + err.Error()})
		return 
	}
	updatedFlight := Flight {Id: id, Number: "AI 845",//
				ArilineName: "Air India", Source: "Mumbai",
				Destination: "Abu Dhabi", Capacity: 180, Price: 15000.0}
	c.JSON(http.StatusOK, 
				gin.H{"message" : "Flight Updated Successfully.",//
			         "flight" : updatedFlight})//
}

func deleteFlight(c *gin.Context) {
	//id := c.Param("id")
	c.JSON(http.StatusOK, 
				gin.H{"message" : "Flight Deleted Successfully."})
}
func main() {
	//router
	r := gin.Default()
	//routes | API Endpoints
	r.GET("/flights", readAllFlights)
	r.GET("/flights/:id", readFlightById)
	r.POST("/flights", createFlight)
	r.PUT("/flights/:id", updateFlight)
	r.DELETE("/flights/:id", deleteFlight)
	//server 
	r.Run(":8080")	
}
