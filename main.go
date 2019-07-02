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
	r.CronRun("0 */1 6-22 * * *", func(err error) {
		if err != nil {
			log.Printf("update err: %s\n", err)
			return
		}
		log.Println("update success")
	})
	select {} // 不退出
}
