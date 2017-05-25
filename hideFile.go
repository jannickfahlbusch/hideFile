package hideFile

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type Hider struct {
	encrypt            bool
	encryptionPassword string
}

func NewHider() *Hider {
	return &Hider{}
}

func (h *Hider) Convert(file *os.File, out string, toType magicNumber) error {
	fileContent, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	newFileContent := append(toType.Number, fileContent...)
	destinationFileName := h.generateFileName(file, toType, out)

	return h.write(destinationFileName, newFileContent)
}

func (h *Hider) Deconvert(file *os.File, out string, fromType magicNumber) error {
	fileContent, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	newFileContent := fileContent[len(fromType.Extension)+1:]
	destinationFileName := h.generateFileName(file, fromType, out)

	return h.write(destinationFileName, newFileContent)
}

func (h *Hider) write(destinationFileName string, content []byte) error {
	newFile, err := os.OpenFile(destinationFileName, os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		return err
	}
	defer newFile.Close()

	_, err = newFile.WriteAt(content, 0)

	return err
}

func (h *Hider) generateFileName(file *os.File, toType magicNumber, out string) string {
	return out + "/" + strings.TrimSuffix(filepath.Base(file.Name()), filepath.Ext(file.Name())) + "." + toType.Extension
}

func (h *Hider) GetType(name string) (magicNumber, error) {
	for _, number := range magicNumberList {
		if number.Name == name {
			return number, nil
		}
	}

	return magicNumber{}, errors.New("No element found")
}

func (h *Hider) GetTypelist() []magicNumber {
	return magicNumberList
}
