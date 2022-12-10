package windowsliding

import (
	"testing"
)

func TestSlidingWindowFindMaxSubstring(t *testing.T) {
	t.Parallel()

	expectedResults := map[string]string{
		"abcabcbb": "abc",
		"":         "",
		" ":        "",
		"bbbbb":    "b",
		"pwwkew":   "wke",
	}

	for k, v := range expectedResults {
		result := slidingWindowFindMaxSubstring(k)
		if result != v {
			t.Fatalf("Expected result: %v, Got: %v", v, result)
		}
	}

}
