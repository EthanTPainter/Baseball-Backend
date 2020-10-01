package internal

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

// GetTeamRecord retrieves team record from Dynamo when provided team id
func GetTeamRecord(teamID string, tableName string) (*dynamodb.GetItemOutput, string, error) {
	teamKey := "team#" + teamID

	// Create session
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"),
	}))

	// Create new instance of service client
	svc := dynamodb.New(sess)
	input := &dynamodb.GetItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"partitionKey": {
				S: aws.String(teamKey),
			},
		},
		TableName: aws.String(tableName),
	}

	// Retrieve dynamo record
	result, err := svc.GetItem(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case dynamodb.ErrCodeProvisionedThroughputExceededException:
				fmt.Println(dynamodb.ErrCodeProvisionedThroughputExceededException, aerr.Error())
				return nil, "", aerr
			case dynamodb.ErrCodeResourceNotFoundException:
				fmt.Println(dynamodb.ErrCodeResourceNotFoundException, aerr.Error())
				return nil, "", aerr
			case dynamodb.ErrCodeRequestLimitExceeded:
				fmt.Println(dynamodb.ErrCodeRequestLimitExceeded, aerr.Error())
				return nil, "", aerr
			case dynamodb.ErrCodeInternalServerError:
				fmt.Println(dynamodb.ErrCodeInternalServerError, aerr.Error())
				return nil, "", aerr
			default:
				fmt.Println(aerr.Error())
				return nil, "", aerr
			}
		} else {
			// Print the error, cast to awserr.Error to get the Code and message from aws error
			fmt.Println(err.Error())
			return nil, "", aerr
		}
	}

	return result, teamKey, nil
}

// FormatTeamRecord should unmarshal map the retrieved item
func FormatTeamRecord(record *dynamodb.GetItemOutput, key string, tableName string) (TeamStatsViewModel, error) {
	if record.Item == nil {
		msg := "Could not find '" + key + "' record in table " + tableName
		fmt.Println(msg)
		return TeamStatsViewModel{}, errors.New(msg)
	}

	// Unmarshal dynamo item
	formattedRecord := TeamStatsDynamoModel{}
	err := dynamodbattribute.UnmarshalMap(record.Item, &formattedRecord)
	if err != nil {
		fmt.Println(err.Error())
		return TeamStatsViewModel{}, err
	}

	// Convert to view model
	viewModel, err2 := convertDynamoToView(formattedRecord)
	if err2 != nil {
		return TeamStatsViewModel{}, err2
	}

	return viewModel, nil
}

// Convert dynamo model to view model
func convertDynamoToView(record TeamStatsDynamoModel) (TeamStatsViewModel, error) {
	// Convert base (root level) stats
	baseStats, err1 := convertBaseStats(record)
	if err1 != nil {
		return TeamStatsViewModel{}, err1
	}

	// Convert hitting stats
	hittingStats, err2 := convertHittingStats(record.Hitting)
	if err2 != nil {
		return TeamStatsViewModel{}, err2
	}

	// Convert pitching stats
	pitchingStats, err3 := convertPitchingStats(record.Pitching)
	if err3 != nil {
		return TeamStatsViewModel{}, err3
	}

	// Convert fielding stats
	fieldingStats, err4 := convertFielingStats(record.Fielding)
	if err4 != nil {
		return TeamStatsViewModel{}, err4
	}

	viewModel := TeamStatsViewModel{
		Wins:        baseStats.Wins,
		Losses:      baseStats.Losses,
		GamesPlayed: baseStats.GamesPlayed,
		Hitting:     hittingStats,
		Pitching:    pitchingStats,
		Fielding:    fieldingStats,
	}
	return viewModel, nil
}

func convertBaseStats(record TeamStatsDynamoModel) (TeamBaseViewModel, error) {
	// Parse wins
	wins, wErr := strconv.Atoi(record.Wins)
	if wErr != nil {
		fmt.Println("Error converting wins (" + record.Wins + ") to int")
		return TeamBaseViewModel{}, wErr
	}

	// Parse losses
	losses, lErr := strconv.Atoi(record.Losses)
	if lErr != nil {
		fmt.Println("Error converting losses (" + record.Losses + ") to int")
		return TeamBaseViewModel{}, lErr
	}

	// Parse games played
	gamesPlayed, gpErr := strconv.Atoi(record.GamesPlayed)
	if gpErr != nil {
		fmt.Println("Error converting gamesPlayed (" + record.GamesPlayed + ") to int")
		return TeamBaseViewModel{}, gpErr
	}

	// Return model with converted root level stats
	return TeamBaseViewModel{Wins: wins, Losses: losses, GamesPlayed: gamesPlayed}, nil
}

// Convert hitting stats from strings to ints, floats
func convertHittingStats(record TeamHittingDynamoModel) (TeamHittingViewModel, error) {
	// Convert at bats to int
	atBats, atBatsErr := strconv.Atoi(record.AtBats)
	if atBatsErr != nil {
		fmt.Println("Error converting at bats (" + record.AtBats + ") to int")
		return TeamHittingViewModel{}, atBatsErr
	}

	// Convert average to float
	average, averageErr := strconv.ParseFloat(record.Average, 32)
	if averageErr != nil {
		fmt.Println("Error converting average (" + record.Average + ") to float64")
		return TeamHittingViewModel{}, averageErr
	}

	// Convert doubles to int
	doubles, doublesErr := strconv.Atoi(record.Doubles)
	if doublesErr != nil {
		fmt.Println("Error converting doubles (" + record.Doubles + ") to int")
		return TeamHittingViewModel{}, doublesErr
	}

	// Convert hits to int
	hits, hitsErr := strconv.Atoi(record.Hits)
	if hitsErr != nil {
		fmt.Println("Error converting hits (" + record.Hits + ") to int")
		return TeamHittingViewModel{}, hitsErr
	}

	// Convert home runs to int
	homeRuns, homeRunsErr := strconv.Atoi(record.HomeRuns)
	if homeRunsErr != nil {
		fmt.Println("Error converting home runs (" + record.HomeRuns + ") to int")
		return TeamHittingViewModel{}, homeRunsErr
	}

	// Convert on base and slugging to float
	onBaseAndSlugging, onBaseAndSluggingErr := strconv.ParseFloat(record.OnBaseAndSlugging, 32)
	if onBaseAndSluggingErr != nil {
		fmt.Println("Error converting onBaseAndSlugging (" + record.OnBaseAndSlugging + ") to float")
		return TeamHittingViewModel{}, onBaseAndSluggingErr
	}

	// Convert on base percentage to float
	onBasePercent, onBasePercentErr := strconv.ParseFloat(record.OnBasePercentage, 64)
	if onBasePercentErr != nil {
		fmt.Println("Error converting onBasePercentage (" + record.OnBasePercentage + ") to float64")
		return TeamHittingViewModel{}, onBasePercentErr
	}

	// Convert runs to int
	runs, runsErr := strconv.Atoi(record.Runs)
	if runsErr != nil {
		fmt.Println("Error converting runs (" + record.Runs + ") to int")
		return TeamHittingViewModel{}, runsErr
	}

	// Convert runs batted in to int
	runsBattedIn, runsBattedInErr := strconv.Atoi(record.RunsBattedIn)
	if runsBattedInErr != nil {
		fmt.Println("Error converting runsBattedIn (" + record.RunsBattedIn + ") to int")
		return TeamHittingViewModel{}, runsBattedInErr
	}

	// Convert slugging to float
	slugging, sluggingErr := strconv.ParseFloat(record.Slugging, 32)
	if sluggingErr != nil {
		fmt.Println("Error converting sluggingErr (" + record.Slugging + ") to float")
		return TeamHittingViewModel{}, sluggingErr
	}

	// Convert triples to int
	triples, triplesErr := strconv.Atoi(record.Triples)
	if triplesErr != nil {
		fmt.Println("Error converting triples (" + record.Triples + ") to int")
		return TeamHittingViewModel{}, triplesErr
	}

	// Convert strikeouts to int
	strikeouts, strikeoutsErr := strconv.Atoi(record.Strikeouts)
	if strikeoutsErr != nil {
		fmt.Println("Error converting strikeouts (" + record.Strikeouts + ") to int")
		return TeamHittingViewModel{}, strikeoutsErr
	}

	// Convert slugging to float
	slugging, sluggingErr := strconv.ParseFloat(record.Slugging, 32)
	if sluggingErr != nil {
		fmt.Println("Error converting slugging (" + record.Slugging + ") to float64")
		return TeamHittingViewModel{}, sluggingErr
	}

	// Convert walks to int
	walks, walksErr := strconv.Atoi(record.Walks)
	if walksErr != nil {
		fmt.Println("Error converting walks (" + record.Walks + ") to int")
		return TeamHittingViewModel{}, walksErr
	}

	// Convert leading batting average value to float
	leadingBattingAvgVal, leadingBattingAvgErr := strconv.ParseFloat(record.LeadingBattingAverage.Value, 32)
	if leadingBattingAvgErr != nil {
		fmt.Println("Error converting leadingBattingAverage (" + record.LeadingBattingAverage.Value + ") to float64")
		return TeamHittingViewModel{}, leadingBattingAvgErr
	}
	leadingBattingAvg := LeadingFloat{
		PlayerID: record.LeadingBattingAverage.PlayerID,
		Value:    leadingBattingAvgVal,
	}

	// Convert leading home runs to int
	leadingHomeRunsVal, leadingHomeRunsErr := strconv.Atoi(record.LeadingHomeRuns.Value)
	if leadingHomeRunsErr != nil {
		fmt.Println("Error converting leadingHomeRuns (" + record.LeadingHomeRuns.Value + ") to int")
		return TeamHittingViewModel{}, leadingHomeRunsErr
	}
	leadingHomeRuns := LeadingInt{
		PlayerID: record.LeadingHomeRuns.PlayerID,
		Value:    leadingHomeRunsVal,
	}

	// Convert leading runs batted in to int
	leadingRunsBattedInVal, leadingRunsBattedInErr := strconv.Atoi(record.LeadingRunsBattedIn.Value)
	if leadingRunsBattedInErr != nil {
		fmt.Println("Error converting leadingRunsBattedIn (" + record.LeadingRunsBattedIn.Value + ") to int")
		return TeamHittingViewModel{}, leadingRunsBattedInErr
	}
	leadingRunsBattedIn := LeadingInt{
		PlayerID: record.LeadingRunsBattedIn.PlayerID,
		Value:    leadingRunsBattedInVal,
	}

	// Convert leading on base percentage to float
	leadingOnBasePercentVal, leadingOnBasePercentErr := strconv.ParseFloat(record.LeadingOnBasePercentage.Value, 64)
	if leadingOnBasePercentErr != nil {
		fmt.Println("Error converting leading (" + record.LeadingOnBasePercentage.Value + ") to float")
		return TeamHittingViewModel{}, leadingOnBasePercentErr
	}
	leadingOnBasePercent := LeadingFloat{
		PlayerID: record.LeadingOnBasePercentage.PlayerID,
		Value:    leadingOnBasePercentVal,
	}

	// Convert leading hits to int
	leadingHitsVal, leadingHitsErr := strconv.Atoi(record.LeadingHits.Value)
	if leadingHitsErr != nil {
		fmt.Println("Error converting leading (" + record.LeadingHits.Value + ") to int")
		return TeamHittingViewModel{}, leadingHitsErr
	}
	leadingHits := LeadingInt{
		PlayerID: record.LeadingHits.PlayerID,
		Value:    leadingHitsVal,
	}

	return TeamHittingViewModel{
		AtBats:            atBats,
		Average:           average,
		Doubles:           doubles,
		Hits:              hits,
		HomeRuns:          homeRuns,
		OnBaseAndSlugging: onBaseAndSlugging,
		OnBasePercentage:  onBasePercent,

		Runs:         runs,
		RunsBattedIn: runsBattedIn,
		Slugging:     slugging,
		Strikeouts:   strikeouts,

		Triples: triples,
		Walks:   walks,

		LeadingBattingAverage:   leadingBattingAvg,
		LeadingHomeRuns:         leadingHomeRuns,
		LeadingRunsBattedIn:     leadingRunsBattedIn,
		LeadingOnBasePercentage: leadingOnBasePercent,
		LeadingHits:             leadingHits,
	}, nil
}

// Convert pitching stats from strings to ints, floats
func convertPitchingStats(record TeamPitchingDynamoModel) (TeamPitchingViewModel, error) {
	// Convert innings pitched to float
	inningsPitched, inningsPitchedErr := strconv.ParseFloat(record.InningsPitched, 32)
	if inningsPitchedErr != nil {
		fmt.Println("Error converting innings pitched (" + record.InningsPitched + ") to float")
		return TeamPitchingViewModel{}, inningsPitchedErr
	}

	// Convert runs to int
	runs, runsErr := strconv.Atoi(record.Runs)
	if runsErr != nil {
		fmt.Println("Error converting runs (" + record.Runs + ") to int")
		return TeamPitchingViewModel{}, runsErr
	}

	// Convert hits to int
	hits, hitsErr := strconv.Atoi(record.Hits)
	if hitsErr != nil {
		fmt.Println("Error converting hits (" + record.Hits + ") to int")
		return TeamPitchingViewModel{}, hitsErr
	}

	// Convert doubles to int
	doubles, doublesErr := strconv.Atoi(record.Doubles)
	if doublesErr != nil {
		fmt.Println("Error converting doubles (" + record.Doubles + ") to int")
		return TeamPitchingViewModel{}, doublesErr
	}

	// Convert triples to int
	triples, triplesErr := strconv.Atoi(record.Triples)
	if triplesErr != nil {
		fmt.Println("Error converting triples (" + record.Triples + ") to int")
		return TeamPitchingViewModel{}, triplesErr
	}

	// Convert home runs to int
	homeRuns, homeRunsErr := strconv.Atoi(record.HomeRuns)
	if homeRunsErr != nil {
		fmt.Println("Error converting home runs (" + record.HomeRuns + ") to int")
		return TeamPitchingViewModel{}, homeRunsErr
	}

	// Convert walks to int
	walks, walksErr := strconv.Atoi(record.Walks)
	if walksErr != nil {
		fmt.Println("Error converting walks (" + record.Walks + ") to int")
		return TeamPitchingViewModel{}, walksErr
	}

	// Convert strikeouts to int
	strikeouts, strikeoutsErr := strconv.Atoi(record.Strikeouts)
	if strikeoutsErr != nil {
		fmt.Println("Error converting strikeouts (" + record.Strikeouts + ") to int")
		return TeamPitchingViewModel{}, strikeoutsErr
	}

	// Convert strikeouts per nine to float
	strikeoutsPerNine, strikeoutsPerNineErr := strconv.ParseFloat(record.StrikeoutsPerNine, 32)
	if strikeoutsPerNineErr != nil {
		fmt.Println("Error converting strikeouts per nine (" + record.StrikeoutsPerNine + ") to float")
		return TeamPitchingViewModel{}, strikeoutsPerNineErr
	}

	// Convert pitches per start to float
	pitchesPerStart, pitchesPerStartErr := strconv.ParseFloat(record.PitchesPerStart, 32)
	if pitchesPerStartErr != nil {

	}

	// Convert walks plus hits per innings pitched to float
	walksPlusHitsPerInnings, walksPlusHitsPerInningsErr := strconv.ParseFloat(record.WalksPlusHitsPerInningsPitched, 32)
	if walksPlusHitsPerInningsErr != nil {

	}

	// Convert earned run average to float
	earnedRunAverage, earnedRunAverageErr := strconv.ParseFloat(record.EarnedRunAverage, 32)
	if earnedRunAverageErr != nil {

	}

	// Convert saves to int
	saves, savesErr := strconv.Atoi(record.Saves)
	if savesErr != nil {

	}

	// Convert stolen bases to int
	stolenBases, stolenBasesErr := strconv.Atoi(record.StolenBases)
	if stolenBasesErr != nil {

	}

	// Convert leading wins to int
	leadingWinsVal, leadingWinsErr := strconv.Atoi(record.LeadingWins.Value)
	if leadingWinsErr != nil {
		fmt.Println("Error converting leading wins (" + record.LeadingWins.Value + ") to int")
		return TeamPitchingViewModel{}, leadingWinsErr
	}
	leadingWins := LeadingInt{
		PlayerID: record.LeadingWins.PlayerID,
		Value:    leadingWinsVal,
	}

	// Convert leading earned run average to float
	leadingEarnedRunAverageVal, leadingEarnedRunAverageErr := strconv.ParseFloat(record.LeadingEarnedRunAverage.Value, 32)
	if leadingEarnedRunAverageErr != nil {
		fmt.Println("Error converting leading earned run average (" + record.LeadingEarnedRunAverage.Value + ") to float")
		return TeamPitchingViewModel{}, leadingEarnedRunAverageErr
	}
	leadingEarnedRunAverage := LeadingFloat{
		PlayerID: record.LeadingEarnedRunAverage.PlayerID,
		Value:    leadingEarnedRunAverageVal,
	}

	// Convert leading strikeouts to int
	leadingStrikeoutsVal, leadingStrikeoutsErr := strconv.Atoi(record.LeadingStrikeouts.Value)
	if leadingStrikeoutsErr != nil {
		fmt.Println("Error converting leading strikeouts (" + record.LeadingStrikeouts.Value + ") to int")
		return TeamPitchingViewModel{}, leadingStrikeoutsErr
	}
	leadingStrikeouts := LeadingInt{
		PlayerID: record.LeadingStrikeouts.PlayerID,
		Value:    leadingStrikeoutsVal,
	}

	// Convert leading saves to int
	leadingSavesVal, leadingSavesErr := strconv.Atoi(record.LeadingSaves.Value)
	if leadingSavesErr != nil {
		fmt.Println("Error converting leading saves (" + record.LeadingSaves.Value + ") to int")
		return TeamPitchingViewModel{}, leadingSavesErr
	}
	leadingSaves := LeadingInt{
		PlayerID: record.LeadingSaves.PlayerID,
		Value:    leadingSavesVal,
	}

	// Convert leading holds to int
	leadingHoldsVal, leadingHoldsErr := strconv.Atoi(record.LeadingHolds.Value)
	if leadingHoldsErr != nil {
		fmt.Println("Error converting leading holds (" + record.LeadingHolds.Value + ") to int")
		return TeamPitchingViewModel{}, leadingHoldsErr
	}
	leadingHolds := LeadingInt{
		PlayerID: record.LeadingHolds.PlayerID,
		Value:    leadingHoldsVal,
	}

	return TeamPitchingViewModel{
		InningsPitched:                 inningsPitched,
		Runs:                           runs,
		Hits:                           hits,
		Doubles:                        doubles,
		Triples:                        triples,
		HomeRuns:                       homeRuns,
		Strikeouts:                     strikeouts,
		StrikeoutsPerNine:              strikeoutsPerNine,
		Walks:                          walks,
		WalksPlusHitsPerInningsPitched: walksPlusHitsPerInnings,
		PitchesPerStart:                pitchesPerStart,
		EarnedRunAverage:               earnedRunAverage,
		Saves:                          saves,
		StolenBases:                    stolenBases,
		LeadingWins:                    leadingWins,
		LeadingEarnedRunAverage:        leadingEarnedRunAverage,
		LeadingStrikeouts:              leadingStrikeouts,
		LeadingSaves:                   leadingSaves,
		LeadingHolds:                   leadingHolds,
	}, nil
}

// Convert fieldingstats from strings to ints, floats
func convertFielingStats(record TeamFieldingDynamoModel) (TeamFieldingViewModel, error) {
	// Parse Errors
	errors, eErr := strconv.Atoi(record.Errors)
	if eErr != nil {
		fmt.Println()
		return TeamFieldingViewModel{}, eErr
	}

	// Parse fielding percentage
	fieldingPercent, fpError := strconv.ParseFloat(record.FieldingPercentage, 64)
	if fpError != nil {
		fmt.Println()
		return TeamFieldingViewModel{}, fpError
	}

	return TeamFieldingViewModel{Errors: errors, FieldingPercentage: fieldingPercent}, nil
}
