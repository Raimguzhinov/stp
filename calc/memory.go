package main

// Memory хранит одно число и позволяет выполнять стандартные действия: MS, MR, M+, MC.
type Memory struct {
	value *FractionNumber
}

func NewMemory() *Memory {
	return &Memory{}
}

func (m *Memory) Save(val *FractionNumber) {
	m.value = val
}

func (m *Memory) Read() *FractionNumber {
	if m.value == nil {
		return nil
	}
	fcopy := &FractionNumber{numerator: m.value.Numerator(), denominator: m.value.Denominator()}
	return fcopy.shrink()
}

func (m *Memory) Add(val *FractionNumber) {
	if m.value == nil {
		m.value = val
	} else {
		m.value = m.value.Add(val)
	}
}

func (m *Memory) Clear() {
	m.value = nil
}
