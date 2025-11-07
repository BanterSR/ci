package alg

import (
	"strconv"
)

func S2U32(msg string) uint32 {
	if msg == "" {
		return 0
	}
	ms, _ := strconv.ParseInt(msg, 10, 32)
	return uint32(ms)
}

func AddList[T any](list *[]*T, n *T) {
	if list == nil {
		list = new([]*T)
	}
	*list = append(*list, n)
}

func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
