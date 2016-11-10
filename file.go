//文件操作模块
<<<<<<< HEAD
//该包直接调用函数即可
=======
>>>>>>> 2865ccf63334543b43a4ded793e8ed3dd3a456b2
package ftmplibs

import (
	"io/ioutil"
	"os"
	"bytes"
	"crypto/sha1"
)

<<<<<<< HEAD
//创建新的文件夹
//支持多级创建
func CreateDir(src string) (bool,error) {
=======
//文件类结构
type FileOperate struct{
}

//创建新的文件夹
//支持多级创建
func (File *FileOperate) CreateDir(src string) (bool,error) {
>>>>>>> 2865ccf63334543b43a4ded793e8ed3dd3a456b2
	err := os.MkdirAll(src,os.ModePerm)
	if err != nil{
		return false,err
	}
	return true,nil
}

//创建文件
<<<<<<< HEAD
func CreateFile(src string) bool{
=======
func (File *FileOperate) CreateFile(src string) bool{
>>>>>>> 2865ccf63334543b43a4ded793e8ed3dd3a456b2
	_,err := os.Create(src)
	if err != nil{
		return true
	}
	return false
}

//读取文件
<<<<<<< HEAD
func ReadFile(src string) ([]byte,error){
=======
func (File *FileOperate) ReadFile(src string) ([]byte,error){
>>>>>>> 2865ccf63334543b43a4ded793e8ed3dd3a456b2
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
<<<<<<< HEAD
func WriteFile(src string, content []byte) (bool,error) {
=======
func (File *FileOperate) WriteFile(src string, content []byte) (bool,error) {
>>>>>>> 2865ccf63334543b43a4ded793e8ed3dd3a456b2
	err := ioutil.WriteFile(src, content, os.ModeAppend)
	if err != nil {
		return false,err
	}
	return true,nil
}

//追加写入文件
<<<<<<< HEAD
func WriteFileAppend(src string, content []byte) (bool,error){
	if IsFile(src) == false{
		writeBool,writeErr := WriteFile(src, content)
		return writeBool,writeErr
	}
	fileContent, fcErr := ReadFile(src)
=======
func (File *FileOperate) WriteFileAppend(src string, content []byte) (bool,error){
	if File.IsFile(src) == false{
		writeBool,writeErr := File.WriteFile(src, content)
		return writeBool,writeErr
	}
	fileContent, fcErr := File.ReadFile(src)
>>>>>>> 2865ccf63334543b43a4ded793e8ed3dd3a456b2
	if fcErr != nil{
		return false,fcErr
	}
	s := [][]byte{
		fileContent,
		content,
	}
	sep := []byte("")
	var newContent []byte = bytes.Join(s,sep)
<<<<<<< HEAD
	writeBool2,writeErr2 := WriteFile(src,newContent)
=======
	writeBool2,writeErr2 := File.WriteFile(src,newContent)
>>>>>>> 2865ccf63334543b43a4ded793e8ed3dd3a456b2
	return writeBool2,writeErr2
}

//修改文件或文件夹名称
//可用于修改路径，即剪切
<<<<<<< HEAD
func EditFileName(src string, newName string) (bool,error) {
=======
func (File *FileOperate) EditFileName(src string, newName string) (bool,error) {
>>>>>>> 2865ccf63334543b43a4ded793e8ed3dd3a456b2
	err := os.Rename(src, newName)
	if err != nil {
		return true,err
	}
	return false,nil
}

//删除文件
<<<<<<< HEAD
func DeleteFile(src string) bool {
=======
func (File *FileOperate) DeleteFile(src string) bool {
>>>>>>> 2865ccf63334543b43a4ded793e8ed3dd3a456b2
	err := os.RemoveAll(src)
	if err != nil {
		return true
	}
	return false
}

//判断路径是否存在
<<<<<<< HEAD
func IsExist(src string) bool{
=======
func (File *FileOperate) IsExist(src string) bool{
>>>>>>> 2865ccf63334543b43a4ded793e8ed3dd3a456b2
	_, err := os.Stat(src)
	return err == nil || os.IsExist(err)
}

//判断是否为文件
<<<<<<< HEAD
func IsFile(src string) bool {
=======
func (File *FileOperate) IsFile(src string) bool {
>>>>>>> 2865ccf63334543b43a4ded793e8ed3dd3a456b2
	info, err := os.Stat(src)
	if err != nil{
		return false
	}
	return !info.IsDir()
}

//判断是否为文件夹
<<<<<<< HEAD
func IsFolder(src string) bool {
=======
func (File *FileOperate) IsFolder(src string) bool {
>>>>>>> 2865ccf63334543b43a4ded793e8ed3dd3a456b2
	info, err := os.Stat(src)
	if err != nil{
		return false
	}
	return info.IsDir()
}

//获取文件列表
<<<<<<< HEAD
func GetFileList(src string) string {
=======
func (File *FileOperate) GetFileList(src string) string {
>>>>>>> 2865ccf63334543b43a4ded793e8ed3dd3a456b2
	return ""
}

//获取文件大小
<<<<<<< HEAD
func GetFileSize(src string) int64 {
=======
func (File *FileOperate) GetFileSize(src string) int64 {
>>>>>>> 2865ccf63334543b43a4ded793e8ed3dd3a456b2
	info, err := os.Stat(src)
	if err != nil{
		return 0
	}
	return info.Size()
}

//获取文件信息
<<<<<<< HEAD
func GetFileInfo(src string) (os.FileInfo ,error) {
=======
func (File *FileOperate) GetFileInfo(src string) (os.FileInfo ,error) {
>>>>>>> 2865ccf63334543b43a4ded793e8ed3dd3a456b2
	info, err := os.Stat(src)
	return info,err
}

//计算文件sha1值
<<<<<<< HEAD
func GetFileSha1(src string) (string,error){
	content,err := ReadFile(src)
=======
func (File *FileOperate) GetFileSha1(src string) (string,error){
	content,err := File.ReadFile(src)
>>>>>>> 2865ccf63334543b43a4ded793e8ed3dd3a456b2
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