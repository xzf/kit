package fileX

import "os"

func MustReadFile(filePath string) []byte {
	content, err := ReadFile(filePath)
	if err != nil {
		panic("[akpf4kfdhm] " + err.Error())
	}
	return content
}

func ReadFile(filePath string) ([]byte, error) {
	return os.ReadFile(filePath)
}

func WriteFile(filePath string, content []byte) error {
	return os.WriteFile(filePath, content, 0777)
}

func MustWriteFile(filePath string, content []byte) {
	err := WriteFile(filePath, content)
	if err != nil {
		panic("[fbyzb8jpzv] " + err.Error())
	}
}

func AppendFile(filePath string, content []byte) error {
	f, err := os.OpenFile(filePath, os.O_APPEND, 0777)
	if err != nil {
		return err
	}
	_, err = f.Write(content)
	if err != nil {
		return err
	}
	return nil
}

func MustAppendFile(filePath string, content []byte) {
	err := AppendFile(filePath, content)
	if err != nil {
		panic("[brjz3q4d68] " + err.Error())
	}
}
