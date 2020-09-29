package main

import (
	"fmt"

	"github.com/ethantpainter/Baseball-Backend/packages/team-stats/src/internal"
)

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

	teamID := teamIDs[0]
	tableName := "test-table"
	teamRecord, teamKey, getErr := internal.GetTeamRecord(teamID, tableName)
	// Error encountered during team record retrieval
	if getErr != nil {
		return
	}

	formattedRecord, formatErr := internal.FormatTeamRecord(teamRecord, teamKey, tableName)
	// Error encountered during formatting team record
	if formatErr != nil {
		return
	}

	fmt.Println(formattedRecord)
}
