package main

func main() {
	var f Article
	url:=GetUrl()
	f.GetInfo()
	DownloadPicture(url, f.Picture_name)
	f.PrintFile()
	f.OpenWithCode()
}
