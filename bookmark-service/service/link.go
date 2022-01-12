package service

import (
	"context"
	"github.com/MenciusCheng/bookmark-manager/bookmark-service/model"
	"regexp"
	"strings"
)

func (s *Service) GetLinkList(ctx context.Context) (infos []model.Link, err error) {
	return s.dao.GetLinkList(ctx)
}

func (s *Service) InitLinkFromString(ctx context.Context, str string) (err error) {
	links := make([]model.Link, 0)

	split := strings.Split(str, "\n")
	for _, s := range split {
		link, match := strToLinkM(s)
		if match {
			links = append(links, link)
		}
	}
	if err = s.dao.CreateLinkList(ctx, links); err != nil {
		return err
	}

	return nil
}

func strToLink(str string) bool {
	linkReg := regexp.MustCompile(`^\s*<DT><A HREF=".+</A>`)
	return linkReg.MatchString(str)
}

func strToLinkM(str string) (info model.Link, match bool) {
	linkReg := regexp.MustCompile(`^\s*<DT><A HREF=".+</A>`)
	if !linkReg.MatchString(str) {
		return
	}

	hrefReg := regexp.MustCompile(`HREF="([a-zA-z]+://[^\s]*)"`)
	hrefSub := hrefReg.FindStringSubmatch(str)

	nameReg := regexp.MustCompile(`">(.+?)</A>`)
	nameSub := nameReg.FindStringSubmatch(str)

	iconReg := regexp.MustCompile(`ICON="([^\s]*)"`)
	iconSub := iconReg.FindStringSubmatch(str)

	info.Name = nameSub[1]
	info.URL = hrefSub[1]
	if len(iconSub) > 1 {
		info.Icon = iconSub[1]
	}
	match = true
	return
}
