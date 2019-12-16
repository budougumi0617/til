package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	root := "./"

	err := filepath.Walk(root,
		func(path string, info os.FileInfo, err error) error {
			if info.IsDir() {
				return nil
			}

			rel, err := filepath.Rel(root, path)

			fmt.Printf("base: %q, full: %q\n", path, rel)

			return nil
		})

	if err != nil {
		fmt.Println(1, err)
	}
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	rel := filepath.Join(wd, "../hoge")

	fmt.Println(rel)
	fmt.Println(filepath.Clean(rel))
	if abs, err := filepath.Abs("/User"); err == nil {
		fmt.Println(abs)
	}
}
