package utils

import (
	"fmt"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var (
	timestampRegex *regexp.Regexp
)

func init() {
	pattern := "^[0-5][0-9]:[0-5][0-9]$"
	timestampRegex, _ = regexp.Compile(pattern)
}

func ValidateTimestamp(time string) bool {
	return timestampRegex.MatchString(time)
}

func ConvertToSeconds(time string) float64 {
	a := strings.Split(time, ":")
	m, _ := strconv.Atoi(a[0])
	s, _ := strconv.Atoi(a[1])
	return float64((60 * m) + s)
}

func RandomNumber(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	return min + rand.Intn(max-min+1)
}

func TimestampFileName() string {
	t := time.Now().Format("2006-01-02")
	CreateIfNotExist("output")
	CreateIfNotExist("output\\" + t)
	return fmt.Sprintf("output\\%s\\REEL-%d-%d.mp4", t, time.Now().Unix(), RandomNumber(100, 999))
}

func RandomTimestamp(vd float64, d float64) (float64, float64) {
	start := float64(RandomNumber(0, int(vd-d)))
	end := start + d

	return start, end
}
