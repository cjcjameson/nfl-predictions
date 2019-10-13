package main

import "fmt"

var teamStrengths = map[string]int{
	"atl": 1400,
	"bal": 1450,
	"chi": 1590,
	"den": 1600,
}

var schedule = [][][]string{
	//week 1
	{{"atl", "bal"}, {"chi", "den"}},
	//week 2
	{{"atl", "chi"}, {"bal", "den"}},
	//week 3
	{{"atl", "den"}, {"bal", "chi"}},
}

func main() {
	for i, weekLineup := range schedule {
		fmt.Println(
			likeliestLoser(weekLineup, teamStrengths),
			"is the most likely to lose in week",
			i+1,
		)
	}
}

func likeliestLoser(matchups [][]string, strengths map[string]int) string {
	likelyLosersByHowMuch := make(map[string]int)
	for _, matchup := range matchups {
		strength0 := strengths[matchup[0]]
		strength1 := strengths[matchup[1]]
		if strength0 > strength1 {
			likelyLosersByHowMuch[matchup[1]] = strength0 - strength1
		} else {
			likelyLosersByHowMuch[matchup[0]] = strength1 - strength0
		}
	}

	var likeliest string
	var worstHowMuchLikely int

	for loser, howMuch := range likelyLosersByHowMuch {
		if howMuch > worstHowMuchLikely {
			likeliest = loser
			worstHowMuchLikely = howMuch
		}
	}

	return likeliest
}
