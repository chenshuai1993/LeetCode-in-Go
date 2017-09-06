package Problem0076

import (
	"sort"
)

func minWindow(s string, t string) string {
	rec := make(map[byte][]int, len(s))

	for i := range s {
		rec[s[i]] = append(rec[s[i]], i)
	}

	// 检查，t 中的字母是否都在 s 中
	for i := range t {
		if len(rec[t[i]]) == 0 {
			return ""
		}
	}

	min := len(s) + 1
	first, last := 0, 0
	list := make([]int, len(t))

	var dfs func(int)
	dfs = func(idx int) {
		if idx == len(t) {
			temp := make([]int, len(t))
			copy(temp, list)
			sort.Ints(temp)
			dist := temp[len(t)-1] - temp[0]
			if min > dist {
				first, last = temp[0], temp[len(t)-1]
			}
		}

		for i := 0; i < len(rec[t[idx]]); i++ {
			list[idx] = rec[t[idx]][i]
			dfs(idx + 1)
		}
	}

	dfs(0)

	return s[first : last+1]
}