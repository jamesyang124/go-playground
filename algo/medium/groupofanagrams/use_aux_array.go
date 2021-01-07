package groupofanagrams

// UseAuxArray ...
//
// T O(n * w * log w + n * log n * w), n is number of words, w is word longest length
// w * n * log n => each sorted words, now need to sort again. take n * log n
//               => in each iteration we need to compare w times for each char comparison
//
// S O(n * w)
func UseAuxArray(words []string) [][]string {
	output := [][]string{}

	return output
}
