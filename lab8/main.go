package main

import (
	"fmt"
	"log"
	"stp/lab8/uproc"
)

func main() {
	proc := uproc.NewTProc(0.0)
	// Установка правого операнда: 4/1
	proc.ROpSet(4)
	// Возведение в квадрат: (4/1)^2
	if err := proc.FuncRun(uproc.Sqr); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("(4/1)^2 = %.2f\n", proc.GetROp())
	// Установка левого операнда: 3/1
	proc.LOpAndResSet(3)
	// Умножение: 3/1 * (4/1)^2
	proc.OperationSet(uproc.Mul)
	if err := proc.OperationRun(); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("3/1 * (4/1)^2 = %.2f\n", proc.GetLOpAndRes())
	// Установка правого операнда: 2/1
	proc.ROpSet(2)
	// Сложение: 2/1 + 3/1 * (4/1)^2
	proc.OperationSet(uproc.Add)
	if err := proc.OperationRun(); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("2/1 + 3/1 * (4/1)^2 = %.2f\n", proc.GetLOpAndRes())
	// Сброс процессора
	proc.Reset()
	fmt.Println("After Reset LOpAndRes:", proc.GetLOpAndRes())
	fmt.Println("After Reset ROp:", proc.GetROp())
	fmt.Println("After Reset Operation:", proc.GetOperation())
}
