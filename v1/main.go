package main

import (
	"develop/go-design/v1/service"
	"log"
	"os"
)

func main() {
	f := os.Args[1]
	file, err := os.Open(f)
	if err != nil {
		log.Fatalf("cannot open file %v: %v", file, err)
	}
	defer file.Close()

	i, err := service.GetParser(f)
	if err != nil {
		log.Fatalln(err)
	}

	data, err := i.Parse(file)
	if err != nil {
		log.Fatalf("missing file parse: %v", err)
	}

	score := service.NewScore(data)
	score.ChangeToRankingFormat()
	score.PrintConsole()
}
