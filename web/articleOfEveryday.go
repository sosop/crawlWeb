package web

import (
	. "crawlWeb/slog"
	"goMagic/downloader"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type (
	EveryDay struct {
		Article
		Voice
		Book
	}
	Article struct {
		Title  string `json:"title"`
		URL    string `json:"url"`
		Author string `json:"author"`
	}
	Voice struct {
		Title  string `json:"title"`
		URL    string `json:"url"`
		Author string `json:"author"`
		PicURL string `json:"pic"`
	}
	Book struct {
		Title  string `json:"title"`
		URL    string `json:"url"`
		Author string `json:"author"`
		PicURL string `json:"pic"`
	}
)

type EveryDayArticleProcessor struct {
}

func (ea *EveryDayArticleProcessor) Process(p *downloader.Page) {
	q, err := p.Parser()

	if err != nil {
		Slogger.Error(err)
		return
	}
	art := q.Find("#article_show")
	title := art.Find("h1").Text()
	p.PutField("Title", strings.TrimSpace(title))
	url := "http://meiriyiwen.com/"
	p.PutField("URL", url)
	author := art.Find(".article_author span").Text()
	p.PutField("Author", strings.TrimSpace(author))
}

type EveryDayVoiceProcessor struct {
}

func (ev *EveryDayVoiceProcessor) Process(p *downloader.Page) {
	q, err := p.Parser()
	if err != nil {
		Slogger.Error(err)
		return
	}
	q.Find(".list_box").Each(func(index int, s *goquery.Selection) {
		url := s.Find(".box_list_img").AttrOr("href", "")
		p.PutField("URL", strings.TrimSpace(url))
		pic := s.Find(".box_list_img img").AttrOr("src", "")
		p.PutField("PicURL", strings.TrimSpace(pic))
		title := s.Find(".list_author a").Text()
		p.PutField("Title", strings.TrimSpace(title))
		author := s.Find(".author_name").Text()
		p.PutField("author", strings.TrimSpace(author))
	})
}
