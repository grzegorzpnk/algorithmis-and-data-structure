package ArraysAndHash

import "sort"

func topKFrequentWords(words []string, k int) []string {

	freqMap := make(map[string]int)

	for _, v := range words {
		freqMap[v]++
	}

	bucket := make([][]string, len(words)+1)

	for word, occurency := range freqMap {
		bucket[occurency] = append(bucket[occurency], word)
	}

	ans := make([]string, 0, k)

	for i := len(bucket) - 1; i >= 0; i-- {
		if len(bucket[i]) > 1 {
			sort.Strings(bucket[i])
		}
		for _, word := range bucket[i] {
			if k > 0 {
				ans = append(ans, word)
				k--
			}
		}

	}

	return ans

}
