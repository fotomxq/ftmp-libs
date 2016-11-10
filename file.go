//文件操作模块
package ftmplibs

import (
	"io/ioutil"
	"os"
	"bytes"
	"crypto/sha1"
)

//文件类结构
type FileOperate struct{
}

//创建新的文件夹
//支持多级创建
func (File *FileOperate) CreateDir(src string) (bool,error) {
	err := os.MkdirAll(src,os.ModePerm)
	if err != nil{
		return false,err
	}
	return true,nil
}

//创建文件
func (File *FileOperate) CreateFile(src string) bool{
	_,err := os.Create(src)
	if err != nil{
		return true
	}
	return false
}

//读取文件
func (File *FileOperate) ReadFile(src string) ([]byte,error){
	fd, fdErr := os.Open(src)
	if fdErr != nil {
		return nil,fdErr
	}
	defer fd.Close()
	c, cErr := ioutil.ReadAll(fd)
	if cErr != nil {
		return nil,cErr
	}
	return c,nil
}

//写入文件
func (File *FileOperate) WriteFile(src string, content []byte) (bool,error) {
	err := ioutil.WriteFile(src, content, os.ModeAppend)
	if err != nil {
		return false,err
	}
	return true,nil
}

//追加写入文件
func (File *FileOperate) WriteFileAppend(src string, content []byte) (bool,error){
	if File.IsFile(src) == false{
		writeBool,writeErr := File.WriteFile(src, content)
		return writeBool,writeErr
	}
	fileContent, fcErr := File.ReadFile(src)
	if fcErr != nil{
		return false,fcErr
	}
	s := [][]byte{
		fileContent,
		content,
	}
	sep := []byte("")
	var newContent []byte = bytes.Join(s,sep)
	writeBool2,writeErr2 := File.WriteFile(src,newContent)
	return writeBool2,writeErr2
}

//修改文件或文件夹名称
//可用于修改路径，即剪切
func (File *FileOperate) EditFileName(src string, newName string) (bool,error) {
	err := os.Rename(src, newName)
	if err != nil {
		return true,err
	}
	return false,nil
}

//删除文件
func (File *FileOperate) DeleteFile(src string) bool {
	err := os.RemoveAll(src)
	if err != nil {
		return true
	}
	return false
}

//判断路径是否存在
func (File *FileOperate) IsExist(src string) bool{
	_, err := os.Stat(src)
	return err == nil || os.IsExist(err)
}

//判断是否为文件
func (File *FileOperate) IsFile(src string) bool {
	info, err := os.Stat(src)
	if err != nil{
		return false
	}
	return !info.IsDir()
}

//判断是否为文件夹
func (File *FileOperate) IsFolder(src string) bool {
	info, err := os.Stat(src)
	if err != nil{
		return false
	}
	return info.IsDir()
}

//获取文件列表
func (File *FileOperate) GetFileList(src string) string {
	return ""
}

//获取文件大小
func (File *FileOperate) GetFileSize(src string) int64 {
	info, err := os.Stat(src)
	if err != nil{
		return 0
	}
	return info.Size()
}

//获取文件信息
func (File *FileOperate) GetFileInfo(src string) (os.FileInfo ,error) {
	info, err := os.Stat(src)
	return info,err
}

//计算文件sha1值
func (File *FileOperate) GetFileSha1(src string) (string,error){
	content,err := File.ReadFile(src)
	if err != nil{
		return "",err
	}
	if content != nil{
		sha := sha1.New()
		sha.Write(content)
		res := sha.Sum(nil)
		return string(res),nil
	}
	return "",nil
}