package minimum_genetic_mutation

func MinMutation(start string, end string, bank []string) (steps int) {
	bankMap := make(map[string]struct{})
	for _, gene := range bank {
		bankMap[gene] = struct{}{}
	}

	if _, ok := bankMap[end]; !ok {
		return -1
	}

	queue := []string{start}
	var tmp string
	var length int

	for len(queue) != 0 {
		length = len(queue)
		for i := 0; i < length; i++ {
			if queue[i] == end {
				return steps
			}

			delete(bankMap, queue[i])

			for j := 0; j < 8; j++ {
				tmp = queue[i]
				tmp = tmp[:j] + "A" + tmp[j+1:]
				if _, ok := bankMap[tmp]; ok {
					queue = append(queue, tmp)
				}
				tmp = tmp[:j] + "C" + tmp[j+1:]
				if _, ok := bankMap[tmp]; ok {
					queue = append(queue, tmp)
				}
				tmp = tmp[:j] + "G" + tmp[j+1:]
				if _, ok := bankMap[tmp]; ok {
					queue = append(queue, tmp)
				}
				tmp = tmp[:j] + "T" + tmp[j+1:]
				if _, ok := bankMap[tmp]; ok {
					queue = append(queue, tmp)
				}
			}
		}
		queue = queue[length:]
		steps++
	}

	return -1
}
