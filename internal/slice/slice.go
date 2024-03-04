package slice

// AppendElement добавляет элемент в конец слайса.
func AppendElement(ints []int, elem int) []int {
	ints = append(ints, elem)
	return ints
}

// RemoveElement удаляет последний элемент из слайса. Если массив пуст, функция возвращает оригинальный пустой массив.
func RemoveElement(ints []int) []int {
	ints = ints[:len(ints)-1]
	return ints
}

// AddOneToAll увеличивает каждый элемент массива на единицу.
func AddOneToAll(ints []int) []int {
	for i := 0; i < len(ints); i++ {
		ints[i] += 1
	}
	return ints
}
