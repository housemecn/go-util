package goutil

import (
	"encoding/json"
	"os"
	
	"github.com/housemecn/goutil/file"
)

// JSONFileToBytes 从json文件中转换为[]byte
func JSONFileToBytes(filepath string) ([]byte, error) {
	open, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	fileInfo, err := open.Stat()
	if err != nil {
		return nil, err
	}

	buf := make([]byte, fileInfo.Size())
	_, err = open.Read(buf)
	if err != nil {
		return nil, err
	}
	return buf, nil
}

// JSONBytesToStruct json []byte 转换为 struct
func JSONBytesToStruct(data []byte, structObj interface{}) error {
	err := json.Unmarshal(data, structObj)
	if err != nil {
		return err
	}
	return nil
}

// JSONBytesToFile json []byte 写入文件
func JSONBytesToFile(data []byte, filepath string) error {
	err := file.BytesToFile(data, filepath)
	if err != nil {
		return err
	}
	return nil
}
