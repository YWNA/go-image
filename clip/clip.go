package clip

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"
)

import _ "image/png"
import _ "image/jpeg"

type imageContainer struct {
	filePath       string
	name           string
	extension      string
	path           string
	fIn            io.Reader
	fOut           io.Writer
	outPutFilePath string
}

//todo 对剪裁尺寸的判断
func Resize(name string, x0 int, y0 int, x1 int, y1 int) string {
	ic := imageContainer{filePath: name}
	InputFile(&ic)
	OutFile(&ic)
	switch extension := ic.extension; extension {
	case ".jpeg":
		originImage, err := jpeg.Decode(ic.fIn)
		CheckError(err)
		imageYCbCr := originImage.(*image.YCbCr)
		err = jpeg.Encode(ic.fOut, imageYCbCr.SubImage(image.Rect(x0, y0, x1, y1)), &(jpeg.Options{Quality: 100}))
		CheckError(err)
		return ic.outPutFilePath
	case ".png":
		originImage, err := png.Decode(ic.fIn)
		CheckError(err)
		imageYCbCr := originImage.(*image.NRGBA)
		err = png.Encode(ic.fOut, imageYCbCr.SubImage(image.Rect(x0, y0, x1, y1)))
		CheckError(err)
		return ic.outPutFilePath
	default:
		log.Fatal(fmt.Sprintf("%v:%v", "不支持的文件格式", ic.extension))
		return "error"
	}
}

func InputFile(ic *imageContainer) {
	filePath, _ := filepath.Abs(ic.filePath)
	fIn, err := os.Open(filePath)
	ic.fIn = fIn
	CheckError(err)
	ic.name = filepath.Base(filePath)
	ic.extension = filepath.Ext(filePath)
	ic.path, _ = filepath.Split(filePath)
	return
}

func OutFile(ic *imageContainer) {
	temp := "temp/"
	ic.outPutFilePath = ic.path + temp + fmt.Sprintf("%v-%v", time.Now().Unix(), ic.name)
	outFile, err := os.Create(ic.outPutFilePath)
	CheckError(err)
	ic.fOut = outFile
	return
}

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
