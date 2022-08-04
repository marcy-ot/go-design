package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"reflect"
	"sort"
	"strconv"
)

func main() {
	csv := NewCsvUtil("./testdata/game_score.csv")
	data := csv.Parse()

	result := NewScoreBoad(data).sumScoreByPlayer().addRank().cut(10)
	result.printConsole()
}

type CsvUtil struct {
	Path    string
	RawData [][]string
}

func NewCsvUtil(path string) *CsvUtil {
	return &CsvUtil{
		Path: path,
	}
}

func (c CsvUtil) Parse() [][]string {
	file, err := os.Open(c.Path)
	if err != nil {
		log.Fatalf("faild file when open %v", err)
	}
	defer file.Close()

	r := csv.NewReader(file)
	rawData, err := r.ReadAll()
	if err != nil {
		log.Fatalf("faild file when reading %v", err)
	}
	// TODO: 正常値がチェック

	return rawData
}

type Score struct {
	playerID int
	score    int
	rank     int
}

type ScoreBoad struct {
	scores []Score
}

var fileHeader = []string{
	"create_timestamp",
	"player_id",
	"score",
}

var consoleHeader = []string{
	"rank",
	"score",
	"player_id",
}

func NewScoreBoad(scoreList [][]string) *ScoreBoad {
	var scoresList []Score
	for _, v := range scoreList {
		if reflect.DeepEqual(fileHeader, v) {
			continue
		}
		playerID, _ := strconv.Atoi(v[1])
		score, _ := strconv.Atoi(v[2])
		scoreList := Score{
			playerID: playerID,
			score:    score,
		}
		scoresList = append(scoresList, scoreList)
	}

	return &ScoreBoad{
		scores: scoresList,
	}
}

func (s ScoreBoad) sumScoreByPlayer() *ScoreBoad {
	aggScoreList := NewScoreBoad([][]string{})
	for _, data := range s.scores {
		if aggScoreList.hasPlayer(data.playerID) {
			aggScoreList.addScore(data.playerID, data.score)
		} else {
			aggScoreList.appendScore(data.playerID, data.score)
		}
	}
	return aggScoreList
}

func (s ScoreBoad) hasPlayer(id int) bool {
	for _, v := range s.scores {
		if v.playerID == id {
			return true
		}
	}
	return false
}

func (s *ScoreBoad) addScore(id, score int) {
	for i, v := range s.scores {
		if v.playerID == id {
			s.scores[i].score += score
		}
	}
}

func (s *ScoreBoad) appendScore(id, score int) {
	scoreList := Score{
		playerID: id,
		score:    score,
	}
	s.scores = append(s.scores, scoreList)
}

func (s *ScoreBoad) addRank() *ScoreBoad {
	// スコアが高い順にソート
	sort.SliceStable(s.scores, func(i, j int) bool {
		return s.scores[i].score > s.scores[j].score
	})

	rank := 1
	for i, v := range s.scores {
		if i == 0 {
			s.scores[i].rank = rank
			continue
		}
		if s.scores[i-1].score > v.score {
			rank = i + 1
		}
		s.scores[i].rank = rank
	}
	return s
}

func (s *ScoreBoad) cut(limit int) *ScoreBoad {
	result := NewScoreBoad([][]string{})
	for i, v := range s.scores {
		if i != 0 {
			// 直前のランクと同じランクかどうか
			samePreviousRank := v.rank == s.scores[i-1].rank
			if (i >= limit) && !samePreviousRank {
				break
			}
		}
		result.scores = append(result.scores, v)
	}
	return result
}

func (s ScoreBoad) printConsole() {
	fmt.Println("rank,player_id,mean_score")
	for _, v := range s.scores {
		fmt.Printf("%v,%v,%v\n", v.rank, v.playerID, v.score)
	}
}
