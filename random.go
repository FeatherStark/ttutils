package ttutils

import (
	"math/rand"
	"time"
)

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// RandomNumber 获取指定数字区间的随机数
// Args: min int 最小值, max int 最大值
// Returns: int 随机数
// e.g.: Args: 1, 10 Returns: 5
func RandomNumber(min, max int) int {
	if min > max {
		min, max = max, min
	}
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min
}

// RandomString 获取指定长度的随机字符串
// Args: length int 字符串长度
// Returns: string 指定长度的随机字符串
// e.g.: Args: 10 Returns: "oq3kefg7ij"
func RandomString(length int) string {
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)
	result := make([]byte, length)
	for i := range result {
		result[i] = letters[r.Intn(len(letters))]
	}
	return string(result)
}
