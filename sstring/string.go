package sstring

func IsPalindrome(word string) bool {
	start := 0
	end := len(word) - 1
	for start < end {
		letterStart := word[start]
		letterEnd := word[end]
		if letterStart != letterEnd {
			return false
		}
		start++
		end--
	}
	return true
}
