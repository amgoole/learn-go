package substr

// MaxSubStrNoRepeat 仅支持英文字母
func MaxSubStrNoRepeat(s string) int {
	lastRepeat := make(map[string]int)
	start := 0
	MaxLen := 0
	for i, ch := range s {
		if index, ok := lastRepeat[string(ch)]; ok && index >= start {
			start = index + 1
		}
		if MaxLen < i-start+1 {
			MaxLen = i - start + 1
		}
		lastRepeat[string(ch)] = i
	}
	return MaxLen
}

// MaxSubStrNoRepeatIntl 支持中文
func MaxSubStrNoRepeatIntl(s string) int {
	lastRepeat := make(map[rune]int)
	start := 0
	MaxLen := 0
	for i, ch := range []rune(s) {
		if index, ok := lastRepeat[ch]; ok && index >= start {
			start = index + 1
		}
		if MaxLen < i-start+1 {
			MaxLen = i - start + 1
		}
		lastRepeat[ch] = i
	}
	return MaxLen
}

// ProfileMaxSubStrNoRepeatIntl 优化速度
func ProfileMaxSubStrNoRepeatIntl(s string) int {
	lastRepeat := make([]int, 0xFFFF)
	for i, _ := range lastRepeat {
		lastRepeat[i] = -1
	}
	start := 0
	MaxLen := 0
	for i, ch := range []rune(s) {
		if index:= lastRepeat[ch]; index != -1 && index >= start {
			start = index + 1
		}
		if MaxLen < i-start+1 {
			MaxLen = i - start + 1
		}
		lastRepeat[ch] = i
	}
	return MaxLen
}
