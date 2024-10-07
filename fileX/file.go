package fileX

import (
	"os"
)

func Write(path string, content string) {
	f, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	_, err = f.WriteString(content)
	if err != nil {
		panic(err)
	}
	err = f.Close()
	if err != nil {
		panic(err)
	}
	return
}

func Read(path string) []byte {
	result, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return result
}

func Append(path string, content string) {
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND, os.ModeAppend)
	if err != nil {
		panic(err)
	}
	_, err = f.WriteString(content)
	if err != nil {
		panic(err)
	}
	err = f.Close()
	if err != nil {
		panic(err)
	}
	return
}
