package models

// TeamStatsDynamoModel represents the dynamo team stats model
type TeamStatsDynamoModel struct {
	Wins        string
	Losses      string
	GamesPlayed string

	Hitting struct {
		Runs                      string
		Hits                      string
		Doubles                   string
		Triples                   string
		RunsBattedIn              string
		PitchesPerPlateAppearance string
		Walks                     string
		Strikeouts                string
		Average                   string
		Slugging                  string
		OnBasePercentage          string
		LeadingBattingAverage     struct {
			PlayerID string
			Value    string
		}
		LeadingHomeRuns struct {
			PlayerID string
			Value    string
		}
		LeadingRunsBattedIn struct {
			PlayerID string
			Value    string
		}
		LeadingOnBasePercentage struct {
			PlayerID string
			Value    string
		}
		LeadingHits struct {
			PlayerID string
			Value    string
		}
	}

	Pitching struct {
		InningsPitched                 string
		Runs                           string
		Hits                           string
		Doubles                        string
		Triples                        string
		HomeRuns                       string
		Walks                          string
		Strikeouts                     string
		StrikeoutsPerNine              string
		PitchesPerStart                string
		WalksPlusHitsPerInningsPitched string
		EarnedRunAverage               string
		Saves                          string
		StolenBases                    string

		LeadingWins struct {
			PlayerID string
			Value    string
		}
		LeadingEarnedRunAverage struct {
			PlayerID string
			Value    string
		}
		LeadingStrikeouts struct {
			PlayerID string
			Value    string
		}
		LeadingSaves struct {
			PlayerID string
			Value    string
		}
		LeadingHolds struct {
			PlayerID string
			Value    string
		}
	}

	Fielding struct {
		Errors             string
		FieldingPercentage string
	}
}

// TeamHittingDynamoModel is the db model for
type TeamHittingDynamoModel struct {
	Runs                      string
	Hits                      string
	Doubles                   string
	Triples                   string
	RunsBattedIn              string
	PitchesPerPlateAppearance string
	Walks                     string
	Strikeouts                string
	Average                   string
	Slugging                  string
	OnBasePercentage          string
	LeadingBattingAverage     struct {
		PlayerID string
		value    string
	}
	LeadingHomeRuns struct {
		PlayerID string
		Value    string
	}
	LeadingRunsBattedIn struct {
		PlayerID string
		Value    string
	}
	LeadingOnBasePercentage struct {
		PlayerID string
		Value    string
	}
	LeadingHits struct {
		PlayerID string
		Value    string
	}
}

// TeamPitchingDynamoModel is the db model for team pitching stats
type TeamPitchingDynamoModel struct {
}

// TeamFieldingDynamoModel is the db model for team fielding stats
type TeamFieldingDynamoModel struct {
	Errors             string
	FieldingPercentage string
}
