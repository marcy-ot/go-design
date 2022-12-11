package pagemaker

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var Database = database{}

type database struct{}

type MailData struct {
	EMail    string
	UserName string
}

type MailDatas []MailData

func (m MailDatas) GetPropaty(mailAddr string) string {
	for _, v := range m {
		if v.EMail == mailAddr {
			return v.UserName
		}
	}
	return ""
}

func (d database) GetProperties(dbName string) MailDatas {
	lines := fromFile(dbName + ".txt")
	var result []MailData
	for _, line := range lines {
		prop := strings.Split(line, "=")
		r := MailData{
			EMail:    prop[0],
			UserName: prop[1],
		}
		result = append(result, r)
	}
	return result
}

func fromFile(fileName string) []string {
	f, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("could not open file %s", fileName)
		os.Exit(1)
	}
	defer f.Close()

	lines := make([]string, 0, 100)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if serr := scanner.Err(); serr != nil {
		fmt.Printf("file scan error %v", serr)
	}

	return lines
}
