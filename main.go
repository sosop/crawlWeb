package main

import (
	"crawlWeb/pipe"
	"crawlWeb/web"

	"goMagic/core"
)

func main() {
	core.NewMagic("everydayArticle", &web.EveryDayArticleProcessor{}).AddURL("http://meiriyiwen.com").SetThread(1).SetPipeline(pipe.NewRPCEverydayPipeline()).SetOutMode(pipe.EVERYDAY_ARTICLE).Run()
	core.NewMagic("everydayVoice", &web.EveryDayVoiceProcessor{}).AddURL("http://voice.meiriyiwen.com").SetThread(1).SetPipeline(pipe.NewRPCEverydayPipeline()).SetOutMode(pipe.EVERYDAY_VOICE).Run()
	core.NewMagic("everydayBook", &web.EveryDayBookProcessor{}).AddURL("http://book.meiriyiwen.com").SetThread(1).SetPipeline(pipe.NewRPCEverydayPipeline()).SetOutMode(pipe.EVERYDAY_BOOK).Run()
}
