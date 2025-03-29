package main

import "fmt"

type Record struct {
	P1, P2    int
	NumberIn  string
	NumberOut string
}

func (r Record) String() string {
	return fmt.Sprintf("p1: %d, число: %s → p2: %d, результат: %s\n", r.P1, r.NumberIn, r.P2, r.NumberOut)
}

type History struct {
	records []Record
}

func NewHistory() *History {
	return &History{records: []Record{}}
}

func (h *History) AddRecord(p1, p2 int, n1, n2 string) {
	h.records = append(h.records, Record{p1, p2, n1, n2})
}

func (h *History) Count() int {
	return len(h.records)
}

func (h *History) Get(index int) Record {
	if index >= 0 && index < len(h.records) {
		return h.records[index]
	}
	return Record{}
}
