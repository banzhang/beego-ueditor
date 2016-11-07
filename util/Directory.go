package util

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

//扫描目录下的文件
func GetFiles(paths string, res *[]map[string]interface{}, allow []string) error {
	//paths = strings.TrimLeft(paths, "/") + "admin/upload/"
	fi, err := os.Stat(paths)
	if err != nil {
		return errors.New(paths + ":[" + err.Error() + "]")
	}
	if !fi.IsDir() {
		return errors.New("path not dir:" + paths)
	}
	l, err := ioutil.ReadDir(paths)
	if err != nil {
		return err
	}
	for _, n := range l {
		nm := n.Name()
		if nm == "." || nm == ".." {
			continue
		}
		if !n.IsDir() {
			allowAdd := true
			if len(allow) > 0 {
				allows := strings.Join(allow, "|")
				re := regexp.MustCompile(`\.`)
				allows = re.ReplaceAllString(allows, "")
				re = regexp.MustCompile(allows)
				ma := re.FindAllString(nm, -1)
				if len(ma) < 1 {
					allowAdd = false
				}
			}
			if allowAdd {
				item := map[string]interface{}{
					"url":   paths + nm,
					"mtime": n.ModTime().Unix(),
				}
				*res = append(*res, item)
			}
		} else {
			fmt.Printf("info: base [%s] \n", paths)
			pathss := paths + nm + "/"
			fmt.Printf("info: goin [%s] \n", paths)
			err := GetFiles(pathss, res, allow)
			if err != nil {
				fmt.Printf("err:[%s]", err)
				return err
			}
		}
	}
	return nil
}
