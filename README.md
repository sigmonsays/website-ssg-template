# features
- index page (listing recent posts)
  - need dates
- directory index page generation
- archive page (listing all posts)
- no tags page (listing tags and pages with tag)
- supports frontmatter data for title, summary and date
- frontmatter draft flag to skip page in production mode

# setup

```bash
  go install github.com/a-h/templ/cmd/templ@latest

  # render tmpl files into go code
  go generate ./site/

  # Compile website
  go install ./website

  # Generate website (results written to ./docs, -dev flag includes drafts)
  website -gen -dev

  # Start static HTTP server
  website -dev

```

# todo
- table of contents?
  see https://pkg.go.dev/go.abhg.dev/goldmark/toc#section-readme
- whole bunch of extensions here https://go.abhg.dev/
  - hash tags (instead of tags)
    see https://pkg.go.dev/go.abhg.dev/goldmark/hashtag
- wikilink support using https://github.com/abhinav/goldmark-wikilink
