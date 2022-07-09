package sol

import "testing"

func BenchmarkTestV1(b *testing.B) {
	s1 := "aabcc"
	s2 := "dbbca"
	s3 := "aadbbbaccc"
	for idx := 0; idx < b.N; idx++ {
		isInterleaveV1(s1, s2, s3)
	}
}
func Test_isInterleaveV1(t *testing.T) {
	type args struct {
		s1 string
		s2 string
		s3 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "s1 = \"aabcc\", s2 = \"dbbca\", s3 = \"aadbbcbcac\"",
			args: args{s1: "aabcc", s2: "dbbca", s3: "aadbbcbcac"},
			want: true,
		},
		{
			name: "s1 = \"aabcc\", s2 = \"dbbca\", s3 = \"aadbbbaccc\"",
			args: args{s1: "aabcc", s2: "dbbca", s3: "aadbbbaccc"},
			want: false,
		},
		{
			name: "s1 = \"\", s2 = \"\", s3 = \"\"",
			args: args{s1: "", s2: "", s3: ""},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isInterleaveV1(tt.args.s1, tt.args.s2, tt.args.s3); got != tt.want {
				t.Errorf("isInterleaveV1() = %v, want %v", got, tt.want)
			}
		})
	}
}
