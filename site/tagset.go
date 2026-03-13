package site

type TagSet struct {
	Tag   string
	Pages []*PageMetadata
}

func (me *TagSet) GetPages() []PageMetadata {
	return SortPages(me.Pages)
}
