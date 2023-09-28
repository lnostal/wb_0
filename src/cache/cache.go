package cache

import (
	"L0/src/db"
	"errors"
)

var cache [][]byte

func New() [][]byte {
	if len(cache) == 0 {

		dataFromDb, err := db.GetData()

		if err != nil {
			return cache
		}

		for _, d := range dataFromDb {
			Add(d)
		}

		return cache
	}

	return cache
}

func Add(element []byte) {
	cache = append(cache, element)
}

func AtIndex(i int) ([]byte, error) {
	if len(cache)-1 < i {
		return nil, errors.New("out of index")
	}

	return cache[i], nil
}
