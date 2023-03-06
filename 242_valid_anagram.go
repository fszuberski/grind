package main

func main_242() {
	// s = "anagram", t = "nagaram" -> true
	// s = "rat", t = "car" -> false
	// s = "ab", t = "a" -> false

}

func isAnagram(s string, t string) bool {
	letters := make(map[rune]int)

	for _, r := range s {
		val, ok := letters[r]
		if ok {
			letters[r] = val + 1
		} else {
			letters[r] = 1
		}
	}

	for _, r := range t {
		val, ok := letters[r]
		if !ok {
			return false
		}

		if val-1 == 0 {
			delete(letters, r)
		} else {
			letters[r] = val - 1
		}
	}

	return len(letters) == 0
}

// Time complexity: O(len(s) + len(t)) ~= O(n). A single pass through both strings is required.
// Space complexity: O(len(s)). A hashmap of all the values from s is required.
