package uset

// TSet представляет множество с элементами типа T.
type TSet[T comparable] struct {
	elements map[T]struct{} // Используем map для хранения уникальных элементов.
}

// NewSet создает новое пустое множество.
func NewSet[T comparable]() *TSet[T] {
	return &TSet[T]{elements: make(map[T]struct{})}
}

// Clear очищает множество.
func (s *TSet[T]) Clear() {
	s.elements = make(map[T]struct{})
}

// Add добавляет элемент в множество.
func (s *TSet[T]) Add(element T) {
	s.elements[element] = struct{}{}
}

// Remove удаляет элемент из множества.
func (s *TSet[T]) Remove(element T) {
	delete(s.elements, element)
}

// IsEmpty проверяет, пустое ли множество.
func (s *TSet[T]) IsEmpty() bool {
	return len(s.elements) == 0
}

// Contains проверяет, принадлежит ли элемент множеству.
func (s *TSet[T]) Contains(element T) bool {
	_, exists := s.elements[element]
	return exists
}

// Union возвращает объединение двух множеств.
func (s *TSet[T]) Union(other *TSet[T]) *TSet[T] {
	result := NewSet[T]()
	for elem := range s.elements {
		result.Add(elem)
	}
	for elem := range other.elements {
		result.Add(elem)
	}
	return result
}

// Difference возвращает разность двух множеств.
func (s *TSet[T]) Difference(other *TSet[T]) *TSet[T] {
	result := NewSet[T]()
	for elem := range s.elements {
		if !other.Contains(elem) {
			result.Add(elem)
		}
	}
	return result
}

// Intersection возвращает пересечение двух множеств.
func (s *TSet[T]) Intersection(other *TSet[T]) *TSet[T] {
	result := NewSet[T]()
	for elem := range s.elements {
		if other.Contains(elem) {
			result.Add(elem)
		}
	}
	return result
}

// Size возвращает количество элементов в множестве.
func (s *TSet[T]) Size() int {
	return len(s.elements)
}

// Elements возвращает срез всех элементов множества.
func (s *TSet[T]) Elements() []T {
	keys := make([]T, 0, len(s.elements))
	for elem := range s.elements {
		keys = append(keys, elem)
	}
	return keys
}
