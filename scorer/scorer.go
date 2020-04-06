package scorer

import (
	"github.com/asafschers/goscore"
	"log"
)

type ScoringInput struct {
	Sepal_length_cm float64 `json:"sepal_length_cm"`
	Sepal_width_cm  float64 `json:"sepal_width_cm"`
	Petal_length_cm float64 `json:"petal_length_cm"`
	Petal_width_cm  float64 `json:"Petal_width_cm"`
}

func ScoreInput(input *ScoringInput, model *goscore.RandomForest) float64 {

	features := map[string]interface{}{
		"sepal_length_cm": input.Sepal_length_cm,
		"sepal_width_cm":  input.Sepal_width_cm,
		"petal_length_cm": input.Petal_length_cm,
		"petal_width_cm":  input.Petal_width_cm,
	}

	log.Println(features)

	score, err := model.Score(features, "1")

	println("score", score)

	if err != nil {
		log.Fatal(err)
	}
	return score

}
