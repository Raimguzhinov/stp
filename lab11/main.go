package main

func main() {
	dictionarySizes := []int{16, 32, 64, 128}
	operatorsRatio := 0.5 // 50% operators, 50% operands

	for _, eta := range dictionarySizes {
		operators := int(operatorsRatio * float64(eta))
		operands := eta - operators
		compareMetrics(operators, operands)
	}
}
