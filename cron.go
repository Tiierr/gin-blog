package main

import (
	"github.com/rayyu03/gin-blog/models"
	"github.com/robfig/cron"
	"log"
	"time"
)

func main()  {
	log.Println("Starting...")

	c := cron.New()
	c.AddFunc("* * * * * *", func() {
		log.Println("Run models.CleanAllTag...")
		models.CleanAllTag()
	})
	c.AddFunc("* * * * * *", func() {
		log.Println("Run models.CleanAllArticle...")
		models.CleanAllArticle()
	})

	c.Start()

	t1 := time.NewTimer(time.Second * 20)
	for {
		select {
		case <-t1.C:
			t1.Reset(time.Second * 10)
		}
	}

}
