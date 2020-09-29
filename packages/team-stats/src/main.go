package main

import (
	"fmt"

	"github.com/ethantpainter/Baseball-Backend/packages/team-stats/src/internal"
)

type teamStatsDynamoModel struct {
	Wins        int
	Losses      int
	GamesPlayed int

	Hitting struct {
		Runs                      int
		Hits                      int
		Doubles                   int
		Triples                   int
		RunsBattedIn              int
		PitchesPerPlateAppearance int
		Walks                     int
		Strikeouts                int
		Average                   float32
		Slugging                  float32
		OnBasePercentage          float32
		LeadingBattingAverage     struct {
			playerID string
			value    float32
		}
		LeadingHomeRuns struct {
			playerID string
			value    float32
		}
		LeadingRunsBattedIn struct {
			playerID string
			value    float32
		}
		LeadingOnBasePercentage struct {
			playerID string
			value    float32
		}
		LeadingHits struct {
			playerID string
			value    float32
		}
	}

	Pitching struct {
		InningsPitched                 float32
		Runs                           int
		Hits                           int
		Doubles                        int
		Triples                        int
		HomeRuns                       int
		Walks                          int
		Strikeouts                     int
		StrikeoutsPerNine              int
		PitchesPerStart                float32
		WalksPlusHitsPerInningsPitched float32
		EarnedRunAverage               float32
		Saves                          float32
		StolenBases                    int

		LeadingWins struct {
			playerID string
			value    int
		}
		LeadingEarnedRunAverage struct {
			playerID string
			value    float32
		}
		LeadingStrikeouts struct {
			playerID string
			value    int
		}
		LeadingSaves struct {
			playerID string
			value    int
		}
		LeadingHolds struct {
			playerID string
			value    int
		}
	}

	Fielding struct {
		Errors             int
		FieldingPercentage float32
	}
}

// import ("fmt" "context" "github.com/aws/aws-lambda-go/lambda")
// type MyEvent string {
//   Name string `json:"name"`
// }
// type HandleRquest(ctx context.Context, name MyEvent) (string error) {
//   return fmt.Sprintf("Hello %s!", name.Name), nil
// }
// func main() {
//   lambda.Start(HandleRequest)
// }

func main() {
	teamAbrs := [30]string{
		"nyy", "tb", "tor", "bal", "bos",
		"min", "cle", "chw", "kc", "det",
		"oak", "hou", "tex", "laa", "sea",
		"atl", "mia", "nym", "phi", "wsh",
		"chc", "stl", "mil", "cin", "pit",
		"lad", "sd", "col", "ari", "sf"}
	teamNames := [30]string{
		"Yankees", "Rays", "Blue Jays", "Orioles", "Red Sox",
		"Twins", "Indians", "White Sox", "Royals", "Tigers",
		"Athletics", "Astros", "Rangers", "Angels", "Mariners",
		"Braves", "Marlins", "Mets", "Phillies", "Nationals",
		"Cubs", "Cardinals", "Brewers", "Reds", "Pirates",
		"Dodgers", "Padres", "Rockies", "Diamondbacks", "Giants"}
	var teamIDs []string

	for i, v := range teamAbrs {
		var idString = v + teamNames[i]
		teamIDs = append(teamIDs, idString)
	}

	var teamID = teamIDs[0]
	var teamRecord = internal.GetTeamRecord(teamID, "test-table")
	fmt.Println(teamRecord)
}
