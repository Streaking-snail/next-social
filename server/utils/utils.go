package utils

import (
	"strconv"
	"strings"

	"github.com/google/uuid"
)

func UUID() string {
	return uuid.New().String()
}
func LongUUID() string {
	uuid.New()
	longUUID := strings.Join([]string{UUID(), UUID(), UUID(), UUID()}, "")
	return strings.ReplaceAll(longUUID, "-", "")
}

func IpToInt(ip string) int64 {
	if len(ip) == 0 {
		return 0
	}
	bits := strings.Split(ip, ".")
	if len(bits) < 4 {
		return 0
	}
	b0 := StringToInt(bits[0])
	b1 := StringToInt(bits[1])
	b2 := StringToInt(bits[2])
	b3 := StringToInt(bits[3])

	var sum int64
	sum += int64(b0) << 24
	sum += int64(b1) << 16
	sum += int64(b2) << 8
	sum += int64(b3)

	return sum
}

func StringToInt(in string) (out int) {
	out, _ = strconv.Atoi(in)
	return
}
