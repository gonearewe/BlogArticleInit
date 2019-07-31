package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"time"
)

type Article struct {
	title        string
	subtitle     string
	date         string
	picture_name string
	picture_tags []string
	article_tags []string
}

func getInput(hint string) string {
	var input string
	for {
		fmt.Print(hint)
		fmt.Scanln(&input)
		if input != "" {
			break
		}
	}
	return input

}
func getTags(hint string) []string {
	str := getInput(hint)
	println(str)
	return strings.Split(str, ",")
}
func (it *Article) GetInfo() {
	it.date = time.Now().Format("2006-01-02")
	it.picture_tags = getTags("请输入图片标签（以逗号分隔）:")
	fmt.Println(it.picture_tags)
	it.title = getInput("请输入文章标题 :")
	it.subtitle = getInput("请输入文章副标题 :")
	it.article_tags = getTags("请输入文章标签（以逗号分隔）:")
}
func (it *Article) PrintFile() {
	fileaddr := strings.Join([]string{ARTICLE_PATH, it.date, "-", it.title,".md" }, "")
	ioutil.WriteFile(fileaddr, []byte(CONTENT), 0644)

}
