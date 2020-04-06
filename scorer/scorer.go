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

	var score = 0.0 

	return score
}
