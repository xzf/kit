package fileX

import (
	"os"
)

func Write(path string, content []byte) {
	f, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	_, err = f.Write(content)
	if err != nil {
		panic(err)
	}
	err = f.Close()
	if err != nil {
		panic(err)
	}
	return
}

func WriteStr(path string, content string) {
	Write(path, []byte(content))
}

func Append(path string, content []byte) {
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND, os.ModeAppend)
	if err != nil {
		panic(err)
	}
	_, err = f.Write(content)
	if err != nil {
		panic(err)
	}
	err = f.Close()
	if err != nil {
		panic(err)
	}
	return
}

func AppendStr(path string, content string) {
	Append(path, []byte(content))
}

func Read(path string) []byte {
	result, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return result
}

func ReadStr(path string) string {
	return string(Read(path))
}
