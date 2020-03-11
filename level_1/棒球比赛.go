package level_1

import "strconv"

func calPoints(ops []string) int {
	if 0 == len(ops) {
		return 0
	}

	var scores []int // 记录得分

	for _, v := range ops {
		scoreCount := len(scores)

		switch v {
		case "+":
			if 0 < len(scores) {
				if 1 < len(scores) {
					score := scores[scoreCount-1] + scores[scoreCount-2]
					scores = append(scores, score)
				} else {
					score := scores[scoreCount-1]
					scores = append(scores, score)
				}
			}
		case "D":
			if 0 < len(scores) {
				score := scores[scoreCount-1] * 2
				scores = append(scores, score)
			}
		case "C":
			if 0 < len(scores) {
				scores = scores[:scoreCount-1]
			}
		default:
			score, _ := strconv.Atoi(v)
			scores = append(scores, score)
		}
	}

	var sumScore int
	for _, v := range scores {
		sumScore += v
	}

	return sumScore
}

func CalPoints(ops []string) int {
	return calPoints(ops)
}
