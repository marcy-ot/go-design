package service

import (
	"reflect"
	"testing"
)

func Test_TotallingScoreAndCountByUser_Nomal_OK(t *testing.T) {
	test := []struct {
		name string
		args [][]string
		want []map[string]int
	}{
		{
			name: "合計値集計　正常系テスト",
			args: [][]string{
				{"created_timestamp", "player_id", "score"},
				{"2021/10/1", "1", "120"},
				{"2021/10/2", "2", "220"},
				{"2021/10/3", "2", "220"},
				{"2021/10/4", "3", "120"},
				{"2021/10/4", "3", "10"},
				{"2021/10/4", "3", "-10"},
			},
			want: []map[string]int{
				{
					"player_id":   1,
					"count":       1,
					"total_score": 120,
				}, {
					"player_id":   2,
					"count":       2,
					"total_score": 440,
				}, {
					"player_id":   3,
					"count":       3,
					"total_score": 120,
				},
			},
		},
	}

	m := NewMeanScoreTopRanking()
	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			got := m.totallingScoreAndCountByUser(tt.args)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Rank() = %v,\n got:\n %v", got, tt.want)
			}
		})
	}
}

func Test_CalcAveratge_Nomal_OK(t *testing.T) {
	test := []struct {
		name string
		args []map[string]int
		want []map[string]int
	}{
		{
			name: "平均値算出ロジック　正常系テスト",
			args: []map[string]int{
				{
					"player_id":   1,
					"count":       1,
					"total_score": 120,
				}, {
					"player_id":   2,
					"count":       2,
					"total_score": 440,
				}, {
					"player_id":   3,
					"count":       3,
					"total_score": 120,
				},
			},
			want: []map[string]int{
				{
					"player_id":  1,
					"mean_score": 120,
				}, {
					"player_id":  2,
					"mean_score": 220,
				}, {
					"player_id":  3,
					"mean_score": 40,
				},
			},
		},
	}

	m := MeanScoreTopRanking{}
	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			got := m.calcAverage(tt.args)
			if !reflect.DeepEqual(tt.want, got) {
				t.Errorf("calcAverage = %v,\n want: \n %v", got, tt.want)
			}
		})
	}
}

func Test_AddRank_Normal_Ok(t *testing.T) {
	test := []struct {
		name string
		args []map[string]int
		want ResultScoreBoad
	}{
		{
			name: "ランキング付与ロジック　正常系テスト",
			args: []map[string]int{
				{
					"player_id":  1,
					"mean_score": 120,
				}, {
					"player_id":  2,
					"mean_score": 220,
				}, {
					"player_id":  3,
					"mean_score": 40,
				},
			},
			want: ResultScoreBoad{
				{
					Rank:        1,
					PlayerID:    2,
					ResultScore: 220,
				}, {
					Rank:        2,
					PlayerID:    1,
					ResultScore: 120,
				}, {
					Rank:        3,
					PlayerID:    3,
					ResultScore: 40,
				},
			},
		},
	}

	m := MeanScoreTopRanking{}
	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			got := m.addRank(tt.args)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addRank = %v,\n want: \n%v", got, tt.want)
			}
		})
	}

}
