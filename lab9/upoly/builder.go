package upoly

// PolyBuilder используется для пошагового создания полинома.
type PolyBuilder struct {
	poly *TPoly
}

// NewPolyBuilder создает новый PolyBuilder.
func NewPolyBuilder() *PolyBuilder {
	return &PolyBuilder{
		poly: NewPoly(),
	}
}

// AddMember добавляет одночлен в полином.
func (b *PolyBuilder) AddMember(coeff, degree int) *PolyBuilder {
	b.poly.AddMember(NewMember(coeff, degree))
	return b
}

// Build возвращает готовый полином.
func (b *PolyBuilder) Build() *TPoly {
	return b.poly
}
