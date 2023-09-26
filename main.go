package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/cetinboran/yavuzlarscraper/models"
	"github.com/cetinboran/yavuzlarscraper/scraper"
)

func main() {
	// http://localhost/Yavuzlar_TODO_PHP/src/register.php
	res, err := http.Get("http://localhost/myBlog/about.php")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	scraper, err := scraper.BodyReader(res.Body)
	if err != nil {
		fmt.Println(err)
	}

	// AutoSave Added.
	scraper.SetConfig(&models.Config{
		AutoSave: false,
	})

	tag := models.TagInit("div")
	tag.SetClasses("description")

	tag1 := models.TagInit("div")
	tag1.SetClasses("title")

	scraper.Find(*tag).Each(func(i int, name string) {
		fmt.Println(i, name)
	})
	scraper.Find(*tag1)

}
