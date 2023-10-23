package pkg

type Log struct {
	Weight    float64
	Type      string
	Exercices []*Exercice
}

type Exercice struct {
	Name   string
	Goal   string
	Reps   []int64
	Weight []float64
}
