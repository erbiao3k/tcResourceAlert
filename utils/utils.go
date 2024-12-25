package utils

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// SplitIntoChunks 将切片按照指定大小进行拆分
func SplitIntoChunks(slice []string, chunkSize int) [][]string {
	var chunks [][]string
	for i := 0; i < len(slice); i += chunkSize {
		end := i + chunkSize
		if end > len(slice) {
			end = len(slice)
		}
		chunks = append(chunks, slice[i:end])
	}
	return chunks
}

func Average(list []*float64) float64 {
	sum := 0.0
	for _, value := range list {
		if value != nil {
			sum += *value
		}
	}
	return sum / float64(len(list))
}

// RangeTime 返回正确的接口时间范围
func RangeTime(endMinute, startMinute time.Duration) (string, string) {
	format := "2006-01-02T15:04:00-07:00"
	currentTime := time.Now()
	currentMinute := currentTime.Truncate(time.Minute)
	previousEndMinute := currentMinute.Add(-time.Minute * endMinute)
	previousStartMinute := currentTime.Add(-time.Minute * startMinute)

	return previousStartMinute.Format(format), previousEndMinute.Format(format)
}

func Nowtime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func Json(key interface{}) string {
	bytes, _ := json.Marshal(key)
	return string(bytes)
}

func Inslice(slice []string, key string) bool {
	for _, s := range slice {
		if s == key {
			return true
		}
	}
	return false
}

func Float64Str(n float64) string {
	return strconv.FormatFloat(n, 'f', 2, 64)
}

func Float(n float64) float64 {
	parsedNum, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", n), 64)
	return parsedNum
}

func ReplaceMsg(str string) string {
	str = strings.ReplaceAll(str, "\"", "")
	str = strings.ReplaceAll(str, ",", "\n")
	str = strings.ReplaceAll(str, ":", "：")
	str = strings.ReplaceAll(str, "{", "")
	str = strings.ReplaceAll(str, "}", "")
	return str
}

func SlowRun() {
	currentTime := time.Now()
	startTime := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 2, 0, 0, 0, currentTime.Location())
	endTime := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 5, 0, 0, 0, currentTime.Location())
	if currentTime.After(startTime) && currentTime.Before(endTime) {
		time.Sleep(time.Hour * 1)
	}
}
