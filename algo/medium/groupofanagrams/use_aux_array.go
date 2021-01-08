package groupofanagrams

import "sort"

// UseAuxArray ...
// https://www.algoexpert.io/questions/Group%20Anagrams
//
// T O(n * w * log w + n * log n * w), n is number of words, w is word longest length
//
// w * n * log n => each sorted words, now need to sort again. take n * log n
//               => in each iteration we need to compare w times for each char comparison
//
// S O(n * w)
func UseAuxArray(words []string) [][]string {
	output := [][]string{}

	if len(words) == 0 {
		return output
	}

	sortedWords := []string{}
	indices := []int{}

	for i, w := range words {
		s := sortWord(w)
		sortedWords = append(sortedWords, s)
		indices = append(indices, i)
	}

	sort.Slice(indices, func(i, j int) bool {
		return sortedWords[indices[i]] < sortedWords[indices[j]]
	})

	currentList := []string{}
	current := sortedWords[indices[0]]
	for _, i := range indices {
		if current == sortedWords[i] {
			currentList = append(currentList, words[i])
			continue
		}

		output = append(output, currentList)
		currentList = []string{words[i]}
		current = sortedWords[i]
	}

	output = append(output, currentList)
	return output
}
