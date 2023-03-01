package main

import (
	"bufio"
	"os"
	"strings"
)

// 换行的数据
func ReadLineData(userDict string) (users []string, err error) {
	file, err := os.Open(userDict)
	if err != nil {
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		user := strings.TrimSpace(scanner.Text())
		if user != "" {
			users = append(users, user)
		}
	}
	return users, err
}
func In_array(needle interface{}, hystack interface{}) bool {
	switch key := needle.(type) {
	case string:
		for _, item := range hystack.([]string) {
			if key == item {
				return true
			}
		}
	case int:
		for _, item := range hystack.([]int) {
			if key == item {
				return true
			}
		}
	case int64:
		for _, item := range hystack.([]int64) {
			if key == item {
				return true
			}
		}
	default:
		return false
	}
	return false
}

func WriteAppend(path, str string) {
	f, err := os.OpenFile(path, os.O_CREATE+os.O_RDWR+os.O_APPEND, 0764)
	if err != nil {
	}

	//jsonBuf := append([]byte(result),[]byte("\r\n")...)
	str += "\n"
	f.WriteString(str)
}

func main() {
	set := []string{}

	hosts, _ := ReadLineData("domain")

	for _, d := range hosts {
		if !In_array(d, set) {
			WriteAppend("qc", d)
		}
		set = append(set, d)
	}

}
