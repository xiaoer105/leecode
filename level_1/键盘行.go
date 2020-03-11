package level_1

/*
给定一个单词列表，只返回可以使用在键盘同一行的字母打印出来的单词。键盘如下图所示。

示例：

输入: ["Hello", "Alaska", "Dad", "Peace"]
输出: ["Alaska", "Dad"]

注意：

你可以重复使用键盘上同一字符。
你可以假设输入的字符串将只包含字母。
*/

func findWords(words []string) []string {
	if 0 == len(words) {
		return nil
	}

	keyBoardMap := map[string]int{
		"q": 1, "w": 1, "e": 1, "r": 1, "t": 1, "y": 1, "u": 1, "i": 1, "o": 1, "p": 1,
		"a": 2, "s": 2, "d": 2, "f": 2, "g": 2, "h": 2, "j": 2, "k": 2, "l": 2,
		"z": 3, "x": 3, "c": 3, "v": 3, "b": 3, "n": 3, "m": 3}

	var strs []string
	for _, word := range words {
		var isRow = true
		var prevRow int
		for _, v := range word {
			if row, ok := keyBoardMap[string(v)]; !ok {
				continue
			} else {
				if 0 != prevRow && prevRow != row {
					isRow = false
					continue
				}

				prevRow = row
			}
		}

		if isRow {
			strs = append(strs, word)
		}
	}

	return strs
}

func FindWords(words []string) []string {
	return findWords(words)
}
