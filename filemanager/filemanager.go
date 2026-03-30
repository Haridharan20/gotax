package filemanager

import (
	"example.com/gotax/utils"
)

type FileManager struct {
	InputFile  string
	OutputFile string
}

func (fm FileManager) LoadPrices() ([]float64, error) {
	lines, err := utils.ReadLines(fm.InputFile)
	if err != nil {
		return nil, err
	}

	return utils.StringToFloat(lines)
}

func (fm FileManager) SaveResult(data any) error {
	return utils.WriteJson(fm.OutputFile, data)
}

func New(inputFile, outputFile string) FileManager {
	return FileManager{
		InputFile:  inputFile,
		OutputFile: outputFile,
	}
}
