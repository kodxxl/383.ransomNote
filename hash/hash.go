package hash

func CanConstruct(ransomNote string, magazine string) bool {
	hash := make(map[rune]int)
	for _, value := range magazine {
		hash[value]++
	}
	for _, value := range ransomNote {
		if i, present := hash[value]; (!present) || i <= 0 {
			return false
		}
		hash[value]--
	}
	return true
}
