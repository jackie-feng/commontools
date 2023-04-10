package lcs

type Pair struct {
	Idx1 int
	Idx2 int
}

type Equal func(i, j int) bool

// Sequence 最长公共子序列 longest_common_sequence
// c[i,j] = 0                           i = 0 || j = 0
//          c[i-1, j-1] + 1             i,j > 0 && xi  = yj
//          max(c[i-1, j], c[i, j-1])   i,j > 0 && xi != yj
func Sequence(l1, l2 int, equal Equal) []Pair {
	rows := make([][]int, l1+1)
	for i := 0; i <= l1; i++ {
		rows[i] = make([]int, l2+1)
	}
	res := []Pair{}
	for i := 0; i < l1+1; i++ {
		for j := 0; j < l2+1; j++ {
			if i == 0 || j == 0 {
				rows[i][j] = 0
				continue
			}

			if equal(i-1, j-1) {
				rows[i][j] = rows[i-1][j-1] + 1
				continue
			}
			if rows[i][j-1] > rows[i-1][j] {
				rows[i][j] = rows[i][j-1]
			} else {
				rows[i][j] = rows[i-1][j]
			}
		}
	}

	i := l1
	j := l2
	for i > 0 && j > 0 {
		if rows[i][j] <= 0 {
			break
		}
		if rows[i-1][j] == rows[i][j] {
			i = i - 1
			continue
		}

		if rows[i][j-1] == rows[i][j] {
			j = j - 1
			continue
		}
		res = append(res, Pair{Idx1: i - 1, Idx2: j - 1})
		i = i - 1
		j = j - 1
	}
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return res
}
