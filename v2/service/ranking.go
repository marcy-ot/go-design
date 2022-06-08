package service

import (
	"fmt"
	"strconv"
	"strings"
)

type Ranker struct {
	logic RankingStrategy
}

func NewRanker(logic RankingStrategy) *Ranker {
	return &Ranker{logic: logic}
}

type PrintList struct {
	Result ResultScoreBoad
	Header []string
}

func (r Ranker) Execute(fileData [][]string) PrintList {
	return PrintList{
		Result: r.logic.Rank(fileData),
		Header: r.logic.GetHeader(),
	}

}

type RankingStrategy interface {
	Rank(fileData [][]string) ResultScoreBoad
	GetHeader() []string
}

type ScoreList struct {
	Rank        int
	PlayerID    int
	ResultScore int
}
type ResultScoreBoad []ScoreList

// MeanScoreTopRanking 平均値が高い順にランキングを作成
type MeanScoreTopRanking struct {
	Header []string
}

func NewMeanScoreTopRanking() *MeanScoreTopRanking {
	h := []string{
		"rank",
		"player_id",
		"mean_score",
	}
	return &MeanScoreTopRanking{
		Header: h,
	}
}

// Rank スコアの平均値の高い順にランキングを作成
func (m MeanScoreTopRanking) Rank(fileData [][]string) ResultScoreBoad {
	r := make([]map[string]int, len(fileData)-1)
	for i, v := range fileData {
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
	return m.calcAverage(r)
}

// プレイヤーごとの平均値を計算して配列に組み込む
func (m MeanScoreTopRanking) calcAverage(playerScoreList []map[string]int) ResultScoreBoad {

	result := []map[string]int{}
	for _, l := range playerScoreList {
		m := make(map[string]int, 2)
		if !isEmpty(l) {
			m["mean_score"] = l["total_score"] / l["count"]
			m["player_id"] = l["player_id"]
			result = append(result, m)
		}
	}

	return m.addRank(result)
}

// ランキングを付与して10位以下は切り捨てる
func (m MeanScoreTopRanking) addRank(score []map[string]int) ResultScoreBoad {
	result := ResultScoreBoad{}
	rank := 1
	count := 1
	before := 0
	for _, v := range score {
		var sl ScoreList
		if before != v["mean_score"] {
			rank = count
			if len(result) >= 10 {
				// 10位以上は切り捨て
				break
			}
		}
		sl.ResultScore = v["mean_score"]
		sl.PlayerID = v["player_id"]
		sl.Rank = rank
		count++
		before = v["mean_score"]
		result = append(result, sl)
	}
	return result
}

// GetHeader 出力するヘッダーを返却する
func (m MeanScoreTopRanking) GetHeader() []string {
	return m.Header
}

// PrintConsole コンソールに結果を出力する
func PrintConsole(header []string, scoreList []ScoreList) {
	fmt.Println(strings.Join(header, ","))
	for _, v := range scoreList {
		fmt.Printf("%v,%v,%v\n", v.Rank, v.PlayerID, v.ResultScore)
	}
}

func isEmpty(m map[string]int) bool {
	return len(m) == 0
}
