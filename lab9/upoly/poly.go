package upoly

import (
	"sort"
	"strings"
)

// TPoly представляет полином как список одночленов.
type TPoly struct {
	Members []*TMember // Список одночленов
}

// NewPoly создает новый пустой полином.
func NewPoly() *TPoly {
	return &TPoly{
		Members: []*TMember{},
	}
}

// AddMember добавляет одночлен в полином.
func (p *TPoly) AddMember(member *TMember) {
	if member.Coeff != 0 {
		p.Members = append(p.Members, member)
	}
	p.Normalize()
}

// Degree возвращает степень полинома.
func (p *TPoly) Degree() int {
	if len(p.Members) == 0 {
		return 0
	}
	maxDegree := 0
	for _, m := range p.Members {
		if m.Degree > maxDegree {
			maxDegree = m.Degree
		}
	}
	return maxDegree
}

// Eval вычисляет значение полинома в точке x.
func (p *TPoly) Eval(x float64) float64 {
	result := 0.0
	for _, m := range p.Members {
		result += m.Eval(x)
	}
	return result
}

// Diff возвращает производную полинома.
func (p *TPoly) Diff() *TPoly {
	result := NewPoly()
	for _, m := range p.Members {
		result.AddMember(m.Diff())
	}
	return result
}

// Normalize приводит полином к нормализованному виду.
func (p *TPoly) Normalize() {
	combined := make(map[int]int)
	for _, m := range p.Members {
		combined[m.Degree] += m.Coeff
	}
	p.Members = []*TMember{}
	for degree, coeff := range combined {
		if coeff != 0 {
			p.Members = append(p.Members, NewMember(coeff, degree))
		}
	}
}

// String возвращает строковое представление полинома.
func (p *TPoly) String() string {
	if len(p.Members) == 0 {
		return "0"
	}
	sort.Slice(p.Members, func(i, j int) bool {
		return p.Members[i].Degree < p.Members[j].Degree
	})
	var terms []string
	for _, m := range p.Members {
		terms = append(terms, m.String())
	}
	return strings.Join(terms, " + ")
}
