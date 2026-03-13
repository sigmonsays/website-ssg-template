package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/sigmonsays/website/site"
)

func genSite(opts *Options) error {
	ctx := context.Background()
	inDir := opts.InDir
	outDir := opts.OutDir

	err := os.MkdirAll(outDir, 0755)
	if err != nil {
		return err
	}

	pages, err := GetPages(inDir, inDir, opts.DevServer)
	if err != nil {
		return err
	}

	// track directories and child object count so we can generate indexes
	directories := make(map[string]int, 0)
	for _, page := range pages {
		dirname := filepath.Dir(page.Filename)
		count, found := directories[dirname]
		if !found {
			count = 1
		}
		if dirname == "." || dirname == opts.InDir {
			continue
		}
		directories[dirname] = count + 1
	}

	siteObj := &site.Site{
		Pages:  pages,
		TagSet: make(map[string]*site.TagSet, 0),
	}

	siteRender := &SiteRender{
		Verbose:       opts.Verbose,
		includeDrafts: opts.DevServer,
		InDir:         opts.InDir,
		OutDir:        opts.OutDir,
		Site:          siteObj,
	}

	// build tag mapping
	for _, page := range pages {
		for _, tag := range page.FrontMatter.Tags {
			tagset, found := siteObj.TagSet[tag]
			if !found {
				tagset = &site.TagSet{
					Tag: tag,
				}
				siteObj.TagSet[tag] = tagset
			}
			tagset.Pages = append(tagset.Pages, page)
		}
	}

	// populate tags from all markdown files
	tags := make(map[string]int, 0)
	for _, page := range pages {
		for _, tag := range page.FrontMatter.Tags {
			count, _ := tags[tag]
			tags[tag] = count + 1
		}
	}
	for tag, count := range tags {
		siteObj.AllTags = append(siteObj.AllTags, tag)
		fmt.Printf("tag:%s page-count:%d\n", tag, count)
	}

	// write each page html
	for _, page := range pages {

		body, err := markdownToHTML(page.Markdown)
		if err != nil {
			log.Printf("markdown convert %s: %v", page.Filename, err)
			continue
		}
		switch page.FrontMatter.Layout {
		case "page":
			err = siteRender.WritePage(ctx, page, page.FrontMatter.Title, body)
			if err != nil {
				log.Printf("render page %s: %v", page.Filename, err)
				continue
			}
		case "directory":
			dir := filepath.Join(inDir, page.FrontMatter.Directory)
			err = siteRender.WriteDirectory(ctx, page, page.FrontMatter.Title, body, dir)
			if err != nil {
				log.Printf("render directory %s: %v", page.Filename, err)
				continue
			}
		}

	}

	// generate index
	err = siteRender.WriteIndexPage(ctx, "_YOUR_SITE_TITLE_")
	if err != nil {
		return err
	}

	// generate archive
	err = siteRender.WriteArchivePage(ctx, "archive")
	if err != nil {
		return err
	}

	// generate a page for each tag
	for tag, tagset := range siteObj.TagSet {

		err := siteRender.WriteTagPage(ctx, tag, tagset.GetPages())
		if err != nil {
			return err
		}

	}

	// generate index page for directories
	for directory, count := range directories {
		outdir := filepath.Join(opts.OutDir, directory)
		fmt.Printf("generate index for %s (%d children, outdir:%s)\n",
			directory, count, outdir)
		err = siteRender.WriteDirectoryIndexPage(ctx, directory, outdir)
		if err != nil {
			log.Printf("Error: generating dir index %s: %s", directory, err)
		}

		// generate a tag index page, ie tags.html
		err = siteRender.WriteTagsIndex(ctx, "tags")
		if err != nil {
			return err
		}
	}

	log.Printf("finished generating site")

	return nil
}
