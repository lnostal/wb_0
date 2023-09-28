package cache

import "errors"

var cache [][]byte

func New() [][]byte {
	if len(cache) == 0 {
		// загрузить данные из бд
	}

	return cache
}

func Add(element []byte) {
	cache = append(cache, element)
}

func AtIndex(i int) ([]byte, error) {
	if len(cache)-1 < i {
		return nil, errors.New("out of memory")
	}

	return cache[i], nil
}
