package main

import (
	"bytes"
	"fmt"

	"github.com/sigmonsays/website/site"
	"gopkg.in/yaml.v2"
)

// parseFrontMatter extracts YAML front matter and the remaining markdown content.
func parseFrontMatter(md []byte) (*site.FrontMatter, []byte, error) {
	fm := &site.FrontMatter{}

	content := bytes.TrimSpace(md)
	if !bytes.HasPrefix(content, []byte("---")) {
		return fm, md, nil // no front matter, return as-is
	}

	parts := bytes.SplitN(content, []byte("---"), 3)
	if len(parts) < 3 {
		return fm, md, nil // malformed front matter
	}

	fmData := bytes.TrimSpace(parts[1])
	body := bytes.TrimSpace(parts[2])

	if err := yaml.Unmarshal(fmData, &fm); err != nil {
		return fm, md, fmt.Errorf("parse front matter: %w", err)
	}

	fm.SetDefaults()

	return fm, body, nil
}
