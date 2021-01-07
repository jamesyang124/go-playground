package groupofanagrams

import "sort"

// UseMap ...
// https://www.algoexpert.io/questions/Group%20Anagrams
// ["ab", "ba", "cde", "dec"] => [["ab", "ba"], ["cde", "dev"]]
// with map or aux aaray to memoize index of each element
//
// without map, need aux array to memo each element's index during sort.
//
// T O(n * (w * log w))
// S O(n * w), n is number of words, w is longest word length
func UseMap(words []string) [][]string {
	output := [][]string{}
	sortedMap := map[string][]string{}

	for _, word := range words {
		sorted := sortWord(word)
		sortedMap[sorted] = append(sortedMap[sorted], word)
	}

	for _, group := range sortedMap {
		output = append(output, group)
	}

	return output
}

func sortWord(word string) string {
	bytes := []byte(word)

	sort.Slice(bytes, func(i, j int) bool {
		return bytes[i] < bytes[j]
	})

	return string(bytes)
}
