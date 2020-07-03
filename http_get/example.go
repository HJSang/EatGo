package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	url := "https://www.google.com/search?q=google&sxsrf=ALeKk02mc4588E4c2tvgqlEkpjPox7Gh_w:1593741086976&source=lnms&tbm=nws&sa=X&ved=2ahUKEwirlema_K_qAhW-hHIEHROjC-sQ_AUoAXoECB8QAw&biw=1440&bih=820" 
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	robots, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", robots)
}
