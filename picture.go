package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

//download the background picture
func GetUrl() string {
	return getInput("请输入图片URL :")
}
func DownloadPicture(url, name string) {

	res, err := http.Get(url)
	if err != nil {
		fmt.Println("Download Failed !!!")
		return
	} else {
		fmt.Println("Download Successfully !!!")
	}
	defer res.Body.Close()

	file, err := os.Create(PICTURE_PATH + name)
	if err != nil {
		panic(err)
	}

	written, _ := io.Copy(file, res.Body)
	fmt.Printf("%s Saved\n", name)
	fmt.Printf("Total Size: %d Bytes\n", written)
}
