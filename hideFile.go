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

// NewHider returns a new instance of a Hider
func NewHider() *Hider {
	return &Hider{}
}

// Converts the file to the given type and writes it to the given location
func (h *Hider) Convert(file *os.File, toType magicNumber, out string) error {
	fileContent, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	newFileContent := append(toType.Number, fileContent...)
	destinationFileName := h.generateFileName(file, toType, out)

	return h.write(destinationFileName, newFileContent)
}

// Deconverts the file from the type and writes it to the given location
func (h *Hider) Deconvert(file *os.File, fromType magicNumber, out string) error {
	fileContent, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	newFileContent := fileContent[len(fromType.Extension):]
	destinationFileName := h.generateFileName(file, fromType, out)

	return h.write(destinationFileName, newFileContent)
}

// write writes the content of the new file to the destination
func (h *Hider) write(destinationFileName string, content []byte) error {
	newFile, err := os.OpenFile(destinationFileName, os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		return err
	}
	defer newFile.Close()

	_, err = newFile.WriteAt(content, 0)

	return err
}

// generates the name of the new file according to the wanted type
func (h *Hider) generateFileName(file *os.File, toType magicNumber, out string) string {
	return out + "/" + strings.TrimSuffix(filepath.Base(file.Name()), filepath.Ext(file.Name())) + "." + toType.Extension
}

// GetType returns the magicNumber to the given Typename
func (h *Hider) GetType(name string) (magicNumber, error) {
	for _, number := range magicNumberList {
		if number.Name == name {
			return number, nil
		}
	}

	return magicNumber{}, errors.New("No element found")
}

// GetTypeList returns the list of currently supported types
func (h *Hider) GetTypelist() []magicNumber {
	return magicNumberList
}
