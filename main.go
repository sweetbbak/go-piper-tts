package main

import (
	"bytes"
	"embed"
	"fmt"
	"os"
)

//go:embed piper.sh
var content embed.FS

func cat_file(files []string) {
	// files := []string{"file1", "file2"}
	var buf bytes.Buffer
	for _, file := range files {
		b, err := os.ReadFile(file)
		if err != nil {
			fmt.Println(err)
		}

		buf.Write(b)
	}

	err := os.WriteFile("output_file.sh", buf.Bytes(), 0755)
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	files := []string{"squashfs-start.sh", "squashfuse.tar", "piper.squashfs"}
	cat_file(files)

	exe, err := content.ReadFile("piper.sh")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(exe))
}
