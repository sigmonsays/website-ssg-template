package main

import (
	"strings"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"go.abhg.dev/goldmark/anchor"
	"go.abhg.dev/goldmark/wikilink"

	htmlr "github.com/yuin/goldmark/renderer/html"
)

// converts markdown to HTML
func markdownToHTML(md []byte) (string, error) {
	mdr := goldmark.New(
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(), // read note
		),
		goldmark.WithExtensions(
			extension.GFM,
			&wikilink.Extender{},
			&anchor.Extender{
				Texter: anchor.Text("_"),
			},
		),
		goldmark.WithRendererOptions(htmlr.WithUnsafe()),
	)
	var sb strings.Builder
	if err := mdr.Convert(md, &sb); err != nil {
		return "", err
	}
	return sb.String(), nil
}
