package roots

import (
	"bufio"
	"errors"
	"os"
	"sync"

	"bitbucket.org/kardianos/osext"
)

func Bury() {
	bury()
}

func Regrowth(url string, wg *sync.WaitGroup) {
	regrowth(url, wg)
}

func CreateFileAndWriteData(fileName string, writeData []byte) error {
	fileHandle, err := os.Create(fileName)
	if err != nil {
		return err
	}
	writer := bufio.NewWriter(fileHandle)
	defer fileHandle.Close()
	writer.Write(writeData)
	writer.Flush()
	return nil
}

func namer() (string, error) {
	filename, err := osext.Executable()
	if err != nil {
		return "", errors.New("unable to get the current filename")
	}
	return filename, nil
}
