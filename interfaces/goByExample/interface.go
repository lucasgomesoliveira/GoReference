package main

	

import "fmt"
import "math"

Aqui temos uma interface básica para formas geométricas.
	

type geometria interface {
    area() float64
    perim() float64
}

Para nosso exemplo vamos implementar essa interface nos tipos quadrado e círculo.
	

type quadrado struct {
    largura, altura float64
}
type círculo struct {
    raio float64
}

Para implementar uma interface no Go, nós apenas precisamos implementar todos os métodos na interface. Aqui nós implementamos geometria nos quadrados.
	

func (q quadrado) area() float64 {
    return q.largura * s.altura
}
func (q quadrado) perim() float64 {
    return 2*q.largura + 2*q.altura
}

A implementação para círculos.
	

func (c círculo) area() float64 {
    return math.Pi * c.raio * c.raio
}
func (c círculo) perim() float64 {
    return 2 * math.Pi * c.raio
}

Se a variável possui um tipo inteface, então nós podemos chamar métodos que estão na interface nomeada. Aqui temos uma função genérica medida tomando vantagem de seu trabalho para qualquer geometria.
	

func medir(g geometria) {
    fmt.Println(g)
    fmt.Println(g.area())
    fmt.Println(g.perim())
}

	

func main() {
    q := quadrado{largura: 3, altura: 4}
    c := círculo{raio: 5}

Os tipos de estruturas círculo e quadrado ambos implementam a interface geometria, então podemos usar instâncias dessas estruturas como argumentos para medir.
	

    medir(q)
    medir(c)
}
