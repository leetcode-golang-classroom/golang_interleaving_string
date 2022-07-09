package sol

func isInterleaveV1(s1 string, s2 string, s3 string) bool {
	s1Len, s2Len, s3Len := len(s1), len(s2), len(s3)
	if s1Len+s2Len != s3Len {
		return false
	}
	dp := make([]bool, s2Len+1)
	dp[s2Len] = true
	// dp[start2] s2 start from start2, s3 start from start2+0, start2+1, ... , start2 + s1Len
	for start1 := s1Len; start1 >= 0; start1-- {
		for start2 := s2Len; start2 >= 0; start2-- {
			if start1 < s1Len && s1[start1] == s3[start1+start2] && dp[start2] {
				dp[start2] = true
				continue
			}
			if start2 < s2Len && s2[start2] == s3[start1+start2] && dp[start2+1] {
				dp[start2] = true
			}
		}
	}
	return dp[0]
}
