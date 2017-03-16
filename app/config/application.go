package config

import (
	"fmt"
	"gopkg.in/gin-gonic/gin.v1"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"errors"
)

var Settings map[string]string

func loadYml() {
	var filename string
	var ok bool
	//当前执行进程中获取相对路径
	filename, err := os.Executable()
	errPanic(err)
	defer func() {
		if err:=recover();err!=nil{
			_, filename, _, ok = runtime.Caller(0)
			if !ok {
				panic("No caller information")
			}
			readFile(filename)
		}

	}()
	readFile(filename)
}

func readFile(filename string) {
	exPath := path.Dir(filename)
	setting_file, _ := filepath.Abs(exPath + "/settings.yml")
	fmt.Fprintf(gin.DefaultWriter, "read setting_file===> %s \r\n", setting_file)
	yamlFile, err := ioutil.ReadFile(setting_file)
	err = yaml.Unmarshal(yamlFile, &Settings)
	if len(Settings) == 0 {
		err = errors.New("settings empty")
	}
	errPanic(err)
}

func errPanic(err error) {
	if err != nil {
		fmt.Println("error: ",err)
		panic(err)
	}
}
