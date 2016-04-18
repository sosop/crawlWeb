package main

import (
	"crawlWeb/site"
	"goMagic/core"
	"goMagic/pipe"
)

func main() {
	core.NewMagic("everydayArticle", &web.EveryDayArticleProcessor{}).AddURL("http://meiriyiwen.com/").SetThread(1).SetPipeline(pipe.NewFilePipeline("article.r")).SetOutMode(pipe.MAPS).Run()
	core.NewMagic("everydayVoice", &web.EveryDayVoiceProcessor{}).AddURL("http://voice.meiriyiwen.com/").SetThread(1).SetPipeline(pipe.NewFilePipeline("voice.r")).SetOutMode(pipe.MAPS).Run()
	core.NewMagic("everydayBook", &web.EveryDayBookProcessor{}).AddURL("http://book.meiriyiwen.com/").SetThread(1).SetPipeline(pipe.NewFilePipeline("book.r")).SetOutMode(pipe.MAPS).Run()
}
