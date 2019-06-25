package ziputils

import (
	"bytes"
	"errors"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/alexmullins/zip"
)

// ZipSecure packages the given files into a single zip file with a given password
func ZipSecure(source, target, password string) error {

	buf := new(bytes.Buffer)
	bufferedWriter := zip.NewWriter(buf)

	// iterating over all the available files at 'target' location
	filepath.Walk(source, func(path string, info os.FileInfo, err error) error {

		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}

		//if baseDir != "" {
		header.Name = filepath.Join("", strings.TrimPrefix(path, source))
		//}
		if info.IsDir() {
			header.Name += "/"
		} else {
			header.Method = zip.Deflate
		}

		if err != nil {
			return err
		}

		// skip directories
		if info.IsDir() {
			_, err = bufferedWriter.CreateHeader(header)
			if err != nil {
				return err
			}
			return nil
		}

		fileToZip, err := os.Open(path)
		if err != nil {
			return err
		}
		defer fileToZip.Close()

		w, err := bufferedWriter.Encrypt(header.Name, password)
		if err != nil {
			return err
		}

		_, err = io.Copy(w, fileToZip)
		return err
	})
	bufferedWriter.Flush()
	bufferedWriter.Close()

	newZipFile, err := os.Create(target)
	if err != nil {
		return err
	}
	newZipFile.Write(buf.Bytes())
	newZipFile.Close()
	return nil
}

// UnzipSecure uncompresses a zip file
func UnzipSecure(zipFilePath, outputDir, password string) error {

	zipr, err := zip.OpenReader(zipFilePath)
	if err != nil {
		return err
	}
	defer zipr.Close()

	for _, z := range zipr.File {

		path := filepath.Join(outputDir, z.Name)
		if z.FileInfo().IsDir() {
			os.MkdirAll(path, z.Mode())
			continue
		}

		z.SetPassword(password)
		rr, err := z.Open()
		if err != nil {
			return err
		}
		content, err := ioutil.ReadAll(rr)
		if err != nil {
			return err
		}

		er := writeFile(path, content)
		if er != nil {
			return er
		}

		_, err = io.Copy(os.Stdout, rr)
		if err != nil {
			return err
		}
		rr.Close()
	}
	return nil
}

func writeFile(filePath string, content []byte) error {
	if !Exists(filePath) {
		file, err := os.Create(filePath)
		if err != nil {
			return err
		}
		file.Write(content)
		file.Close()
		return nil
	}
	return errors.New("File already exists")
}

// Exists reports whether the named file or directory exists.
func Exists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

/*
func getFileWriter(path string) (*Writer, error) {

	if len(path) <= 0 {
		buf := new(bytes.Buffer)
		w := zip.NewWriter(buf)
		defer w.Close()
		return w, nil
	}

	newZipFile, err := os.Create(path)
	zipWriter := zip.NewWriter(newZipFile)
	defer newZipFile.Close()
	defer zipWriter.Close()
	return zipWriter, err
}*/
