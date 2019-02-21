package clip

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"
)

type imageFile struct {
	filePath  string
	name      string
	extension string
	path      string
	fIn       io.Reader
	fOut      io.Writer
}

func Resize(name string) io.Writer {
	newImageFile := imageFile{filePath: name}
	InputFile(&newImageFile)
	return newImageFile.fOut
}

func InputFile(newImageFile *imageFile) {
	filePath, _ := filepath.Abs(newImageFile.filePath)
	fIn, err := os.Open(filePath)
	newImageFile.fIn = fIn
	CheckError(err)
	newImageFile.name = filepath.Base(filePath)
	newImageFile.extension = filepath.Ext(filePath)
	newImageFile.path, _ = filepath.Split(filePath)
	newImageFile.fOut = OutFile(newImageFile.path + fmt.Sprintf("%v-%v", time.Now().Unix(), newImageFile.name))
	fmt.Println(newImageFile)
	defer fIn.Close()
	return
}

func OutFile(filePath string) io.Writer {
	outFile, err := os.Create(filePath)
	CheckError(err)
	defer outFile.Close()
	return outFile
}

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
