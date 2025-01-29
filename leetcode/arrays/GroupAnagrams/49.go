package main

import "sort"

func groupAnagrams(strs []string) [][]string {
	mapS := make(map[string][]string, 0)

	for _, val := range strs {
		bytesStr := []byte(val)
		sort.Slice(bytesStr, func(i, j int) bool {
			return bytesStr[i] < bytesStr[j]
		})

		sortedStr := string(bytesStr)
		mapS[sortedStr] = append(mapS[sortedStr], val)
	}

	ans := make([][]string, 0)
	for _, val := range mapS {
		ans = append(ans, val)
	}
	return ans
}

func isAnagram(s string, t string) bool {

	if len(s) != len(t) {
		return false
	}

	charC := [26]int{}

	for i := 0; i < len(s); i++ {
		charC[s[i]-'a']++
		charC[t[i]-'a']--
	}

	for _, c := range charC {
		if c != 0 {
			return false
		}
	}
	return true
}
