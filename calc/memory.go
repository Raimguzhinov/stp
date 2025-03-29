package main

// Memory хранит одно число и позволяет выполнять стандартные действия: MS, MR, M+, MC.
type Memory struct {
	value Number
}

func NewMemory() *Memory {
	return &Memory{}
}

func (m *Memory) Save(val Number) {
	m.value = val
}

func (m *Memory) Read() Number {
	if m.value == nil {
		return nil
	}
	return m.value.Copy()
}

func (m *Memory) Add(val Number) {
	if m.value == nil {
		m.value = val
	} else {
		m.value = m.value.Add(val)
	}
}

func (m *Memory) Clear() {
	m.value = nil
}
