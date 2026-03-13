package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/a-h/templ"
	"github.com/sigmonsays/website/site"
)

type SiteRender struct {
	Verbose       bool
	includeDrafts bool
	InDir         string
	OutDir        string
	Site          *site.Site
}

func (me *SiteRender) WriteComponent(ctx context.Context, c templ.Component, filename string) error {
	out, err := templ.ToGoHTML(ctx, c)
	if err != nil {
		return err
	}
	// create leading directory
	dir := filepath.Dir(filename)
	os.MkdirAll(dir, 0755)

	err = os.WriteFile(filename, []byte(out), 0644)
	if err != nil {
		return err
	}
	if me.Verbose {
		fmt.Printf("wrote %s\n", filename)
	}
	return nil
}

func (me *SiteRender) WriteIndexPage(ctx context.Context, title string) error {
	outfile := filepath.Join(me.OutDir, "index.html")
	comp := site.Index(*me.Site, title)
	return me.WriteComponent(ctx, comp, outfile)
}

func (me *SiteRender) WriteArchivePage(ctx context.Context, title string) error {
	outfile := filepath.Join(me.OutDir, "archive.html")
	comp := site.Archive(*me.Site, title)
	return me.WriteComponent(ctx, comp, outfile)
}

func (me *SiteRender) WriteTagPage(ctx context.Context, tag string, pages []site.PageMetadata) error {
	outfile := filepath.Join(me.OutDir, "tag", tag+".html")
	comp := site.PagesWithTag(*me.Site, pages, tag)
	return me.WriteComponent(ctx, comp, outfile)
}

func (me *SiteRender) WritePage(ctx context.Context, page *site.PageMetadata, title string, body string) error {
	outName := strings.TrimSuffix(page.RelPath, ".md") + ".html"
	outfile := filepath.Join(me.OutDir, outName)
	comp := site.Page(*me.Site, *page, title, body)
	return me.WriteComponent(ctx, comp, outfile)
}
func (me *SiteRender) WriteDirectory(ctx context.Context, page *site.PageMetadata, title string, body string, directory string) error {
	outName := strings.TrimSuffix(page.RelPath, ".md") + ".html"
	outfile := filepath.Join(me.OutDir, outName)
	files, err := site.ListFiles(directory)
	if err != nil {
		return err
	}
	fmt.Printf("directory %s (%d files)\n", directory, len(files))
	for _, file := range files {
		fmt.Printf(" file %s\n", file.Path)
	}
	comp := site.Directory(*me.Site, *page, title, body, directory, files)
	return me.WriteComponent(ctx, comp, outfile)
}

func (me *SiteRender) WriteDirectoryIndexPage(ctx context.Context, inDir, outDir string) error {
	relDir, err := filepath.Rel(me.InDir, inDir)
	if err != nil {
		return err
	}
	outfile := filepath.Join(me.OutDir, relDir, "index.html")
	//log.Printf("Write directory index %s", outfile)
	title := fmt.Sprintf("%s", relDir)
	_pages, err := GetPages(me.InDir, inDir, me.includeDrafts)
	if err != nil {
		return err
	}
	pages := site.SortPages(_pages)

	comp := site.DirectoryIndex(*me.Site, title, pages)
	return me.WriteComponent(ctx, comp, outfile)
}

func (me *SiteRender) WriteTagsIndex(ctx context.Context, title string) error {
	outfile := filepath.Join(me.OutDir, "tags.html")
	comp := site.TagIndex(*me.Site, title)
	return me.WriteComponent(ctx, comp, outfile)
}
