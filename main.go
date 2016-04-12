package main

import (
	"crawlWeb/web"
	"goMagic/core"
	"goMagic/pipe"
)

func main() {
	core.NewMagic("everydayArticle", &web.EveryDayArticleProcessor{}).AddURL("http://meiriyiwen.com/").SetThread(1).SetPipeline(pipe.NewFilePipeline("article")).SetOutMode(pipe.MAPS).Run()
	core.NewMagic("everydayVoice", &web.EveryDayVoiceProcessor{}).AddURL("http://voice.meiriyiwen.com/").SetThread(1).SetPipeline(pipe.NewFilePipeline("voice")).SetOutMode(pipe.MAPS).Run()
}
