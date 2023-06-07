package common

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"regexp"
	"time"
)

// 获取字符串中指定的子串
//  @param pattern 	正则匹配公式
//  @param s 		匹配串
//  @return map		母串对应子串的map
func GetStringAssignfield(pattern string, s []string) map[string]string {
	m := make(map[string]string)
	reg, _ := regexp.Compile(pattern)

	for _, ss := range s {
		substr := reg.FindString(ss)
		if len(substr) > 0 {
			m[ss] = substr
		}
	}
	return m
}

// 获取指定目录下所有的文件名称, 支持后缀过滤
//  @param pathname 	路径
//  @param ext 		文件后缀
//  @return 
func GetDirFile(pathname, ext string) (s []string, e error) {
	rd, err := ioutil.ReadDir(pathname)
	if err != nil {
		fmt.Printf("读取目录失败: %s \n", err.Error())
		return s, err
	}

	var extflag bool
	if len(ext) == 0 {
		extflag = false
	} else {
		extflag = true
	}

	for _, file := range rd {
		if !file.IsDir() {
			if extflag && (ext != path.Ext(file.Name())) {
				continue
			}
			fullName := pathname + "/" + file.Name()
			s = append(s, fullName)
		}
	}
	return s, nil
}

type TimeStat struct {
	t time.Time
}

func (ts *TimeStat) TimeStatInit() {
	ts.t = time.Now()
}

func (ts *TimeStat) TimeStatShow() {
	e := time.Now().Sub(ts.t)
	fmt.Printf("执行时长：%v \n", e)
}

func PressKeyExit() {
	fmt.Printf("请按任意键退出... \n")
	k := make([]byte, 1)
	os.Stdin.Read(k)
}

// 打印错误信息并退出
func ShowErrExit(err error, errno int) {
	if err != nil {
		fmt.Printf("报错信息(%v) \n", err)
		os.Exit(errno)
	}
}

// 打印错误信息
func ShowErr(err error) {
	if err != nil {
		fmt.Printf("报错信息(%v) \n", err)
	}
}



