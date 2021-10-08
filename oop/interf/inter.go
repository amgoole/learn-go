package interf

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type HttpBase interface {
	Get(url string) ([]byte, error)
	Post(url string) ([]byte, error)
}

type Downloader struct {
	Url string
}

func (download Downloader) Get(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	content, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	return content, err
}

type Uploader struct {
	Location string
}

func (upload Uploader) Post(url string) ([]byte, error) {
	fmt.Println("it's a post")
	return nil, nil
}

func (upload *Uploader) Get(url string) ([]byte, error){
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	content, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	return content, err
}

func Spyder(client HttpBase) string {
	content, _ := client.Get("https://www.baidu.com")
	content, _ = client.Post("https://www.baidu.com")
	fmt.Printf("%s\n", string(content))
	return "finished"
}
