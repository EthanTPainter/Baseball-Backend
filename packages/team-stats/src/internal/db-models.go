package internal

// TeamStatsDynamoModel represents the dynamo team stats model
type TeamStatsDynamoModel struct {
	Wins        string
	Losses      string
	GamesPlayed string
	Hitting     TeamHittingDynamoModel
	Pitching    TeamPitchingDynamoModel
	Fielding    TeamFieldingDynamoModel
}

// TeamHittingDynamoModel is the db model for
type TeamHittingDynamoModel struct {
	AtBats            string
	Average           string
	Doubles           string
	Hits              string
	HomeRuns          string
	OnBaseAndSlugging string
	OnBasePercentage  string
	Runs              string
	RunsBattedIn      string
	Slugging          string
	Strikeouts        string
	TotalBases        string
	Triples           string
	Walks             string

	// Leading Stats
	LeadingBattingAverage struct {
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

// TeamPitchingDynamoModel is the db model for team pitching stats
type TeamPitchingDynamoModel struct {
	EarnedRunAverage               string
	Hits                           string
	Holds                          string
	HomeRuns                       string
	InningsPitched                 string
	Losses                         string
	PitchesPerStart                string
	QualityStarts                  string
	Saves                          string
	Strikeouts                     string
	StrikeoutsPerNine              string
	Walks                          string
	WalksPlusHitsPerInningsPitched string
	Wins                           string

	// Leading stats
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

// TeamFieldingDynamoModel is the db model for team fielding stats
type TeamFieldingDynamoModel struct {
	Errors             string
	FieldingPercentage string
}
