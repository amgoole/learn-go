package substr

import (
	"testing"
)

func TestMaxSubStrNoRepeat(t *testing.T) {
	tests := []struct {
		s      string
		length int
	}{
		{"abc", 3},
		{"abcabcbb", 3},
		{"bbbbb", 1},
		{"pwwkpw", 3},
	}
	for _, tt := range tests {
		if actual := MaxSubStrNoRepeat(tt.s); actual != tt.length {
			t.Errorf("str: %s | expected %d, result %d", tt.s, tt.length, actual)
		}
	}
}

func TestMaxSubStrNoRepeatIntl(t *testing.T) {
	tests := []struct {
		s      string
		length int
	}{
		{"你好呀", 3},
		{"yes中国y", 5},
		{"黑化肥挥发发灰会花飞灰化肥挥发发黑会飞花", 8},
	}
	for _, tt := range tests {
		if actual := MaxSubStrNoRepeatIntl(tt.s); actual != tt.length {
			t.Errorf("str: %s | expected %d, result %d", tt.s, tt.length, actual)
		}
	}
}

func TestProfileMaxSubStrNoRepeatIntl(t *testing.T) {
	tests := []struct {
		s      string
		length int
	}{
		{"你好呀", 3},
		{"yes中国y", 5},
		{"黑化肥挥发发灰会花飞灰化肥挥发发黑会飞花", 8},
	}
	for _, tt := range tests {
		if actual := ProfileMaxSubStrNoRepeatIntl(tt.s); actual != tt.length {
			t.Errorf("str: %s | expected %d, result %d", tt.s, tt.length, actual)
		}
	}
}

func BenchmarkMaxSubStrNoRepeatIntl(b *testing.B) {
	s := "黑化肥挥发发灰会花飞灰化肥挥发发黑会飞花"
	ans := 8
	for i := 0; i < b.N; i++ {
		actual := MaxSubStrNoRepeatIntl(s)
		if actual != ans {
			b.Errorf("expected %d, result %d", ans, actual)
		}
	}
}

func BenchmarkProfileMaxSubStrNoRepeatIntl(b *testing.B) {
	s := "黑化肥挥发发灰会花飞灰化肥挥发发黑会飞花"
	for i := 0; i < 13; i++ {
		s = s + s
	}
	ans := 8
	b.Logf("len(s)=%d\n", len(s))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		actual := ProfileMaxSubStrNoRepeatIntl(s)
		if actual != ans {
			b.Errorf("expected %d, result %d", ans, actual)
		}
	}
}
