package threesum

// Time limit exceed,
// since we have to check all output again, this make time wasted take 3 from n
// n^3 wasted in worst case
func checkOutput(o [][]int, a, b, c int) bool {
	set := map[int]int{a: 0, b: 0, c: 0}
	set[a] = set[a] + 1
	set[b] = set[b] + 1
	set[c] = set[c] + 1

	isUnused := true
	for _, ov := range o {
		oSet := make(map[int]int)
		oSet[ov[0]] = oSet[ov[0]] + 1
		oSet[ov[1]] = oSet[ov[1]] + 1
		oSet[ov[2]] = oSet[ov[2]] + 1

		if oSet[a] == set[a] && oSet[b] == set[b] && oSet[c] == set[c] {
			isUnused = false
		}
	}
	return isUnused
}

func threeSumOverTime(nums []int) [][]int {
	//target := 0
	var m = map[int][][]int{}
	output := [][]int{}

	for x, xv := range nums {
		m[-xv] = append(m[-xv], []int{x, xv})
	}

	for i, iv := range nums {
		for j := i + 1; j < len(nums); j++ {
			jv := nums[j]
			ary, _ := m[iv+jv]
			// ary => [[idx, iv]....]
			for _, av := range ary {
				if av[0] != i && av[0] != j {
					if checkOutput(output, iv, jv, av[1]) == true {
						output = append(output, []int{iv, jv, av[1]})
					}
				}
			}
		}
	}

	return output
}
