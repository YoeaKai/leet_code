package text_justification

import "bytes"

func FullJustify(words []string, maxWidth int) []string {
	lineNums := []int{}
	l := 0

	// Find the number of words in each line.
	for _, word := range words {
		if l == 0 || l+len(word)+1 > maxWidth {
			l = len(word)
			lineNums = append(lineNums, 1)
		} else {
			l = l + len(word) + 1
			lineNums[len(lineNums)-1]++
		}
	}

	res := make([]string, len(lineNums))
	start := 0

	for i, wordNums := range lineNums {
		if wordNums == 1 || i == len(lineNums)-1 {
			res[i] = fillRight(words[start:start+wordNums], maxWidth)
		} else {
			res[i] = fillMedium(words[start:start+wordNums], maxWidth)
		}
		start += wordNums
	}

	return res
}

func fillMedium(words []string, maxWidth int) string {
	wordsLen := 0

	for w := 0; w < len(words); w++ {
		wordsLen += len(words[w])
	}

	var res bytes.Buffer
	spaceNum := (maxWidth - wordsLen) / (len(words) - 1)
	leftover := (maxWidth - wordsLen) % (len(words) - 1)

	for i := 0; i < len(words)-1; i++ {
		res.WriteString(words[i])
		for x := 0; x < spaceNum; x++ {
			res.WriteString(" ")
		}
		if leftover > 0 {
			res.WriteString(" ")
			leftover--
		}
	}

	res.WriteString(words[len(words)-1])

	return res.String()
}

func fillRight(words []string, maxWidth int) string {
	var res bytes.Buffer
	l := len(words[0])

	res.WriteString(words[0])

	for i := 1; i < len(words); i++ {
		l += len(words[i])
		l++
		res.WriteString(" ")
		res.WriteString(words[i])
	}

	for j := 0; j < maxWidth-l; j++ {
		res.WriteString(" ")
	}

	return res.String()
}
