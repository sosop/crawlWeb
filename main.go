package main

import (
	"crawlWeb/pipe"
	"crawlWeb/web"
	"os"
	"os/signal"
	"syscall"

	"goMagic/core"

	"github.com/robfig/cron"
)

func main() {
	sc := make(chan os.Signal, 1)
	signal.Notify(sc,
		os.Kill,
		os.Interrupt,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)
	c := cron.New()
	c.AddFunc("0 */30 * * * *", func() {
		crawlEveryday()
	})
	//c.AddFunc("@hourly",      func() { fmt.Println("Every hour") })
	//c.AddFunc("@every 1h30m", func() { fmt.Println("Every hour thirty") })
	c.Start()
	<-sc
	c.Stop()
}

func crawlEveryday() {
	// defer wg.Done()
	// wg.Add(1)
	core.NewMagic("everydayArticle", &web.EveryDayArticleProcessor{}).AddURL("http://meiriyiwen.com").SetThread(1).SetPipeline(pipe.NewRPCEverydayPipeline()).SetOutMode(pipe.EVERYDAY_ARTICLE).Run()
	core.NewMagic("everydayVoice", &web.EveryDayVoiceProcessor{}).AddURL("http://voice.meiriyiwen.com").SetThread(1).SetPipeline(pipe.NewRPCEverydayPipeline()).SetOutMode(pipe.EVERYDAY_VOICE).Run()
	core.NewMagic("everydayBook", &web.EveryDayBookProcessor{}).AddURL("http://book.meiriyiwen.com").SetThread(1).SetPipeline(pipe.NewRPCEverydayPipeline()).SetOutMode(pipe.EVERYDAY_BOOK).Run()
}
