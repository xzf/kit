package jsonX

import "encoding/json"

func MarshalIndentStr(obj interface{}) string {
	byteSlice, err := json.MarshalIndent(obj, "", "\t")
	if err != nil {
		panic(err)
	}
	return string(byteSlice)
}

func MarshalStr(obj interface{}) string {
	byteSlice, err := json.Marshal(obj)
	if err != nil {
		panic(err)
	}
	return string(byteSlice)
}
