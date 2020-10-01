package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/ethantpainter/Baseball-Backend/packages/team-stats/src/internal"
)

var teamAbrs = [30]string{
	"nyy", "tb", "tor", "bal", "bos",
	"min", "cle", "chw", "kc", "det",
	"oak", "hou", "tex", "laa", "sea",
	"atl", "mia", "nym", "phi", "wsh",
	"chc", "stl", "mil", "cin", "pit",
	"lad", "sd", "col", "ari", "sf"}

var teamNames = [30]string{
	"Yankees", "Rays", "Blue Jays", "Orioles", "Red Sox",
	"Twins", "Indians", "White Sox", "Royals", "Tigers",
	"Athletics", "Astros", "Rangers", "Angels", "Mariners",
	"Braves", "Marlins", "Mets", "Phillies", "Nationals",
	"Cubs", "Cardinals", "Brewers", "Reds", "Pirates",
	"Dodgers", "Padres", "Rockies", "Diamondbacks", "Giants"}

// TeamStatsLambdaEvent is the expected lambda event for the team stats lambda
type TeamStatsLambdaEvent struct {
	teamAbrs  []string
	teamNames []string
	playerID  string
}

// TeamStatsResponse is the expected response for the team stats lambda
type TeamStatsResponse struct {
	Response     internal.TeamStatsViewModel
	ErrorMessage string `json:"errorMessage"`
}

// HandleLambdaEvent handles all lambda event progress
func HandleLambdaEvent(event TeamStatsLambdaEvent) (TeamStatsResponse, error) {
	// Check if body is empty
	if event.teamAbrs == nil && event.teamNames == nil && event.playerID == "" {
		errorMsg := "Error: no expected parameters found (teamAbrs, teamNames, playerID)"
		fmt.Println(errorMsg)
		return TeamStatsResponse{
			ErrorMessage: errorMsg,
		}, nil
	}

	var teamIDs []string

	for i, v := range teamAbrs {
		var idString = v + teamNames[i]
		teamIDs = append(teamIDs, idString)
	}

	teamID := teamIDs[0]
	tableName := os.Getenv("TABLE_NAME")

	// retrieve team record from dynamo
	teamRecord, teamKey, getErr := internal.GetTeamRecord(teamID, tableName)
	if getErr != nil {
		return TeamStatsResponse{
			ErrorMessage: getErr.Error(),
		}, nil
	}

	// Format the team record from a data model to a view model
	formattedRecord, formatErr := internal.FormatTeamRecord(teamRecord, teamKey, tableName)
	if formatErr != nil {
		return TeamStatsResponse{
			ErrorMessage: formatErr.Error(),
		}, nil
	}

	return TeamStatsResponse{
		Response: formattedRecord,
	}, nil
}

func main() {
	lambda.Start(HandleLambdaEvent)
}
