package datastructures

import (
	"math"
)

const RandomPrimeNumber = 31
const HashTableLength = 13

func hash(key string) int {
	minIteration := int(math.Min(float64(len(key)), 100))
	hashKey := 0
	for i := 0; i < minIteration; i++ {
		charAt := rune(key[i])
		hashKey = (hashKey*RandomPrimeNumber + int(charAt)) % HashTableLength
	}
	return hashKey
}

func (h *HashTable) Set(key string, value string) {
	keyHash := hash(key)
	valuesAt := &h.Values[keyHash]
	valueAlreadyExists := false
	for _, values := range *valuesAt { //check for existing key
		if values[0] == key {
			values[1] = value
			valueAlreadyExists = true
			break
		}
	}
	if !valueAlreadyExists {
		h.Values[keyHash] = append(h.Values[keyHash], []string{key, value})
	}
}

func (h *HashTable) Get(key string) string {
	keyHash := hash(key)
	valuesAt := &h.Values[keyHash]
	value := ""
	if len(*valuesAt) != 0 {
		for _, values := range *valuesAt {
			if values[0] == key {
				value = values[1]
				break
			}
		}
	}
	return value
}

func (h *HashTable) Keys() []string {
	keys := []string{}
	for _, valuesAt := range h.Values {
		for _, values := range valuesAt {
			keys = append(keys, values[0])
		}
	}
	return keys
}

func (h *HashTable) KeyValues() []string {
	keyValues := []string{}
	for _, valuesAt := range h.Values {
		for _, values := range valuesAt {
			keyValues = append(keyValues, values[1])
		}
	}
	return keyValues
}
