package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"time"
)

type Article struct {
	Picture_name string   //public for passing to DownloadPicture()
	title        string
	subtitle     string
	date         string
	picture_tags []string
	article_tags []string
}

func getInput(hint string) string {
	var input string
	for {
		fmt.Print(hint)
		fmt.Scanln(&input)
		if input != "" {  //make sure to get an valid input
			break
		}
	}
	return input

}
func getTags(hint string) []string {
	str := getInput(hint)
	return strings.Split(str, ",")
}
func (it *Article) GetInfo() {
	it.date = time.Now().Format("2006-01-02")
	it.picture_tags = getTags("请输入图片标签（以逗号分隔）:")
	it.Picture_name=strings.Join(it.picture_tags ,"-") //high efficiency
	it.Picture_name=strings.Join([]string{"post-bg-",it.Picture_name,".jpg"} ,"")
	it.title = getInput("请输入文章标题 :")
	it.subtitle = getInput("请输入文章副标题 :")
	it.article_tags = getTags("请输入文章标签（以逗号分隔）:")
}
// PrintFile() will NOT check if Article's properties are invalid
func (it *Article) PrintFile() {
	fileaddr := strings.Join([]string{ARTICLE_PATH, it.date, "-", it.title,".md" }, "")

	var tmp string
	tmp=strings.Replace(CONTENT,"TITLE" , it.title, 1)
	tmp=strings.Replace(tmp,"SUBTITLE" , it.subtitle, 1)
	tmp=strings.Replace(tmp,"DATE" , it.date, 1)//simple but low efficiency
	tmp=strings.Replace(tmp,"PICTURE_NAME" , it.Picture_name, 1)
	for _,tag:=range it.article_tags{
		tmp=strings.Replace(tmp,"ARTICLE_TAG" ,tag, 1)
	}
	tmp=strings.ReplaceAll(tmp,"- ARTICLE_TAG" ,"")   //delete extra tag template
	ioutil.WriteFile(fileaddr, []byte(tmp), 0644)

}
