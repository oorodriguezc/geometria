package figuras

import "fmt"

type Figura interface {
	area() float64
	perimetro() float64
}

func ObtenerArea(figura Figura) {
	fmt.Println(figura.area())
}

func ObtenerPerimetro(figura Figura) {
	fmt.Println(figura.perimetro())
}
