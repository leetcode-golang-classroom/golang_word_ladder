package sol

type Words []string

func ladderLength(beginWord string, endWord string, wordList []string) int {
	// check if endWord in WordList
	if !contain(endWord, wordList) {
		return 0
	}
	patternMap := make(map[string]Words)
	visit := make(map[string]struct{})
	wordList = append(wordList, beginWord)
	for _, word := range wordList {
		for idx := range word {
			pattern := word[:idx] + "*" + word[idx+1:]
			patternMap[pattern] = append(patternMap[pattern], word)
		}
	}
	// init
	visit[beginWord] = struct{}{}
	queue := []string{beginWord}
	res := 1
	for len(queue) != 0 {
		qLen := len(queue)
		for idx := 0; idx < qLen; idx++ {
			// pop element from queue
			top := queue[0]
			queue = queue[1:]
			if top == endWord {
				return res
			}
			// generate pattern to check
			for idx := range top {
				pattern := top[:idx] + "*" + top[idx+1:]
				patternList := patternMap[pattern]
				for _, word := range patternList {
					if _, ok := visit[word]; !ok {
						visit[word] = struct{}{}
						queue = append(queue, word)
					}
				}
			}
		}
		res++
	}
	return 0
}

func contain(endWord string, wordList []string) bool {
	for _, word := range wordList {
		if word == endWord {
			return true
		}
	}
	return false
}
