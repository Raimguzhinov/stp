package main

import "fmt"

// HistoryRecord — одна запись в истории вычислений.
type HistoryRecord struct {
	Operand1 Number
	Operand2 Number
	Operator string
	Result   string
}

type History struct {
	records []HistoryRecord
}

func NewHistory() *History {
	return &History{records: make([]HistoryRecord, 0)}
}

func (h *History) Add(op1 Number, op string, op2 Number, result string) {
	h.records = append(h.records, HistoryRecord{
		Operand1: op1,
		Operand2: op2,
		Operator: op,
		Result:   result,
	})
}

func (h *History) Delete(index int) {
	if index >= 0 && index < len(h.records) {
		h.records = append(h.records[:index], h.records[index+1:]...)
	}
}

func (h *History) Strings() []string {
	var res []string
	for _, r := range h.records {
		line := fmt.Sprintf("%s %s %s = %s", r.Operand1, r.Operator, r.Operand2, r.Result)
		res = append(res, line)
	}
	return res
}
