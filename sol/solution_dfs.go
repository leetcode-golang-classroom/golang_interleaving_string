package sol

type StartRecord struct {
	start1, start2 int
}

func isInterleaveDFS(s1 string, s2 string, s3 string) bool {
	s1Len, s2Len, s3Len := len(s1), len(s2), len(s3)
	if s1Len+s2Len != s3Len {
		return false
	}
	cache := make(map[StartRecord]bool)
	var dfs func(start1, start2 int) bool
	dfs = func(start1, start2 int) bool {
		if start1+start2 == s3Len {
			return true
		}
		record := StartRecord{start1: start1, start2: start2}
		if val, exists := cache[record]; exists {
			return val
		}
		if start1 < s1Len && s1[start1] == s3[start1+start2] && dfs(start1+1, start2) {
			return true
		}
		if start2 < s2Len && s2[start2] == s3[start1+start2] && dfs(start1, start2+1) {
			return true
		}
		cache[record] = false
		return false
	}
	return dfs(0, 0)
}
