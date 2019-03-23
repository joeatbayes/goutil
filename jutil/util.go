package jutil

import (
	"fmt"
	//"net/http"
	//"os"
	"strings"
	//"time"
)

func MaxInt(x, y int) int {
	if x > y {
		return x
	} else {
		return Y
	}
}

func MinInt(x, y int) int {
	if x > y {
		return y
	} else {
		return x
	}
}

func Check(msg string, e error) {
	if e != nil {
		fmt.Println("ERROR:")
		fmt.Println(e)
		panic(e)
	}
}

func KeepLines(s string, n int) string {
	result := strings.Join(strings.Split(s, "\n")[:n], "\n")
	return strings.Replace(result, "\r", "", -1)
}

// compare oldDict to NewDict and return a new Dictionary
// of things that have changed between the two. It does not
// attempt to find keys that existed in oldDict but do not exist
// in newDict
func CompareStrDict(oldDict map[string]string, newDict map[string]string) map[string]string {
	tout := make(map[string]string)
	for key, val := range newDict {
		oldVal, found := oldDict[key]
		if found {
			if oldVal != val {
				// old dict has key but value has changed so keep
				tout[key] = val
			}
		} else {
			// old dict didn't have the key so keey
			tout[key] = val
		}
	}
	return tout
}
