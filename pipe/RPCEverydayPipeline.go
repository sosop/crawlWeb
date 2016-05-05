package pipe

import (
	"crawlWeb/config"
	. "crawlWeb/slog"
	"crawlWeb/web"
	"fmt"
	"goMagic/downloader"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	EVERYDAY_ARTICLE = iota
	EVERYDAY_VOICE
	EVERYDAY_BOOK
)

type RPCEverydayPipeline struct {
	mode int
}

func NewRPCEverydayPipeline() *RPCEverydayPipeline {
	return &RPCEverydayPipeline{}
}

func (pipe *RPCEverydayPipeline) Out(p *downloader.Page) error {
	realUrl := config.GetString("rpc.host", true, "http://127.0.0.1:8088/do")
	params := url.Values{"cmd": {"SET"}}
	switch pipe.mode {
	case EVERYDAY_ARTICLE:
		var arts []web.Article
		p.Objects(&arts)
		params["args"] = []string{"everyday_articles," + fmt.Sprint(arts)}
	case EVERYDAY_VOICE:
		var voices []web.Voice
		p.Objects(&voices)
		params["args"] = []string{"everyday_voices," + fmt.Sprint(voices)}
	case EVERYDAY_BOOK:
		var books []web.Book
		p.Objects(&books)
		params["args"] = []string{"everyday_books," + fmt.Sprint(books)}
	}
	resp, err := http.PostForm(realUrl, params)
	if err != nil {
		Logger.Error(err)
		return err
	}
	defer resp.Body.Close()
	ret, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		Logger.Error(err)
		return err
	}
	Logger.Info(string(ret))
	return nil
}
func (pipe *RPCEverydayPipeline) Close() error {
	return nil
}
func (pipe *RPCEverydayPipeline) Mode(mode int) {
	pipe.mode = mode
}
