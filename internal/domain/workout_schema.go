package domain

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Log struct {
	gorm.Model

	Weight    float64
	Type      string
	Exercices []Exercice
}

type Exercice struct {
	gorm.Model

	LogID   uint
	Name    string
	Goal    string
	Reps    pq.Int64Array   `gorm:"type:integer[]"`
	Weights pq.Float64Array `gorm:"type:float[]"`
}

type NewLogInput struct {
	Weight   float64
	Type     string
	Exercice []Exercice
}
