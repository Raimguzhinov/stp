package main

import (
	"fmt"
	"stp/lab9/upoly"
)

func main() {
	// Создание полинома с использованием Builder
	builder := upoly.NewPolyBuilder()
	poly := builder.AddMember(3, 2).AddMember(5, 1).AddMember(2, 0).Build()
	fmt.Println("Полином:", poly)
	fmt.Println("Степень:", poly.Degree())
	fmt.Printf("Вычисление значения полинома в точке x = 2:\n\tpoly(2) = %.2f\n", poly.Eval(2))
	derivative := poly.Diff()
	fmt.Println("Производная:", derivative)
	fmt.Printf("Вычисление значения производной в точке x = 2:\n\tderivative(2) = %.2f\n", derivative.Eval(2))
}
