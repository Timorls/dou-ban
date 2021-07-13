package main

import (
	"fmt"
	"net/http"
	"github.com/PuerkitoBio/goquery"
	"strconv"
	"os"
	"sort"
)

func main(){
	m := make(map[string]string)

	f,_ := os.Create("D:\\Timor\\jwt\\豆瓣Top250.txt")

	defer f.Close()

	for i:=0;i<=250;i=i+25{
		r := "https://movie.douban.com/top250?start="+strconv.Itoa(i)+"&filter="
		url := (r)
		getUrl(url,m)
	}

	getMap(m,f)
}

func getUrl(urls string,m map[string]string){
	name := make(map[int]string)
	rate := make(map[int]string)
	client := &http.Client{}
	req, _ := http.NewRequest("GET", urls, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.90 Safari/537.36")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("error:", err)
	}
	defer resp.Body.Close()
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		fmt.Println("error:", err)
	}
	doc.Find(".title:first-child").Each(func(i int, s *goquery.Selection) {
		content := s.Text()
		name[i] = content
		//_,err := f.WriteString(content+"  ")
		if err != nil {
			fmt.Println(err)
		}
	})

	doc.Find(".rating_num").Each(func(i int, s *goquery.Selection) {
		rating := s.Text()
		rate[i] = rating
		if err != nil {
			fmt.Println(err)
		}
	})

	for i, v := range name {
		m[v]=rate[i]

	}
}

func getMap(m map[string]string,f *os.File){
	type str struct{
		key string
		value string
	}

	var slice []str

	for k,v := range m{
		slice = append(slice,str{k,v})
	}

	sort.Slice(slice,func(i,j int) bool {
		return slice[i].value > slice[j].value
	})

	for i,v := range slice{
		f.WriteString(strconv.Itoa(i+1) + " : " + v.key + "   "+ v.value + "\n")
	}
}
