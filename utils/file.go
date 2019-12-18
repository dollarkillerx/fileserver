/**
 * @Author: DollarKillerX
 * @Description: file 文件相关的utils
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 下午5:08 2019/12/18
 */
package utils

import "os"

// 判断文件夹是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// 如果文件夹不存在就会创建
func DirPing(path string) error {
	b, e := PathExists(path)
	if e != nil {
		return e
	}
	if !b {
		e := os.MkdirAll(path, 00777)
		if e != nil {
			return e
		}
	}
	return nil
}
