package figuras

type Rectangulo struct {
	Base   float64
	Altura float64
}

func (rectangulo *Rectangulo) area() float64 {
	return rectangulo.Base * rectangulo.Altura
}

func (rectangulo *Rectangulo) perimetro() float64 {
	return 2 * (rectangulo.Base + rectangulo.Altura)
}
