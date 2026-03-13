package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/sigmonsays/website/site"
)

func GetPages(rootdir, dir string, includeDrafts bool) ([]*site.PageMetadata, error) {
	ret := make([]*site.PageMetadata, 0)
	entries := make([]*site.FileEntry, 0)
	walkfn := func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return nil
		}

		e := &site.FileEntry{
			Path: path,
			Info: info,
		}
		entries = append(entries, e)
		return nil
	}
	err := filepath.Walk(dir, walkfn)
	if err != nil {
		return nil, err
	}

	for _, ent := range entries {
		e := ent.Info
		if e.IsDir() || !strings.HasSuffix(e.Name(), ".md") {
			continue
		}

		md, err := os.ReadFile(ent.Path)
		if err != nil {
			log.Printf("read %s: %v", ent.Path, err)
			continue
		}

		// Parse front matter
		fm, bodyMD, err := parseFrontMatter(md)
		if err != nil {
			log.Printf("front matter %s: %v", ent.Path, err)
			continue
		}

		if includeDrafts == false && fm.Draft == true {
			log.Printf("skipping draft page %s", ent.Path)
			continue

		}

		// skip pages without a title
		if fm.Title == "" {
			log.Printf("skipping page without title %s", ent.Path)
			continue
		}

		relPath, err := filepath.Rel(rootdir, ent.Path)
		if err != nil {
			fmt.Printf("ERROR: RelPath: %s", err)
		}

		pageId := strings.TrimSuffix(relPath, ".md")
		pageUrl := "/" + pageId + ".html"

		page := &site.PageMetadata{
			FileEntry:   ent,
			Filename:    ent.Path,
			RelPath:     relPath,
			ID:          pageId,
			Url:         pageUrl,
			Title:       fm.Title,
			Markdown:    bodyMD,
			FrontMatter: fm,
		}
		// fmt.Printf("Page %s Url:%q RelPath:%q\n",
		// 	page.Filename, page.Url, page.RelPath)
		ret = append(ret, page)
	}

	return ret, nil
}
