package rand

import (
	"math/rand"
	"time"
)

const (
	RAND_KIND_NUM   = 0 // 纯数字
	RAND_KIND_LOWER = 1 // 小写字母
	RAND_KIND_UPPER = 2 // 大写字母
	RAND_KIND_ALL   = 3 // 数字、大小写字母
)

func MakeRand(length int, mode int) []byte {
	ikind, kinds, result := mode, [][]int{[]int{10, 48}, []int{26, 97}, []int{26, 65}}, make([]byte, length)
	isAll := mode > 2 || mode < 0
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < length; i++ {
		if isAll { // random ikind
			ikind = rand.Intn(3)
		}
		scope, base := kinds[ikind][0], kinds[ikind][1]
		result[i] = uint8(base + rand.Intn(scope))
	}
	return result
}
