package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {

	images, _ := filepath.Glob("*.JPG")

	if _, err := os.Stat("out"); os.IsNotExist(err) {
		err = os.MkdirAll("out", 0700)
	}

	for _, img := range images {
		err := convert(img, "out/"+img)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func convert(src string, dst string) (err error) {

	cmd, err := exec.LookPath("convert")
	if err == nil {
		out, err := exec.Command(cmd, "-resize", "25%", src, dst).CombinedOutput()
		fmt.Println(src, "->", dst)
		if err != nil {
			log.Fatal(string(out), err)
		}
		return err
	}
	return err
}
