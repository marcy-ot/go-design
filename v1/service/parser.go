package service

import (
	"encoding/csv"
	"errors"
	"io"
	"path/filepath"
)

type Parser interface {
	Parse(f io.Reader) ([][]string, error)
}

// GetParser ファイルタイプによって読み込みの形式を変更する
func GetParser(f string) (Parser, error) {
	switch filepath.Ext(f) {
	case ".csv":
		return &Csv{}, nil
	}
	return nil, errors.New("no match file type")
}

type Csv struct{}

// Parse CSVの内容を読み込む
func (c Csv) Parse(f io.Reader) ([][]string, error) {
	r := csv.NewReader(f)
	result := [][]string{}
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		result = append(result, record)
	}
	return result, nil
}
