package ueditor

import (
	"os"
	"io/ioutil"
	"encoding/xml"
	"errors"
)

type configs struct {
	ImagePath string
	FilePath  string
}

var config *configs

func init() {
	var result, err = getConfig()
	if err != nil {
		panic(err)
	}
	if result.FilePath == "" {
		panic(errors.New("附件存放路径未指定"))
	}
	if result.ImagePath == "" {
		panic(errors.New("图片上传路径未指定"))
	}
	config = result
}

func getConfig() (config *configs, err error) {
	file, err := os.Open("config_ue.xml")
	if err != nil {
		return nil, err
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	var result = new(configs)
	err = xml.Unmarshal(data, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
