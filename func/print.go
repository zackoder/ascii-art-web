package web

import (
	"strings"
)

func SplitAndPrint(s string, mapFont map[int][]string) string {
	if mapFont == nil {
		return ""
	}
	out_split := strings.Split(s, string([]byte{13, 10})) // split with \r and \n
	var out string
	for _, word := range out_split { // rang in array
		if err, i := Checkout(word); !err { // if err ==false return err message
			return "you can't use this litter : " + string(i)
		}
		out += OutOfFont(word, mapFont) // add line by line to var string
	}
	return out
}

func OutOfFont(s string, mapFont map[int][]string) string {
	var strreturn string
	for i := 0; i < 8; i++ { // range 8 time
		for _, p := range s { // range in word and take font in map
			strreturn += mapFont[int(p)][i] // add head by head and body by body and last bottum
		}
		strreturn += "\n" // make /n in last
	}
	return strreturn
}
