package main

import "testing"

func TestLikeliestLoser(t *testing.T) {
	type testWeek struct {
		matchups [][]string
		strengths map[string]int
		expectedLoser string
	}

	testWeeks := []testWeek{
		{
			matchups: [][]string{{"a", "b"},{"c","d"}},
			strengths: map[string]int{"a": 1, "b": 2, "c": 3, "d": 5},
			expectedLoser: "c",
		},
	}

	for _, tw := range testWeeks {
		got := likeliestLoser(tw.matchups, tw.strengths)
		if got != tw.expectedLoser {
			t.Fatal("expected", tw.expectedLoser, "but got", got)
		}
	}
}
