package main

import (
	"github.com/iikira/ippush/internal/mytask"
	"github.com/iikira/ippush/syncrunner"
	"log"
)

func main() {
	r, err := syncrunner.NewRunner()
	if err != nil {
		log.Fatalln(err)
	}

	r.SetTask(mytask.NewMyRunTask())

	c := r.PrepareCron("@every 1m", func(err error) {
		if err != nil {
			log.Printf("update err: %s\n", err)
			return
		}
		log.Println("update success")
	})
	c.Run()
}
