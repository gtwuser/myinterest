package main

import (
	"fmt"
	"strings"
)

func main() {
	//url := flag.String(`url`, ``, `Url for data inspection`)
	//flag.Parse()
	//resp, err := http.Get(*url)
	//fmt.Println(`Url :`,*url)
	//if err != nil {
	//	log.Fatal("Url could not be reached due to err ", err)
	//}
	//defer resp.Body.Close()
	//
	//all, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	log.Fatal("Url ", *url, "output could not be read ", resp.StatusCode)
	//}
	//file, _ := ioutil.ReadFile(`./data.xml`)
	//content := string(file)
	//re := regexp.MustCompile(`<li.*?>.*?title="(.*?)"`)
	//allString := re.FindAllStringSubmatch(content, -1)
	//for i, r := range allString {
	//	fmt.Println(r[1], i)
	//}
	s := []string{"1808596"}
	fmt.Println(strings.Join(s, `","`))
}
