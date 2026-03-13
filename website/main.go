package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

type Options struct {
	Verbose   bool
	DevServer bool
	InDir     string
	OutDir    string
	BindAddr  string
}

func main() {

	gen := false
	opts := &Options{
		OutDir:   "docs",
		InDir:    "content",
		BindAddr: ":8888",
	}

	flag.StringVar(&opts.InDir, "in", opts.InDir, "input directory")
	flag.StringVar(&opts.OutDir, "out", opts.OutDir, "output directory")
	flag.BoolVar(&opts.DevServer, "dev", opts.DevServer, "start dev server")
	flag.BoolVar(&gen, "gen", gen, "generate and exit")
	flag.BoolVar(&opts.Verbose, "verbose", opts.Verbose, "be verbose")
	flag.StringVar(&opts.BindAddr, "bind", opts.BindAddr, "dev http server bind address")
	flag.Parse()

	err := genSite(opts)
	if err != nil {
		log.Fatalf("GenServer %s", err)
	}

	if gen {
		log.Printf("exiting after generation.")
		os.Exit(0)
	}

	if opts.DevServer {
		log.Printf("Serving %s on %s ...", opts.OutDir, opts.BindAddr)
		fs := http.FileServer(http.Dir(opts.OutDir))
		http.Handle("/", fs)

		err := http.ListenAndServe(opts.BindAddr, nil)
		if err != nil {
			log.Fatalf("ListenAndServe: %s: %s", opts.BindAddr, err)
		}
	}
}
