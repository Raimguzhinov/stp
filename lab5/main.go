package main

import (
	"errors"
	"math"
	"strconv"
)

type TPNumber struct {
	value     float64
	base      int
	precision int
}

func (t *TPNumber) roundToDecimalPlace() *TPNumber {
	multiplier := math.Pow10(t.precision)
	t.value = math.Round(t.value*multiplier) / multiplier
	return t
}

// 1. Конструктор числа
func NewTPNumber(a float64, b, c int) (*TPNumber, error) {
	if b < 2 || b > 16 {
		return nil, errors.New("основание должно быть в пределах от 2 до 16")
	}
	if c < 0 {
		return nil, errors.New("точность не может быть отрицательной")
	}
	num := &TPNumber{value: a, base: b, precision: c}
	return num.roundToDecimalPlace(), nil
}

// 2. Конструктор из строки
func NewTPNumberFromString(a, bs, cs string) (*TPNumber, error) {
	b, err := strconv.Atoi(bs)
	if err != nil || b < 2 || b > 16 {
		return nil, err
	}
	c, err := strconv.Atoi(cs)
	if err != nil || c < 0 {
		return nil, err
	}
	n, err := strconv.ParseFloat(a, 64)
	if err != nil {
		return nil, err
	}
	return NewTPNumber(n, b, c)
}

// 3. Копировать
func (t *TPNumber) Copy() *TPNumber {
	return &TPNumber{value: t.value, base: t.base, precision: t.precision}
}

// 4. Сложить
func (t *TPNumber) Add(d *TPNumber) (*TPNumber, error) {
	if t.base != d.base || t.precision != d.precision {
		return nil, errors.New("разные основания или точности")
	}
	return NewTPNumber(t.value+d.value, t.base, t.precision)
}

// 5. Умножить
func (t *TPNumber) Mul(d *TPNumber) (*TPNumber, error) {
	if t.base != d.base || t.precision != d.precision {
		return nil, errors.New("разные основания или точности")
	}
	return NewTPNumber(t.value*d.value, t.base, t.precision)
}

// 6. Вычесть
func (t *TPNumber) Sub(d *TPNumber) (*TPNumber, error) {
	if t.base != d.base || t.precision != d.precision {
		return nil, errors.New("разные основания или точности")
	}
	return NewTPNumber(t.value-d.value, t.base, t.precision)
}

// 7. Делить
func (t *TPNumber) Div(d *TPNumber) (*TPNumber, error) {
	if d.value == 0 {
		return nil, errors.New("нельзя делить на 0")
	}
	if t.base != d.base || t.precision != d.precision {
		return nil, errors.New("разные основания или точности")
	}
	return NewTPNumber(t.value/d.value, t.base, t.precision)
}

// 8. Обратить
func (t *TPNumber) Inverse() (*TPNumber, error) {
	if t.value == 0 {
		return nil, errors.New("нельзя делить на 0")
	}
	return NewTPNumber(1/t.value, t.base, t.precision)
}

// 9. Квадрат
func (t *TPNumber) Square() (*TPNumber, error) {
	return NewTPNumber(t.value*t.value, t.base, t.precision)
}

// 10. Взять значение числа
func (t *TPNumber) GetValue() float64 {
	return t.value
}

func intToBase(value int, base int) string {
	const chars = "0123456789ABCDEF"
	result := ""

	if value == 0 {
		return "0"
	}

	for value > 0 {
		result = string(chars[value%base]) + result
		value /= base
	}
	return result
}

func fracToBase(value float64, base int, precision int) string {
	const chars = "0123456789ABCDEF"
	result := ""
	frac := value

	for i := 0; i < precision; i++ {
		frac *= float64(base)
		digit := int(frac)
		result += string(chars[digit])
		frac -= float64(digit)
	}
	return result
}

func (t *TPNumber) GetValueString() string {
	// Проверка на корректность базы (системы счисления)
	if t.base < 2 || t.base > 36 {
		return "Ошибка: неподдерживаемая система счисления"
	}
	// Разделение числа на целую и дробную части
	intPart, fracPart := math.Modf(t.value)

	// Преобразование целой части
	intString := strconv.FormatInt(int64(intPart), t.base)

	// Преобразование дробной части
	fracString := ""
	if t.precision > 0 {
		fracString = "."
		for i := 0; i < t.precision; i++ {
			fracPart *= float64(t.base)
			digit := int(fracPart)
			fracString += strconv.FormatInt(int64(digit), t.base)
			fracPart -= float64(digit)

			// Если дробная часть стала нулем, можно прекратить
			if math.Abs(fracPart) < 1e-10 {
				break
			}
		}
	}

	return intString + fracString
}

// 11. Возвращает значение в виде строки в системе счисления b
//func (t *TPNumber) GetValueString() string {
//	//intPart := int64(t.value)
//	intPart, fracPart := math.Modf(t.value)
//	//fracPart := t.value - float64(intPart)
//	intStr := intToBase(int(intPart), t.base)
//	if t.precision > 0 {
//		fracStr := fracToBase(fracPart, t.base, t.precision)
//		return fmt.Sprintf("%s.%s", intStr, fracStr)
//	}
//	return intStr
//}

// 12. Взять основание числа
func (t *TPNumber) GetBase() int {
	return t.base
}

// 13. Взять основание в виде строки
func (t *TPNumber) GetBaseString() string {
	return strconv.Itoa(t.base)
}

// 14. Взять точность числа
func (t *TPNumber) GetPrecision() int {
	return t.precision
}

// 15. Взять точность в виде строки
func (t *TPNumber) GetPrecisionString() string {
	return strconv.Itoa(t.precision)
}

// 16. Установить основание (число)
func (t *TPNumber) SetBase(newb int) error {
	if newb < 2 || newb > 16 {
		return errors.New("основание должно быть в пределах от 2 до 16")
	}
	t.base = newb
	return nil
}

// 17. Установить основание (строка)
func (t *TPNumber) SetBaseFromString(bs string) error {
	newb, err := strconv.Atoi(bs)
	if err != nil || newb < 2 || newb > 16 {
		return errors.New("основание должно быть в пределах от 2 до 16")
	}
	t.base = newb
	return nil
}

// 18. Установить точность (число)
func (t *TPNumber) SetPrecision(newc int) error {
	if newc < 0 {
		return errors.New("точность должна быть неотрицательной")
	}
	t.precision = newc
	return nil
}

// 19. Установить точность (строка)
func (t *TPNumber) SetPrecisionFromString(newc string) error {
	precision, err := strconv.Atoi(newc)
	if err != nil || precision < 0 {
		return errors.New("точность должна быть неотрицательной")
	}
	t.precision = precision
	return nil
}

func main() {
	return
}
