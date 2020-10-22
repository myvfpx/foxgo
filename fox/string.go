package fox

import (
	"math"
	"strings"
)

// 由一个表达式返回一个字符串，并从两边用空格或字符把该字符串填充到指定长度。
func Padc(s string, width int, pad string) string {
	gap := width - len(s)
	if gap > 0 {
		gapLeft := int(math.Ceil(float64(gap / 2)))
		gapRight := gap - gapLeft
		return strings.Repeat(pad, gapLeft) + s + strings.Repeat(pad, gapRight)
	}
	return s
}

// 由一个表达式返回一个字符串，并从右边用空格或字符把该字符串填充到指定长度。
func Padr(s string, width int, pad string) string {
	gap := width - len(s)
	if gap > 0 {
		return s + strings.Repeat(pad, gap)
	}
	return s
}

// 由一个表达式返回一个字符串，并从左边用空格或字符把该字符串填充到指定长度。
func Padl(s string, width int, pad string) string {
	gap := width - len(s)
	if gap > 0 {
		return strings.Repeat(pad, gap) + s
	}
	return s
}
