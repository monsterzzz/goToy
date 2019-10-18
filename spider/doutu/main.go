package doutu

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"runtime"
	"strings"
	"sync"
	"time"
)

const (
	SaveDir = "/spider/img"
	StartPage = 3 // must be greater than 1
	EndPage = 100
	MaxRunning = 20
)

var WG sync.WaitGroup
var ch = make(chan bool,MaxRunning)

func checks(e error){
	if e != nil {
		var buf [2 << 10]byte
		fmt.Println(string(buf[:runtime.Stack(buf[:], true)]))
		log.Fatal(e)
	}

}


func SpiderGetMethod(url string) *http.Response{
	client := &http.Client{}
	resq, err := http.NewRequest("GET", url, nil)
	checks(err)
	resq.Header.Set("user-agent","Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/71.0.3578.98 Safari/537.36")
	resq.Header.Set("upgrade-insecure-requests","1")
	resq.Header.Set("Referer","https://www.doutula.com/article/list/?page=5")
	resp, err := client.Do(resq)
	checks(err)
	return resp
}

func Run() {
	startUrl := "https://www.doutula.com/article/list/?page=%d"

	//startId := 3

	for i := StartPage ; i < EndPage; i++ {
		sLi := IndexPageParser(fmt.Sprintf(startUrl,i))
		//fmt.Println(fmt.Sprintf(startUrl,i))
		fmt.Println(sLi)
		for _,v := range sLi{
			WG.Add(1)
			go ImageGroupDownload(v)
		}
	}

	WG.Wait()

	//err:= os.Mkdir("./img/出门出门后表情",0777)
	//fmt.Println(err)

	//b1 := []rune{46,47,105,109,103,47,20986,38376,21069,8,20986,38376,21518,34920,24773}
	//fmt.Println(string([]rune{21069}))
	//fmt.Println(string([]rune{8}))
	//fmt.Println(string([]rune{21069,8}))
	//fmt.Println(string([]rune{21069,8,20986}))
	//fmt.Println([]rune("./img/出门前出门后表情"))
	//for i := 0; i < 200 ; i++ {
	//	fmt.Println(i,"*",string(rune(i)))
	//}

}


func MakeDoc(url string) *goquery.Document {
	resp := SpiderGetMethod(url)
	doc,err := goquery.NewDocumentFromReader(resp.Body)
	checks(err)
	return doc
}


func IndexPageParser(url string) []string{
	sLi := make([]string,0)
	doc := MakeDoc(url)
	doc.Find("a.list-group-item.random_list").Each(func(i int, selection *goquery.Selection) {
		a,_ := selection.Attr("href")
		sLi = append(sLi,a)
	})
	return sLi
}

func SonPageParser(url string) map[string]interface{} {

	s := make(map[string]interface{})
	s["href"] = make([]string,0)
	doc := MakeDoc(url)
	title := doc.Find("title").Text()
	tLi := strings.Split(title," ")
	s["name"] = tLi[0]
	doc.Find("img[referrerpolicy]").Each(func(i int, selection *goquery.Selection) {
		href,_ := selection.Attr("src")
		s["href"] = append(s["href"].([]string),href)
	})
	return s
}

func ImageDownloader(url string,saveDir string){
	defer func(){
		<- ch
	}()
	defer WG.Done()
	ch <- false
	imgNameLi := strings.Split(url,"/")
	imgName := imgNameLi[len(imgNameLi) - 1]
	resp, err := http.Get(url)
	checks(err)
	bytes, err := ioutil.ReadAll(resp.Body)
	checks(err)
	f,err := os.Create(saveDir + "/" + imgName )
	if err != nil{
		f,_ = os.Create(saveDir + "/" + fmt.Sprintf("unknown_%d",rand.Intn(64)) )
	}
	defer f.Close()
	f.Write(bytes)
	time.Sleep(500)


}

func ImageGroupDownload(url string){
	m := SonPageParser(url)
	pwd := "."
	name := m["name"].(string)
	runeName := []rune(name)
	for i := len(runeName) - 1 ; i >= 0 ; i-- {
		if runeName[i] < 47{
			runeName = append(runeName[:i], runeName[i+1:]...)
		}
	}
	name = string(runeName)
	for _,v := range []string{"/","\\","?","*","？","<",">","|",":","："}{
		name = strings.ReplaceAll(name,v,"")
	}
	dirName := strings.TrimSpace(pwd  + SaveDir + "/" + name)

	fmt.Println(dirName,url)
	//for i,v := range dirName{
	//	fmt.Println(i,"*",string(v),v)
	//}
	if _,err := os.Stat(dirName);os.IsNotExist(err){
		e:= os.MkdirAll(dirName,os.ModePerm)
		if e != nil {
			dirName = strings.TrimSpace(pwd  + SaveDir + "/" + fmt.Sprintf("unknown_%d",rand.Intn(64)))
		}
	}

	for _,v := range m["href"].([]string){
		WG.Add(1)
		go ImageDownloader(v,dirName)
	}
	WG.Done()
}