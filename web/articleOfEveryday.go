package web

import (
	. "crawlWeb/slog"
	"strings"

	"goMagic/downloader"

	"github.com/PuerkitoBio/goquery"
)

type (
	EveryDay struct {
		Article
		Voices []Voice
		Books  []Book
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
		Logger.Error(err)
		return
	}
	art := q.Find("#article_show")
	title := art.Find("h1").Text()
	p.PutField("title", strings.TrimSpace(title))
	url := "http://meiriyiwen.com/"
	p.PutField("url", url)
	author := art.Find(".article_author span").Text()
	p.PutField("author", strings.TrimSpace(author))
}

type EveryDayVoiceProcessor struct {
}

func (ev *EveryDayVoiceProcessor) Process(p *downloader.Page) {
	q, err := p.Parser()
	if err != nil {
		Logger.Error(err)
		return
	}
	q.Find(".list_box").Each(func(index int, s *goquery.Selection) {
		url := s.Find(".box_list_img").AttrOr("href", "")
		p.PutField("url", strings.TrimSpace(p.URL+url))
		pic := s.Find(".box_list_img img").AttrOr("src", "")
		p.PutField("pic", strings.TrimSpace(p.URL+pic))
		title := s.Find(".list_author a").Text()
		p.PutField("title", strings.TrimSpace(title))
		author := s.Find(".author_name").Text()
		p.PutField("author", strings.TrimSpace(author))
	})
}

type EveryDayBookProcessor struct {
}

func (ev *EveryDayBookProcessor) Process(p *downloader.Page) {
	q, err := p.Parser()
	if err != nil {
		Logger.Error(err)
		return
	}
	q.Find(".book-list li").Each(func(index int, s *goquery.Selection) {
		url := s.Find(".book-bg").AttrOr("href", "")
		p.PutField("url", strings.TrimSpace(p.URL+url))
		pic := s.Find(".book-bg img").AttrOr("src", "")
		p.PutField("pic", strings.TrimSpace(p.URL+pic))
		title := s.Find(".book-name a").Text()
		p.PutField("title", strings.TrimSpace(title))
		author := s.Find(".book-author").Text()
		p.PutField("author", strings.TrimSpace(author))
	})
}
