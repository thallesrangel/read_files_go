package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/google/uuid"
)

func main() {
	outputPath := "files/"

	numFiles := 2000000

	for i := 0; i < numFiles; i++ {
		fileName := fmt.Sprintf("ex_%s.txt", uuid.New().String())

		filePath := filepath.Join(outputPath, fileName)

		err := ioutil.WriteFile(filePath, []byte("Conteúdo do arquivo"), 0644)
		if err != nil {
			fmt.Printf("Erro ao salvar o arquivo %s: %s\n", fileName, err.Error())
			continue
		}

		fmt.Printf("Arquivo %s salvo com sucesso.\n", fileName)
	}
}
