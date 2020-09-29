package internal

import (
	"errors"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

// TeamStatsDynamoModel represents the dynamo team stats model
type TeamStatsDynamoModel struct {
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

// FormatDynamoRecord should unmarshal map the retrieved item
func FormatDynamoRecord(record *dynamodb.GetItemOutput, key string, tableName string) (*TeamStatsDynamoModel, error) {
	if record.Item == nil {
		msg := "Could not find '" + key + "' record in table " + tableName
		fmt.Println(msg)
		return nil, errors.New(msg)
	}

	// Format return value
	formattedRecord := TeamStatsDynamoModel{}
	err := dynamodbattribute.UnmarshalMap(record.Item, &formattedRecord)
	if err != nil {
		panic(fmt.Sprintf("Failed to unmarshal Record, %v", err))
	}

	return &formattedRecord, nil
}
