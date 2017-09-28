package main

import (
	"archive/zip"
	"io"
	"log"
	"os"
	"path/filepath"
)

func CmdInit(path string) {
	_, err := os.Stat(path)
	if err == nil || !os.IsNotExist(err) {
		log.Fatal("Path Exist?!")
	}

	err = os.MkdirAll(path, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	root, _ := os.Getwd()
	SEPARATOR := string(filepath.Separator)
	z, err := zip.OpenReader(root + SEPARATOR + "gor-content.zip")

	for _, zf := range z.File {
		if zf.FileInfo().IsDir() {
			continue
		}
		dst := path + "/" + zf.Name
		os.MkdirAll(filepath.Dir(dst), os.ModePerm)
		f, err := os.OpenFile(dst, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}
		rc, err := zf.Open()
		if err != nil {
			log.Fatal(err)
		}
		_, err = io.Copy(f, rc)
		if err != nil {
			log.Fatal(err)
		}
		f.Sync()
		f.Close()
		rc.Close()
	}
	log.Println("Done")
}
