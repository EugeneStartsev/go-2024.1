package slice

// AppendElement добавляет элемент в конец слайса.
func AppendElement(ints []int, elem int) []int {
	return append(ints, elem)
}

// RemoveElement удаляет последний элемент из слайса. Если массив пуст, функция возвращает оригинальный пустой массив.
func RemoveElement(ints []int) []int {
	if len(ints) != 0 {
		ints = ints[:len(ints)-1]
	}

	return ints
}

// AddOneToAll увеличивает каждый элемент массива на единицу.
func AddOneToAll(ints []int) []int {
	for i, v := range ints {
		ints[i] = v + 1
	}

	return ints
}
