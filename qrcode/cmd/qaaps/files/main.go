package main
import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"os"
	"regexp"
	"strings"
)

func main() {
	//destination := `/Users/kapjoshi/Desktop/cx-collect-temp/`
	//source := `/Users/kapjoshi/Desktop/cx-collect`
	//copy(source, destination)
	//os.RemoveAll(destination)

	r, _ := regexp.Compile("([A-Z]+)")
	allString := r.ReplaceAllString("SimpleText.go", "_$1")
	fmt.Println(allString)
	r, _ = regexp.Compile("^_([A-Z]+)")
	fmt.Println(strings.ToLower(r.ReplaceAllString(allString, "$1")))
}
func copy(source, destination string) error {
	if _, err := os.Stat(destination); os.IsNotExist(err){
			os.MkdirAll(destination, 0755)
	}
	var err error = filepath.Walk(source, func(path string, info os.FileInfo, err error) error {
		var relPath string = strings.Replace(path, source, "", 1)
		if relPath == "" {
			return nil
		}
		if info.IsDir() {
			return os.Mkdir(filepath.Join(destination, relPath), 0755)
		} else {
			var data, err1 = ioutil.ReadFile(filepath.Join(source, relPath))
			if err1 != nil {
				return err1
			}
			return ioutil.WriteFile(filepath.Join(destination, relPath), data, 0777)
		}
	})
	return err
}