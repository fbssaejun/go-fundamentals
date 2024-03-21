package main

func getNameCounts(names []string) map[rune]map[string]int {
	counts := make(map[rune]map[string]int)
	for _, name := range names {
		if name == "" {
			continue
		}
		firstChar := rune(name[0])
		_, ok := counts[firstChar]
		if !ok {
			counts[firstChar] = make(map[string]int)
		}
		counts[firstChar][name]++
	}

	return counts
}

func main() {
	names := []string{"John", "Jane", "Doe", "John", "Jane", "Doe", "John", "Jane", "Doe"}
	nameCounts := getNameCounts(names)
	for firstChar, names := range nameCounts {
		for name, count := range names {
			println(string(firstChar), name, count)
		}
	}
}
