package main

import (
	"encoding/base64"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

func getAllBase64ImageFromMarkDownFile(filePaths []string) {
	//text := "(data:image/JPG;base64,/9j/4AAQSkZJRgABAQAAAQABAAD/2wBDAAoHBwkHBgoJCAkLCwoMDxkQDw4ODx4WFxIZJCAmJSMgIyIoLTkwKCo2KyIjMkQyNjs9QEBAJjBGS0U)"
	base64ImgRegexpPattern, err := regexp.Compile("data:image/(.*?);base64,(.*?)\\)")
	if err != nil {
		panic("1")
	}

	for _, rawPath := range filePaths {
		_, rawFileName := filepath.Split(rawPath)

		rawFileExt := filepath.Ext(rawFileName)
		if rawFileExt != ".md" {
			println("not a .md file:", rawPath)
			continue
		}

		currentPath, _ := os.Getwd()
		savePath := filepath.Join(currentPath, strings.TrimSuffix(rawFileName, rawFileExt))

		err := os.MkdirAll(savePath, 0666)
		if err != nil {
			println("file to create path:", savePath, "reason:", err.Error())
			return
		}

		rawFileBytes, err := ioutil.ReadFile(rawPath)
		if err != nil {
			println("fail to read file:", rawPath, "reason:", err.Error())
			continue
		}

		result := base64ImgRegexpPattern.FindAllStringSubmatch(string(rawFileBytes), -1)
		if result == nil {
			println("no image found")
			continue
		}

		for index, a := range result {
			//println(index, a[1], a[2])
			ext, base64String := a[1], a[2]
			ext = strings.ToLower(ext)
			filePath := filepath.Join(savePath, strconv.Itoa(index+1)+"."+ext)

			fileBytes, err := base64.StdEncoding.DecodeString(base64String)

			if err != nil {
				println("图片文本文解码错误!", filePath, "reason: ", err.Error())
				continue
			}

			err = ioutil.WriteFile(filePath, fileBytes, 0666)
			if err != nil {
				println("save file error: ", filePath, "reason: ", err.Error())
				continue
			}
		}
	}
}

func main() {
	args := os.Args
	if len(args) <= 1 {
		println("need command")
		return
	}

	command := os.Args[1]
	if command == "getAllBase64ImageFromMarkDownFile" {

		if len(os.Args) == 2 {
			println("need file path")
			return
		}

		//var rawFilePaths []string
		for i, argPath := range os.Args[2:] {
			println(i, argPath)
		}
		getAllBase64ImageFromMarkDownFile(os.Args[2:])
	}

}
