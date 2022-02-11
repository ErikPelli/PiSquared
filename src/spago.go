// Copyright (c) 2022 Erik Pellizzon
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package src

import (
	"github.com/nlpodyssey/spago/pkg/mat32"
	"github.com/nlpodyssey/spago/pkg/nlp/transformers/bert"
	"log"
)

var model *bert.Model

// LoadModel loads NPL BERT model in memory
func LoadModel(folder string) error {
	var err error
	model, err = bert.LoadModel(folder)
	return err
}

// CloseModel closes BERT model
func CloseModel() {
	model.Close()
}

func vector(text string) []float32 {
	vector, err := model.Vectorize(text, bert.ClsToken)
	if err != nil {
		log.Fatal(err)
	}
	return vector.(*mat32.Dense).Normalize2().Data()
}

func dotProduct(v1, v2 []float32) float32 {
	s := float32(0)
	_ = v2[len(v1)-1] // avoid bounds check
	for i, a := range v1 {
		s += a * v2[i]
	}
	return s
}
