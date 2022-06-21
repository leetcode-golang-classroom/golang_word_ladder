# golang_word_ladder

A **transformation sequence** from word `beginWord` to word `endWord` using a dictionary `wordList` is a sequence of words `beginWord -> s1 -> s2 -> ... -> sk` such that:

- Every adjacent pair of words differs by a single letter.
- Every `si` for `1 <= i <= k` is in `wordList`. Note that `beginWord` does not need to be in `wordList`.
- `sk == endWord`

Given two words, `beginWord` and `endWord`, and a dictionary `wordList`, return *the **number of words** in the **shortest transformation sequence** from* `beginWord` *to* `endWord`*, or* `0` *if no such sequence exists.*

## Examples

**Example 1:**

```
Input: beginWord = "hit", endWord = "cog", wordList = ["hot","dot","dog","lot","log","cog"]
Output: 5
Explanation: One shortest transformation sequence is "hit" -> "hot" -> "dot" -> "dog" -> cog", which is 5 words long.

```

**Example 2:**

```
Input: beginWord = "hit", endWord = "cog", wordList = ["hot","dot","dog","lot","log"]
Output: 0
Explanation: The endWord "cog" is not in wordList, therefore there is no valid transformation sequence.
```

**Constraints:**

- `1 <= beginWord.length <= 10`
- `endWord.length == beginWord.length`
- `1 <= wordList.length <= 5000`
- `wordList[i].length == beginWord.length`
- `beginWord`, `endWord`, and `wordList[i]` consist of lowercase English letters.
- `beginWord != endWord`
- All the words in `wordList` are **unique**.

## 解析

題目給定一個起始字串 beginWord, 結束字串 endWord 還有一個字串串列 wordList 

找到一個字串串列從 beginWord 開始,到 endWord 結束

中間每兩個鄰近的 string 之間都只能有一個字元不同

要求寫一個演算法來計算這個串列的最短長度，假設存在這個串列的話，否則回傳 0

困難的點在於如何去找到下一個可能的鄰近單字

如果從 pattern的角度來思考

假設字串 s = “hot”

那 s 可能的下個字串 pattern 

“*ot” ：只有第一個字元不同

“h*t"  : 只有第二個字元不同

“ho*” : 只有第三個字元不同

只要能把這類 pattern 根據字串不同個別蒐集在一個 hashTable

這樣從第一個字開始

做 BFS 找尋下一個可能字串

直到遇到字串是 endWord 就是結尾

如果整個 hashTable 都走訪完

都沒有走到 endWord 代表沒有這樣的 path

## 程式碼
```go
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

```
## 困難點

1. 設計出可以找到一下個可能 word的方式
2. 透過 HashSet 來避免重複走訪相同 word

## Solve Point

- [x]  需要找到 match 下一個可能的word 的作法
- [x]  透過 BFS 來實作
- [x]  透過 visit 來避免重複拜訪相同的 word