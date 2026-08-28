// Command create_mod packs a module source zip from a checkout using
// golang.org/x/mod/zip, which applies the module proxy's own inclusion rules
// (dropping .git/, vendor/ and nested modules) so the resulting file set
// matches what proxy.golang.org serves for the same version.
//
// It lives under .github/ so the Go tool never sees it (the go command ignores
// directories whose names begin with "."), and so the zip step's rsync copy,
// which excludes .github, never packages it.
package main

import (
	"log"
	"os"

	"golang.org/x/mod/module"
	"golang.org/x/mod/zip"
)

func main() {
	if len(os.Args) != 5 {
		log.Fatal("usage: create_mod <module-path> <version> <source-dir> <output-zip>")
	}
	m := module.Version{Path: os.Args[1], Version: os.Args[2]}
	f, err := os.Create(os.Args[4])
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	if err := zip.CreateFromDir(f, m, os.Args[3]); err != nil {
		log.Fatal(err)
	}
	log.Printf("created module zip: %s", os.Args[4])
}
