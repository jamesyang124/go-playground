package groupofanagrmas

import "sort"

func sortWord(word string) string {
	bytes := []byte(word)

	sort.Slice(bytes, func(i, j int) bool {
		return bytes[i] < bytes[j]
	})

	return string(bytes)
}
