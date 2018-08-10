package letter

// ConcurrentFrequency ...
func ConcurrentFrequency(texts []string) FreqMap {
	cn := make(chan FreqMap, len(texts))
	for _, words := range texts {
		go func(w string) {
			cn <- Frequency(w)
		}(words)
	}
	freq := FreqMap{}
	for range texts {
		for key, value := range <-cn {
			freq[key] += value
		}
	}
	return freq
}
