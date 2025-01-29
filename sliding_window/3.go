package slidingwindow

func slidingWindow(s string) int {
	left := 0
	result := 0
	memory := make(map[byte]int) // Для хранения состояния окна

	for right := 0; right < len(s); right++ {
		// Расширяем окно: добавляем s[right]
		memory[s[right]]++

		// Проверяем условие задачи
		for /* условие нарушения */ {
			// Сужаем окно: удаляем s[left]
			memory[s[left]]--
			if memory[s[left]] == 0 {
				delete(memory, s[left])
			}
			left++
		}

		// Обновляем результат
		// result = max(result, right-left+1)
	}

	return result
}

// lengthOfLongestSubstring returns the length of the longest substring
// without repeating characters.
func lengthOfLongestSubstring(s string) int {
	var left, right, result int

	// m is a map to store the index of the characters in the string.
	// key is the character, value is the index.
	m := make(map[byte]int)

	for right < len(s) {
		c := s[right]
		// If the character is already in the map,
		// update the left boundary to the right of the previous occurrence.
		if _, ok := m[c]; ok {
			left = max(left, m[c]+1)
		}
		// Update the index of the character in the map.
		m[c] = right
		// Update the result with the maximum length.
		result = max(result, right-left+1)
		right++
	}
	return result
}

// max returns the maximum of two integers.
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// countGoodSubstrings returns the number of substrings of length 3
// that have all distinct characters in the given string s.
func countGoodSubstrings(s string) int {
	// Initialize left boundary of the sliding window and result count
	var left, res int

	// memory is a map to store the frequency of characters in the current window
	memory := make(map[byte]int)

	// Iterate over the string with the right boundary of the sliding window
	for right := 0; right < len(s); right++ {
		// Increment the frequency count of the character at the right boundary
		memory[s[right]]++

		// If the window size exceeds 3, shrink it from the left
		if right-left+1 > 3 {
			left_c := s[left]
			// Decrement the frequency count of the character at the left boundary
			memory[left_c]--
			// If the frequency becomes zero, remove the character from the map
			if memory[left_c] == 0 {
				delete(memory, left_c)
			}
			// Move the left boundary to the right
			left++
		}

		// Check if the current window has exactly 3 distinct characters
		if len(memory) == 3 && right-left+1 == 3 {
			// Increment the result count
			res++
		}
	}
	return res
}

// minWindow returns the minimum window in string s which contains all the characters of string t.
// If there is no such window, it returns an empty string.
func minWindow(s string, t string) string {
	// Initialize maps to store character counts
	memory := make(map[byte]int)   // For the current window in s
	memory_t := make(map[byte]int) // For the string t

	// Fill memory_t with character counts from t
	for i := 0; i < len(t); i++ {
		memory_t[t[i]]++
	}

	// Sliding window variables
	l := 0               // Left pointer
	minLen := len(s) + 1 // Minimum window length (initialized to a value larger than possible)
	res := ""            // Resulting substring
	cur_size := 0        // Number of characters from t found in the current window

	// Main loop over string s
	for r := 0; r < len(s); r++ {
		char := s[r]

		// If the character is in t, update memory and cur_size
		if count, ok := memory_t[char]; ok {
			memory[char]++
			if memory[char] == count {
				cur_size++
			}
		}

		// If all characters from t are found in the current window
		for cur_size == len(memory_t) {
			// Update the result if the current window is smaller than the minimum
			if r-l+1 < minLen {
				minLen = r - l + 1
				res = s[l : r+1]
			}

			// Shift the left pointer to try to minimize the window
			leftChar := s[l]
			if _, ok := memory_t[leftChar]; ok {
				memory[leftChar]--
				if memory[leftChar] < memory_t[leftChar] {
					cur_size--
				}
			}
			l++
		}
	}

	return res
}

// totalFruit returns the length of the longest subarray with no more than two distinct elements.
func totalFruit(fruits []int) int {
	// memory is a map to store the frequency of fruits in the current window
	memory := make(map[int]int)
	// res is the result of the longest subarray with no more than two distinct elements
	res := 0
	// l is the left pointer of the sliding window
	l := 0

	// Iterate over the fruits array with the right pointer of the sliding window
	for r := 0; r < len(fruits); r++ {
		// Increment the frequency count of the fruit at the right boundary
		memory[fruits[r]]++

		// If the number of distinct fruits in the window exceeds 2, shrink the window from the left
		for len(memory) > 2 {
			// Decrement the frequency count of the fruit at the left boundary
			memory[fruits[l]]--
			// If the frequency becomes zero, remove the fruit from the map
			if memory[fruits[l]] == 0 {
				delete(memory, fruits[l])
			}
			// Move the left boundary to the right
			l++
		}

		// Update the result if the current window has a longer length
		res = max(res, r-l+1)

	}

	return res

}
