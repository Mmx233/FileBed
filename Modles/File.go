package Modles

import "os"

type FileStruct struct {}
var File FileStruct

func (*FileStruct)Exists(path string) bool {
	_, err := os.Stat(path)    //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

func (*FileStruct)IsDir(path string)bool{
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}