package jsonX

import "encoding/json"

func MarshalIndent(obj interface{}) []byte {
	byteSlice, err := json.MarshalIndent(obj, "", "\t")
	if err != nil {
		panic(err)
	}
	return byteSlice
}

func MarshalIndentStr(obj interface{}) string {
	return string(MarshalIndent(obj))
}

func Marshal(obj interface{}) []byte {
	byteSlice, err := json.Marshal(obj)
	if err != nil {
		panic(err)
	}
	return byteSlice
}

func MarshalStr(obj interface{}) string {
	return string(Marshal(obj))
}
