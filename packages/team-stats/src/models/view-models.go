package models

// TeamStatsViewModel represents the team stats view model
type TeamStatsViewModel struct {
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
			PlayerID string
			Value    float32
		}
		LeadingHomeRuns struct {
			PlayerID string
			Value    float32
		}
		LeadingRunsBattedIn struct {
			PlayerID string
			Value    float32
		}
		LeadingOnBasePercentage struct {
			PlayerID string
			Value    float32
		}
		LeadingHits struct {
			PlayerID string
			Value    float32
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
			PlayerID string
			Value    int
		}
		LeadingEarnedRunAverage struct {
			PlayerID string
			Value    float32
		}
		LeadingStrikeouts struct {
			PlayerID string
			Value    int
		}
		LeadingSaves struct {
			PlayerID string
			Value    int
		}
		LeadingHolds struct {
			PlayerID string
			Value    int
		}
	}

	Fielding struct {
		Errors             int
		FieldingPercentage float32
	}
}

// TeamBaseViewModel is the view model for basic team stats
type TeamBaseViewModel struct {
	Wins        int
	Losses      int
	GamesPlayed int
}

// TeamHittingViewModel is the view model for team hitting stats
type TeamHittingViewModel struct {
	Runs                      int
	Hits                      int
	Doubles                   int
	Triples                   int
	RunsBattedIn              int
	PitchesPerPlateAppearance int
	Walks                     int
	Strikeouts                int
	Average                   int
	Slugging                  float64
	OnBasePercentage          float64
	LeadingBattingAverage     struct {
		PlayerID string
		value    float64
	}
	LeadingHomeRuns struct {
		PlayerID string
		Value    float64
	}
	LeadingRunsBattedIn struct {
		PlayerID string
		Value    float64
	}
	LeadingOnBasePercentage struct {
		PlayerID string
		Value    float64
	}
	LeadingHits struct {
		PlayerID string
		Value    float64
	}
}

// TeamPitchingViewModel is the view model for team pitching stats
type TeamPitchingViewModel struct {
	InningsPitched                 float64
	Runs                           int
	Hits                           int
	Doubles                        int
	Triples                        int
	HomeRuns                       int
	Walks                          int
	Strikeouts                     int
	StrikeoutsPerNine              int
	PitchesPerStart                float64
	WalksPlusHitsPerInningsPitched float64
	EarnedRunAverage               float64
	Saves                          float64
	StolenBases                    int

	LeadingWins struct {
		PlayerID string
		Value    int
	}
	LeadingEarnedRunAverage struct {
		PlayerID string
		Value    float64
	}
	LeadingStrikeouts struct {
		PlayerID string
		Value    int
	}
	LeadingSaves struct {
		PlayerID string
		Value    int
	}
	LeadingHolds struct {
		PlayerID string
		Value    int
	}
}

// TeamFieldingViewModel is the view model for team fielding stats
type TeamFieldingViewModel struct {
	Errors             int
	FieldingPercentage float64
}
