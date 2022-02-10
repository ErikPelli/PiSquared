package src

import (
	"github.com/nlpodyssey/spago/pkg/mat32"
	"github.com/nlpodyssey/spago/pkg/nlp/transformers/bert"
	"log"
)

var model *bert.Model

// LoadModel loads NPL BERT model in memory
func LoadModel(folder string) (err error) {
	model, err = bert.LoadModel(folder)
	return
}

// CloseModel closes BERT model
func CloseModel() {
	model.Close()
}

func vectorize(text string) []float32 {
	vector, err := model.Vectorize(text, bert.ClsToken)
	if err != nil {
		log.Fatal(err)
	}
	return vector.(*mat32.Dense).Normalize2().Data()
}

func dotProduct(v1, v2 []float32) float32 {
	var s float32 = 0
	_ = v2[len(v1)-1] // avoid bounds check
	for i, a := range v1 {
		s += a * v2[i]
	}
	return s
}
