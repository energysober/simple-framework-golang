package consistenthash

import (
	"strconv"
	"testing"
)

func TestHashing(t *testing.T) {
	hash := New(3, func(data []byte) uint32 {
		i, _ := strconv.Atoi(string(data))
		return uint32(i)
	})

	// 3, 6 ,9, 13, 16, 19, 23, 26, 29
	hash.Add("9", "6", "3")

	testCases := map[string]string{
		"1":  "3",
		"5":  "6",
		"9":  "9",
		"20": "3",
		"27": "9",
		"33": "3",
	}

	for k, v := range testCases {
		if hash.Get(k) != v {
			t.Errorf("Asking for %s, should have yielded %s", k, v)
		}
	}

	// add 8, 18, 28
	hash.Add("8")
	testCases["27"] = "8"

	for k, v := range testCases {
		if hash.Get(k) != v {
			t.Errorf("Asking for %s, should have yielded %s", k, v)
		}
	}
}
