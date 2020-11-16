package main

import (
	"archive/tar"
	"bufio"
	"compress/gzip"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// Policy bundle
func main() {
	BundlePolicy("/Users/kapjoshi/Documents/amp/myinterest/qrcode/cmd/d1/d2/test/best/status/", "/Users/kapjoshi/Desktop/my.tar.gz")
}

func BundlePolicy(pFileDir, opDir string) {
	var pkgName []string
	dirs, err := ioutil.ReadDir(pFileDir)
	if err != nil {
		log.Println("Error while reading directory containing policy")
		return
	}
	for _, f := range dirs {
		fe := filepath.Ext(f.Name())
		if strings.Contains(fe, ".rego") {
			f, err := os.Open(filepath.Join(pFileDir, f.Name()))
			defer f.Close()
			if err != nil {
				log.Println("Error while opening policy file ", err)
				return
			}
			scanner := bufio.NewScanner(f)
			for scanner.Scan() {
				line := scanner.Text()
				if strings.Contains(line, `package`) {
					fields := strings.Fields(line)
					pkgName = strings.Split(fields[1], ".")
					break
				}
			}

			if err := scanner.Err(); err != nil {
				log.Fatal(err)
			}
			break
		}
	}

	createTar(pFileDir, opDir, pkgName)
}

func createTar(source string, target string, packageName []string) error {
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

	tarball := tar.NewWriter(tarfile)
	gz := gzip.NewWriter(tarfile)
	defer gz.Close()
	tarball = tar.NewWriter(gz)
	defer tarball.Close()

	var baseDir string
	baseDir = filepath.Join(packageName...)
	return filepath.Walk(source,
		func(path string, info os.FileInfo, err error) error {
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
