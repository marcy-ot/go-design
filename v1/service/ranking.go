package service

import (
	"fmt"
	"strconv"
)

type score struct {
	fileData  [][]string
	printData []map[string]int
}

func NewScore(f [][]string) *score {
	return &score{
		fileData: f,
	}
}

type scoreBoad struct {
	playerID   int
	count      int
	totalScore int
}

// ファイルデータをランキング形式に整形する
func (s *score) ChangeToRankingFormat() {
	playerScore := s.sumScore()
	playerAverageList := s.calcAverage(playerScore)
	s.printData = s.addRank(playerAverageList)
}

// プレイヤーのスコアを合計すると参加回数を集計する
func (s *score) sumScore() []map[string]int {
	r := make([]map[string]int, len(s.fileData)-1)
	for i, v := range s.fileData {
		if i == 0 {
			continue
		}

		playerID, err := strconv.Atoi(v[1])
		score, err := strconv.Atoi(v[2])
		if err != nil {
			continue
		}

		if isEmpty(r[playerID]) {
			r[playerID] = map[string]int{
				"player_id":   playerID,
				"count":       1,
				"total_score": score,
			}
		} else {
			r[playerID] = map[string]int{
				"player_id":   playerID,
				"count":       r[playerID]["count"] + 1,
				"total_score": r[playerID]["total_score"] + score,
			}
		}
	}
	return r
}

// プレイヤーごとの平均値を計算して配列に組み込む
func (s score) calcAverage(playerScoreList []map[string]int) []map[string]int {

	result := []map[string]int{}
	for _, l := range playerScoreList {
		m := make(map[string]int, 2)
		if !isEmpty(l) {
			m["mean_score"] = l["total_score"] / l["count"]
			m["player_id"] = l["player_id"]
			result = append(result, m)
		}
	}

	return result
}

// ランキングを付与して10位以下は切り捨てる
func (s score) addRank(score []map[string]int) []map[string]int {
	result := []map[string]int{}
	rank := 1
	count := 1
	before := 0
	for _, v := range score {
		m := make(map[string]int, 3)
		if before != v["mean_score"] {
			rank = count
			if len(result) >= 10 {
				// 10位以上は切り捨て
				break
			}
		}
		m["mean_score"] = v["mean_score"]
		m["player_id"] = v["player_id"]
		m["rank"] = rank
		count++
		before = v["mean_score"]
		result = append(result, m)
	}
	return result
}

func (s score) PrintConsole() {
	fmt.Println("rank,player_id,mean_score")
	for _, v := range s.printData {
		fmt.Printf("%v,%v,%v\n", v["rank"], v["player_id"], v["mean_score"])
	}
}

func isEmpty(m map[string]int) bool {
	return len(m) == 0
}
