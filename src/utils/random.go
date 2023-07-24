package utils

import (
	"math/rand"
	"strings"
	"time"
)

var r *rand.Rand

// Выполняется при превом обращении к пакету
func init(){
	r = rand.New(rand.NewSource(time.Now().Unix()))
}

func RandInt(min, max int64) int{
	return int(min + r.Int63n(max - min + 1))
}

func RandString(n int) string{
	sb := strings.Builder{}
	l := 'z' - 'a' + 1
	for i := 0; i < n; i++{
		b := 'a' + byte(RandInt(0, int64(l)))
		sb.WriteByte(b)
	}
	return sb.String()
}