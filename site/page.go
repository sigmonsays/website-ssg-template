package site

import (
	"io/fs"
	"slices"
	"strings"
)

type PageMetadata struct {
	FileEntry   *FileEntry
	ID          string
	Filename    string
	RelPath     string
	Url         string
	Title       string
	FrontMatter *FrontMatter
	Markdown    []byte
}

type FileEntry struct {
	Path string
	Info fs.FileInfo
}

type PageSlugItem struct {
	Url   string
	Label string
}

func (me *PageMetadata) GetSlug() []PageSlugItem {
	ret := make([]PageSlugItem, 0)
	tmp := strings.Split(me.Url, "/")
	for i := 1; i < len(tmp); i++ {
		s := PageSlugItem{}
		s.Label = tmp[i]
		if i < len(tmp)-1 {
			s.Url = strings.Join(tmp[0:i+1], "/")
			if s.Url == "" {
				s.Url = "/"
			}
			if s.Label == "" {
				s.Label = "/"
			}
			ret = append(ret, s)
		}
	}
	return ret
}

// sort by date, newest at top
func SortPages(pages []*PageMetadata) []PageMetadata {
	ret := make([]PageMetadata, 0)
	for _, page := range pages {
		ret = append(ret, *page)
	}

	slices.SortFunc(ret, func(a, b PageMetadata) int {
		return strings.Compare(b.FrontMatter.Date, a.FrontMatter.Date)
	})

	return ret
}
