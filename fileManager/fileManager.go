package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
)

type FileManager struct {
	InputFilepath  string
	OutputFilepath string
}

func (fm FileManager) ReadLines() ([]string, error) {
	file, err := os.Open(fm.InputFilepath)
	if err != nil {
		return nil, errors.New("could not open file !")
	}

	Scanner := bufio.NewScanner(file)
	var lines []string
	for Scanner.Scan() {
		lines = append(lines, Scanner.Text())
	}
	err = Scanner.Err()
	if err != nil {
		file.Close()
		return nil, errors.New("Reading the file got failed!")
	}
	file.Close()
	return lines, nil
}

func (fm FileManager) WriteResult(data interface{}) error {
	file, err := os.Create(fm.OutputFilepath)
	if err != nil {
		file.Close()
		return errors.New("failed to create file..")
	}
	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)
	if err != nil {
		file.Close()
		return errors.New("failed to encode json file..")
	}
	file.Close()
	return nil
}

func New(inputfilepath, outputfilepath string) FileManager {
	return FileManager{
		InputFilepath:  inputfilepath,
		OutputFilepath: outputfilepath,
	}
}
