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
	AtBats                  int
	Average                 float64
	Doubles                 int
	Hits                    int
	HomeRuns                int
	OnBaseAndSlugging       float64
	OnBasePercentage        float64
	Runs                    int
	RunsBattedIn            int
	Slugging                float64
	Strikeouts              int
	TotalBases              int
	Triples                 int
	Walks                   int
	LeadingBattingAverage   LeadingFloat
	LeadingHomeRuns         LeadingInt
	LeadingRunsBattedIn     LeadingInt
	LeadingOnBasePercentage LeadingFloat
	LeadingHits             LeadingInt
}

// TeamPitchingViewModel is the view model for team pitching stats
type TeamPitchingViewModel struct {
	EarnedRunAverage               float64
	Hits                           int
	Holds                          int
	HomeRuns                       int
	InningsPitched                 float64
	Losses                         int
	PitchesPerStart                float64
	QualityStarts                  int
	Saves                          int
	Strikeouts                     int
	StrikeoutsPerNine              float64
	Walks                          int
	WalksPlusHitsPerInningsPitched float64
	Wins                           int
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
