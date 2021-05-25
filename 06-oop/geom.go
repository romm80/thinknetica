package geom

import (
	"errors"
	"math"
)

type Geom struct {
	//X1, Y1, X2, Y2 float64
	x1, x2, y1, y2 float64 // приватные поля
}

func New(x1, x2, y1, y2 float64) (*Geom, error) {
	if x1 < 0 || x2 < 0 || y1 < 0 || y2 < 0 { // Проверка на валидность координат из метода вычисления расстояния
		return nil, errors.New("Координаты не могут быть меньше нуля")
	}
	g := Geom{x1, x2, y1, y2}
	return &g, nil
}

//func (geom Geom) CalculateDistance() (distance float64) {
func (g Geom) Distance() float64 { // g вместо geom, название метода короче, не именованнный возварат
	// возврат расстояния между точками
	//return distance
	return math.Sqrt(math.Pow(g.x2-g.x1, 2) + math.Pow(g.y2-g.y1, 2)) // возварт без промежуточной переменной
}
