package utils

import "encoding/json"

func PrintJSON(v interface{}) {
	bs, _ := json.MarshalIndent(v, "", "    ")
	println(string(bs))
}
