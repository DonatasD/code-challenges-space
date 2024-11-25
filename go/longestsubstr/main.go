package main

func lengthOfLongestSubstring(s string) int {
	m := make(map[int32]rune)
	l := 0
	res := 0
	strArr := []rune(s)

	for r, _ := range strArr {
		for _, exist := m[strArr[r]]; exist; {
			delete(m, strArr[l])
			l += 1
			_, exist = m[strArr[r]]
		}
		m[strArr[r]] = strArr[r]
		if r-l+1 > res {
			res = r - l + 1
		}
	}
	return res
}

func main() {
	println(lengthOfLongestSubstring("dvdf"))

}
