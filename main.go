package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func main() {
	dir, err := os.Open("files/")

	if err != nil {
		panic(err)
	}

	for {
		files, err := dir.Readdir(1)
		if err != nil {
			if err == io.EOF {
				break
			}

			fmt.Printf("ERROR: %v\n", err)
			continue
		}

		nomeArquivo := files[0].Name()
		caminhoCompleto := filepath.Join("files/", nomeArquivo)

		fmt.Println(nomeArquivo)
		os.Remove(caminhoCompleto)
	}
}
