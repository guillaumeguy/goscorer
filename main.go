package main

import (
	"encoding/xml"
	"fmt"
	"github.com/asafschers/goscore"
	"github.com/gin-gonic/gin"
	"goscorer/scorer"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	route := gin.Default()
	route.Use(gin.Logger())
	route.Static("/static", "static")

	port, _ := os.LookupEnv("PORT")

	if port == "" {
		port = "8000"
	}

	log.Println(port)

	// Load model
	modelXml, _ := ioutil.ReadFile("static/iris.pmml")
	var model goscore.RandomForest // or goscore.GradientBoostedModel
	xml.Unmarshal([]byte(modelXml), &model)

	route.GET("/home", helloPage)

	route.GET("/testing", func(c *gin.Context) { startPage(c, &model) })

	route.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	route.Run(":" + port)
}

func helloPage(c *gin.Context) {

	c.String(http.StatusOK, "hello world")

}

func startPage(c *gin.Context, model *goscore.RandomForest) {

	var input scorer.ScoringInput
	if c.BindJSON(&input) == nil {

		score := scorer.ScoreInput(&input, model)
		score_str := fmt.Sprintf("%f", score)
		
		c.String(200, score_str)
	} else {

		c.String(200, "Failure")
	}

}
