package main

func canConstruct(ransomNote string, magazine string) bool {
	letters := make(map[rune]int)

	for _, r := range magazine {
		val, ok := letters[r]
		if ok {
			letters[r] = val + 1
		} else {
			letters[r] = 1
		}
	}

	for _, r := range ransomNote {
		val, ok := letters[r]
		if ok {
			if val-1 == 0 {
				delete(letters, r)
			} else {
				letters[r] = val - 1
			}
		} else {
			return false
		}
	}

	return true
}

// Time complexity: O(len(magazine)) + O(len(randomNote)) ~= O(n). Going through both strings once.
// Space complexity: O(len(magazine). We keep the letters of the magazine in memory.
