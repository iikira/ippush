package main

import (
	"flag"
	"github.com/iikira/ippush/syncrunner"
	"log"
)

var (
	zoneName string
	aName    string
)

func init() {
	flag.StringVar(&zoneName, "z", "example.com", "zone name")
	flag.StringVar(&aName, "a", "a", "a name")
	flag.Parse()
}

func main() {
	r, err := syncrunner.NewRunner()
	if err != nil {
		log.Fatalln(err)
	}

	r.SetTask(syncrunner.NewStaticTask(zoneName, aName))

	c := r.PrepareCron("@every 1m", true, func(err error) {
		if err != nil {
			log.Printf("update err: %s\n", err)
			return
		}
		log.Println("update success")
	})
	c.Run()
}
