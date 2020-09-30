package internal

// TeamStatsViewModel represents the team stats view model
type TeamStatsViewModel struct {
	Wins        int
	Losses      int
	GamesPlayed int
	Hitting     TeamHittingViewModel
	Pitching    TeamPitchingViewModel
	Fielding    TeamFieldingViewModel
}

// TeamBaseViewModel is the view model for basic team stats
type TeamBaseViewModel struct {
	Wins        int
	Losses      int
	GamesPlayed int
}

// LeadingInt represents a leader model with an int value
type LeadingInt struct {
	PlayerID string
	Value    int
}

// LeadingFloat represents a leader model with a float value
type LeadingFloat struct {
	PlayerID string
	Value    float64
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
	Average                   float64
	Slugging                  float64
	OnBasePercentage          float64
	LeadingBattingAverage     LeadingFloat
	LeadingHomeRuns           LeadingInt
	LeadingRunsBattedIn       LeadingInt
	LeadingOnBasePercentage   LeadingFloat
	LeadingHits               LeadingInt
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
	StrikeoutsPerNine              float64
	PitchesPerStart                float64
	WalksPlusHitsPerInningsPitched float64
	EarnedRunAverage               float64
	Saves                          int
	StolenBases                    int
	LeadingWins                    LeadingInt
	LeadingEarnedRunAverage        LeadingFloat
	LeadingStrikeouts              LeadingInt
	LeadingSaves                   LeadingInt
	LeadingHolds                   LeadingInt
}

// TeamFieldingViewModel is the view model for team fielding stats
type TeamFieldingViewModel struct {
	Errors             int
	FieldingPercentage float64
}
