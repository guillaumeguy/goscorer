package main

import (
	"encoding/xml"
	"fmt"
	"github.com/asafschers/goscore"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
)

func main() {
	log.Println("Hello World")
	route := gin.Default()

	// Load model
	modelXml, _ := ioutil.ReadFile("static/iris.pmml")
	var model goscore.RandomForest // or goscore.GradientBoostedModel
	xml.Unmarshal([]byte(modelXml), &model)

	route.GET("/testing", func(c *gin.Context) { startPage(c, &model) })
	route.Run(":8085")
}

func startPage(c *gin.Context, model *goscore.RandomForest) {

	var input Input
	if c.BindJSON(&input) == nil {
		log.Println(input.Sepal_length_cm)
		log.Println(input.Sepal_width_cm)
		log.Println(input.Petal_length_cm)
		log.Println(input.Petal_width_cm)
		log.Println("Binding success...............")
		score := ScoreInput(&input, model)
		score_str := fmt.Sprintf("%f", score)
		c.String(200, score_str)
	} else {
		log.Println("Binding failed...............")
		c.String(200, "Failure")
	}

}
