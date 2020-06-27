package main

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	//gzipit("/tmp/webex", ".")
	err := tarit("/Users/kapjoshi/Documents/amp/myinterest/qrcode/cmd/d1/d2/test/best/status/", "/Users/kapjoshi/Desktop/my.tar.gz")
	fmt.Println(err)
}
func gzipit(source, target string) error {
	reader, err := os.Open(source)
	if err != nil {
		return err
	}

	filename := filepath.Base(source)
	target = filepath.Join(target, fmt.Sprintf("%s.gz", filename))
	writer, err := os.Create(target)
	if err != nil {
		return err
	}
	defer writer.Close()

	archiver := gzip.NewWriter(writer)
	join := filepath.Join(filename)
	fmt.Println(join)
	archiver.Name = join
	defer archiver.Close()

	_, err = io.Copy(archiver, reader)
	return err
}

func tarit(source, target string) error {
	tarName := filepath.Base(target)
	if !strings.Contains(tarName, ".tar.gz") {
		tarName = fmt.Sprintf("%s.tar.gz", tarName)
		target = filepath.Join(target, tarName)
	}
	abs, err := filepath.Abs(target)
	fmt.Println(abs)
	tarfile, err := os.Create(abs)
	if err != nil {
		log.Println("Error in creating tar file ", err)
		return err
	}
	defer tarfile.Close()

	//tarball := tar.NewWriter(tarfile)
	gz := gzip.NewWriter(tarfile)
	defer gz.Close()
	tarball := tar.NewWriter(gz)
	defer tarball.Close()

	var baseDir string
	baseDir = filepath.Join(`custom`, `portal`)
	return filepath.Walk(source,
		func(path string, info os.FileInfo, err error) error {
			fmt.Println("Path ", path)
			if err != nil {
				return err
			}
			header, err := tar.FileInfoHeader(info, info.Name())
			if err != nil {
				return err
			}
			if baseDir != "" {
				header.Name = filepath.Join(filepath.Clean(baseDir), strings.TrimPrefix(path, source))
			}
			if err := tarball.WriteHeader(header); err != nil {
				return err
			}

			if info.IsDir() {
				return nil
			}

			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()
			_, err = io.Copy(tarball, file)
			return err
		})
}
