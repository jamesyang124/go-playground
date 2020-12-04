package palindromcheck

// https://www.algoexpert.io/questions/Palindrome%20Check
// given string ex: "abcdcba"
// is palindrome if tail and end the same in each sub-arrays which shrink from tail and head by 1
// a == a, then check "bcdcb" and so on

// IsPalindrome ...
// https://yourbasic.org/golang/for-loop-range-array-slice-map-channel/
// n -2 -2 -2... => O(n) time, O(1) space
func IsPalindrome(str string) bool {
	// Write your code here.
	check := true
	for i := 0; i < len(str); i++ {
		j := len(str) - 1 - i
		if i >= j {
			break
		}
		if str[i] == str[j] {
			continue
		} else {
			check = false
			break
		}
	}
	return check
}
