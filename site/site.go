package site

import (
	"cmp"
	"slices"
)

type Site struct {
	AllTags []string
	Pages   []*PageMetadata
	TagSet  map[string]*TagSet
}

func (me *Site) GetPages(max int) []PageMetadata {
	ret := SortPages(me.Pages)
	if max == 0 || len(ret) < max {
		return ret
	}

	return ret[:max]
}

func (me *Site) GetTags() []*TagSet {
	ret := make([]*TagSet, 0)
	for _, tagset := range me.TagSet {
		ret = append(ret, tagset)
	}
	slices.SortFunc(ret, func(a, b *TagSet) int {
		return cmp.Compare(a.Tag, b.Tag)

	})
	return ret
}
