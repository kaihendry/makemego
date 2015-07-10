package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
)

func in(slice []string, str string) bool {
	for _, s := range slice {
		if s == str {
			return true
		}
	}
	return false
}

func main() {

	var acceptedImageExt = []string{".jpg", ".jpeg"}
	var images = []string{}

	filepath.Walk(".", func(filePath string, info os.FileInfo, err error) error {
		if err == nil && in(acceptedImageExt, strings.ToLower(path.Ext(filePath))) {
			images = append(images, filePath)
		}
		return nil
	})

	for key, value := range images {
		fmt.Println(key, value)
	}
}

func convert(src string, dst string) (err error) {

	dir, _ := filepath.Split(dst)
	err = os.MkdirAll(dir, 0700)
	if err != nil {
		return
	}

	cmd, err := exec.LookPath("convert")
	if err == nil {
		out, err := exec.Command(cmd, "-resize", "25%", src, dst).CombinedOutput()
		if err != nil {
			fmt.Printf("The output is %s\n", out)
			log.Fatal(err)
		}
		return err
	}
	return err
}
