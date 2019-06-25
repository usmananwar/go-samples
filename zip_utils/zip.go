package ziputils

import (
	"archive/zip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func exportKeystore() (interface{}, error) {
	fileToZip, err := os.Open("C:\\Users\\Usman\\zipTest\\Zipped.zip")
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(fileToZip)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(len(body))

	return body, nil
}

// Zip packages the given files into a zip file
func Zip(source, target string) error {
	files, err := ReadFilesPathInADirectory(source)
	if err != nil {
		return err
	}

	newZipFile, err := os.Create(target)
	if err != nil {
		return err
	}
	defer newZipFile.Close()

	zipWriter := zip.NewWriter(newZipFile)
	defer zipWriter.Close()

	// Add files to zip
	for _, file := range files {
		if err = AddFileToZip(zipWriter, file, source); err != nil {
			return err
		}
	}
	return nil
}

// Unzip extracts the files, from a given zipped file, to a given target location
func Unzip(zippedFile, target string) error {
	reader, err := zip.OpenReader(zippedFile)
	if err != nil {
		return err
	}

	if err := os.MkdirAll(target, 0755); err != nil {
		return err
	}

	for _, file := range reader.File {
		path := filepath.Join(target, file.Name)
		if file.FileInfo().IsDir() {
			os.MkdirAll(path, file.Mode())
			continue
		}

		fileReader, err := file.Open()
		if err != nil {
			return err
		}
		defer fileReader.Close()

		targetFile, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
		if err != nil {
			return err
		}
		defer targetFile.Close()

		if _, err := io.Copy(targetFile, fileReader); err != nil {
			return err
		}
	}

	return nil
}

// ReadFilesPathInADirectory returns a list of absolute paths of all files present in the given direcotry
func ReadFilesPathInADirectory(directoryPath string) ([]string, error) {
	var files []string
	err := filepath.Walk(directoryPath, func(path string, info os.FileInfo, err error) error {

		files = append(files, path)

		return nil
	})
	if err != nil {
		return nil, err
	}
	return files, nil
}

// AddFileToZip reads and adds a file to a given folder, which is goi,g to be zipped, after modifying the base path of the file
func AddFileToZip(zipWriter *zip.Writer, filePath string, sourceDirectory string) error {

	fileToZip, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer fileToZip.Close()

	// Get the file information
	info, err := fileToZip.Stat()
	if err != nil {
		return err
	}

	header, err := zip.FileInfoHeader(info)
	if err != nil {
		return err
	}

	header.Name = filepath.Join("", strings.TrimPrefix(filePath, sourceDirectory))

	if info.IsDir() {
		header.Name += "/"
	} else {
		header.Method = zip.Deflate
	}

	writer, err := zipWriter.CreateHeader(header)
	if err != nil {
		return err
	}

	if info.IsDir() {
		return nil
	}
	_, err = io.Copy(writer, fileToZip)
	return err
}
