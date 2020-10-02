package main

import (
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

// HandleLambdaEvent handles all lambda event progress
func HandleLambdaEvent(event internal.TeamStatsLambdaEvent) (internal.TeamStatsResponse, error) {
	// Parse body if the expected properties exist
	// Priority: teamAbrs <- teamNames <- playerIDs
	parsedBody, errorMessage := internal.ParseLambdaEventBody(event, teamAbrs, teamNames)
	if errorMessage != "" {
		return internal.TeamStatsResponse{
			ErrorMessage: errorMessage,
		}, nil
	}

	// Generate team IDs
	var teamIDs []string
	for i, v := range teamAbrs {
		var idString = v + teamNames[i]
		teamIDs = append(teamIDs, idString)
	}
	teamID := teamIDs[0]
	tableName := os.Getenv("TABLE_NAME")

	// Retrieve team record from dynamo
	teamRecord, teamKey, getErr := internal.GetTeamRecord(teamID, tableName)
	if getErr != nil {
		return internal.TeamStatsResponse{
			ErrorMessage: getErr.Error(),
		}, nil
	}

	// Format the team record from a data model to a view model
	formattedRecord, formatErr := internal.FormatTeamRecord(teamRecord, teamKey, tableName)
	if formatErr != nil {
		return internal.TeamStatsResponse{
			ErrorMessage: formatErr.Error(),
		}, nil
	}

	return internal.TeamStatsResponse{
		Response: formattedRecord,
	}, nil
}

func main() {
	lambda.Start(HandleLambdaEvent)
}
