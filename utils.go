package golodash

import (
	"log"
	"time"
	"os"
	"strings"
)

// Get cwd path.
func GetCwd() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return dir
}

// []string indexOf
func HasValue(a *[]string, it string) (bool, string) {
	for _, v := range *a {
		if v == it {
			return true, v
		}
	}
	return false, ""
}

// Get current date format "2006-01-02"
func GetCurrentDate(date time.Time) string {
	t := date.String()
	return strings.Split(t, " ")[0]
}
