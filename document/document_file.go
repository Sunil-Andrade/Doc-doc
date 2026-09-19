package document

import (
	"os"
	"path/filepath"
)

func CreateDocumentFile(code string) (string, error) {

	folderPath := filepath.Join("documents", code)

	err := os.MkdirAll(folderPath, 0755)
	if err != nil {
		return "", err
	}

	filePath := filepath.Join(folderPath, "document.txt")

	file, err := os.Create(filePath)
	if err != nil {
		return "", err
	}

	defer file.Close()

	return filePath, nil
}
