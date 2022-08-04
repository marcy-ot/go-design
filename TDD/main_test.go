package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestZeroCsvParse(t *testing.T) {
	tests := []struct {
		name string
		want [][]string
		path string
	}{
		{
			name: "０件データのCSVファイル読み込み",
			path: "./testdata/game_score_no_data.csv",
		}, {
			name: "正常なCSVファイル読み込み",
			want: [][]string{
				{
					"create_timestamp",
					"player_id",
					"score",
				},
				{
					"2021/10/1",
					"1",
					"0",
				}, {
					"2021/10/1",
					"2",
					"123",
				},
			},
			path: "./testdata/game_score_nomal.csv",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewCsvUtil(tt.path)
			got := c.Parse()
			if !reflect.DeepEqual(tt.want, got) {
				fmt.Printf("len %v, cap %v, val %v\n", len(got), cap(got), got)
				fmt.Printf("len %v, cap %v, val %v\n", len(tt.want), cap(tt.want), tt.want)
				t.Errorf("data = %v, want %v", got, tt.want)
			}
		})
	}
}

var emptyCsvData = [][]string{}
var normalCsvDagta = [][]string{
	{
		"create_timestamp",
		"player_id",
		"score",
	},
	{
		"2021/10/1",
		"1",
		"0",
	}, {
		"2021/10/1",
		"2",
		"123",
	},
}

func TestNewScoreBoad(t *testing.T) {
	scoreBoad := NewScoreBoad(normalCsvDagta)
	want := &ScoreBoad{
		scores: []Score{
			{
				playerID: 1,
				score:    0,
				rank:     0,
			},
			{
				playerID: 2,
				score:    123,
				rank:     0,
			},
		},
	}
	t.Run("正常系のデータをScoreにマッピングする", func(t *testing.T) {
		if !reflect.DeepEqual(scoreBoad, want) {
			t.Errorf("scoreBoad = %v, want %v", scoreBoad, want)
		}
	})

	scoreBoad = NewScoreBoad(emptyCsvData)
	want = &ScoreBoad{}
	t.Run("空リストのマッピング", func(t *testing.T) {
		if !reflect.DeepEqual(scoreBoad, want) {
			t.Errorf("scoreBoad = %v, want %v", scoreBoad, want)
		}
	})
}

var emptyScoreBoad = &ScoreBoad{}
var normalScoreBoad = &ScoreBoad{}

func TestSumScore(t *testing.T) {
	want := NewScoreBoad([][]string{
		{
			"2021/10/1",
			"1",
			"0",
		}, {
			"2021/10/1",
			"2",
			"300",
		},
	})
	scoreBoad := NewScoreBoad([][]string{
		{
			"create_timestamp",
			"player_id",
			"score",
		},
		{
			"2021/10/1",
			"1",
			"0",
		}, {
			"2021/10/1",
			"2",
			"123",
		}, {
			"2021/10/1",
			"2",
			"177",
		},
	})
	result := scoreBoad.sumScoreByPlayer()
	t.Run("ユーザースコアの加算", func(t *testing.T) {
		if !reflect.DeepEqual(want, result) {
			t.Errorf("result = %v, want %v", result, want)
		}
	})
}

func TestAddRank(t *testing.T) {
	want := &ScoreBoad{
		scores: []Score{
			Score{
				playerID: 5,
				score:    200,
				rank:     1,
			},
			Score{
				playerID: 4,
				score:    100,
				rank:     2,
			},
			Score{
				playerID: 3,
				score:    90,
				rank:     3,
			},
			Score{
				playerID: 2,
				score:    70,
				rank:     4,
			},
			Score{
				playerID: 1,
				score:    50,
				rank:     5,
			},
			Score{
				playerID: 6,
				score:    40,
				rank:     6,
			},
			Score{
				playerID: 7,
				score:    30,
				rank:     7,
			},
		},
	}
	scoreBoad := NewScoreBoad([][]string{
		{
			"create_timestamp",
			"player_id",
			"score",
		},
		{
			"2021/10/1",
			"1",
			"50",
		}, {
			"2021/10/1",
			"2",
			"70",
		}, {
			"2021/10/1",
			"3",
			"90",
		}, {
			"2021/10/1",
			"4",
			"100",
		}, {
			"2021/10/1",
			"5",
			"200",
		}, {
			"2021/10/1",
			"6",
			"40",
		}, {
			"2021/10/1",
			"7",
			"30",
		},
	})
	scoreBoad.addRank()
	t.Run("ランキング付与（スコアの重複がない場合）", func(t *testing.T) {
		if !reflect.DeepEqual(scoreBoad, want) {
			t.Errorf("got = %v, want %v", scoreBoad, want)
		}
	})

	want = &ScoreBoad{
		scores: []Score{
			Score{
				playerID: 4,
				score:    200,
				rank:     1,
			},
			Score{
				playerID: 5,
				score:    200,
				rank:     1,
			},
			Score{
				playerID: 1,
				score:    90,
				rank:     3,
			},
			Score{
				playerID: 2,
				score:    90,
				rank:     3,
			},
			Score{
				playerID: 3,
				score:    90,
				rank:     3,
			},
			Score{
				playerID: 6,
				score:    40,
				rank:     6,
			},
			Score{
				playerID: 7,
				score:    30,
				rank:     7,
			},
		},
	}
	scoreBoad = NewScoreBoad([][]string{
		{
			"create_timestamp",
			"player_id",
			"score",
		},
		{
			"2021/10/1",
			"1",
			"90",
		}, {
			"2021/10/1",
			"2",
			"90",
		}, {
			"2021/10/1",
			"3",
			"90",
		}, {
			"2021/10/1",
			"4",
			"200",
		}, {
			"2021/10/1",
			"5",
			"200",
		}, {
			"2021/10/1",
			"6",
			"40",
		}, {
			"2021/10/1",
			"7",
			"30",
		},
	})
	scoreBoad.addRank()
	t.Run("ランキング付与（スコアの重複があり場合）", func(t *testing.T) {
		if !reflect.DeepEqual(scoreBoad, want) {
			t.Errorf("\ngot = %v\nwant %v", scoreBoad, want)
		}
	})
}

func TestRecordLimit(t *testing.T) {
	want := &ScoreBoad{
		scores: []Score{
			Score{
				playerID: 5,
				score:    200,
				rank:     1,
			},
			Score{
				playerID: 4,
				score:    100,
				rank:     2,
			},
			Score{
				playerID: 3,
				score:    90,
				rank:     3,
			},
			Score{
				playerID: 2,
				score:    70,
				rank:     4,
			},
			Score{
				playerID: 1,
				score:    50,
				rank:     5,
			},
		},
	}
	scoreBoad := &ScoreBoad{
		scores: []Score{
			Score{
				playerID: 5,
				score:    200,
				rank:     1,
			},
			Score{
				playerID: 4,
				score:    100,
				rank:     2,
			},
			Score{
				playerID: 3,
				score:    90,
				rank:     3,
			},
			Score{
				playerID: 2,
				score:    70,
				rank:     4,
			},
			Score{
				playerID: 1,
				score:    50,
				rank:     5,
			},
			Score{
				playerID: 6,
				score:    40,
				rank:     6,
			},
			Score{
				playerID: 7,
				score:    30,
				rank:     7,
			},
		},
	}
	result := scoreBoad.cut(5)
	t.Run("上位5名のレコードに絞る(ランク重複なし）", func(t *testing.T) {
		if !reflect.DeepEqual(result, want) {
			t.Errorf("scoreBoad = %v, want %v", result, want)
		}
	})

	want = &ScoreBoad{
		scores: []Score{
			Score{
				playerID: 4,
				score:    200,
				rank:     1,
			},
			Score{
				playerID: 5,
				score:    200,
				rank:     1,
			},
			Score{
				playerID: 1,
				score:    90,
				rank:     3,
			},
			Score{
				playerID: 2,
				score:    90,
				rank:     3,
			},
			Score{
				playerID: 3,
				score:    90,
				rank:     3,
			},
			Score{
				playerID: 6,
				score:    90,
				rank:     3,
			},
		},
	}
	scoreBoad = &ScoreBoad{
		scores: []Score{
			Score{
				playerID: 4,
				score:    200,
				rank:     1,
			},
			Score{
				playerID: 5,
				score:    200,
				rank:     1,
			},
			Score{
				playerID: 1,
				score:    90,
				rank:     3,
			},
			Score{
				playerID: 2,
				score:    90,
				rank:     3,
			},
			Score{
				playerID: 3,
				score:    90,
				rank:     3,
			},
			Score{
				playerID: 6,
				score:    90,
				rank:     3,
			},
			Score{
				playerID: 7,
				score:    30,
				rank:     7,
			},
		},
	}
	result = scoreBoad.cut(5)
	t.Run("上位5名のレコードに絞る(ランク重複あり）", func(t *testing.T) {
		if !reflect.DeepEqual(result, want) {
			t.Errorf("scoreBoad = %v, want %v", result, want)
		}
	})
}
