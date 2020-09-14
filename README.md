# Baseball-Backend
Baseball application backend for tracking team or player performance across a baseball season. The frontend application can be found [here](https://github.com/EthanTPainter/Baseball-Frontend)

# Architecture Diagram

[ Insert diagram here when fully built]

## Explanation

There are three web scraper lambdas for updating game, team, and player statistics. These lambdas are run as cron jobs each day at 7 AM UTC (3 EST) time. 

There are three lambdas that are used directly for retrieving game, team, and player statistics. 

The data is stored in a DynamoDB table, and partitioned in various ways 

# Database Partitioning

## Game Partitioning

Game partitioning has the following partition key format: `game#{gameId}` where `gameId` is the id for the game played. The sort key will also exist and have a format: `{homeTeamId}#{awayTeamId}` where `homeTeamId` is the team id of the home team, and `awayTeamId` is the away team id.

The following properties could also be present on a game record:

```
{
   HTRuns: number;
   HTHits: number;
   HTErrors: number;
   HTRunsByInning: number[];
   HTAtBats: number;
   HTRBIs: number;
   HTWalks: number;
   HTRuns: number;
   HTInningsPitched: number;
   HTEarnedRuns: number;
   HTStrikeouts: number;
   HTPitchCount: number;
   HTStrikeCount: number;
   HTBallCount: number;  

   ATHits: number;
   ATErrors: number;
   ATRunsByInning: number[];
   ATAtBats: number;
   ATRBIs: number;
   ATWalks: number;
   ATRuns: number;
   ATInningsPitched: number;
   ATEarnedRuns: number;
   ATStrikeouts: number;
   ATPitchCount: number;
   ATStrikeCount: number;
   ATBallCount: number;  

   WinningPitcher: string;
   LosingPitcher: string;
}
```

## Team Partitioning
Team partition has the following parititon key format: `team#{teamId}` where `teamId` is the id for the team. The team dynamo record does not have a sort key.

The following properties could also be present on a team record:

```
{
  Wins: number;
  Losses: number;
  GamesPlayed: number;

  Hitting: {
    Runs: number;
    Hits: number;
    Doubles: number;
    Triples: number;
    HomeRuns: number;
    RunsBattedIn: number;
    PitchesPerPlateAppearance: number;
    Walks: number;
    Strikeouts: number;
    Average: float;
    Slugging: float;
    OnBasePercentage: float;
    OnBaseAndSlugging: float;

    LeadingBattingAverage: {
      playerId: string;
      value: float;
    }
    LeadingHomeRuns: {
      playerId: string;
      value: number;
    }
    LeadingRunsBattedIn: {
      playerId: string;
      value: number;
    }
    LeadingOnBasePercentage: {
      playerId: string;
      value: float;
    }
    LeadingHits: {
      palyerId: string;
      value: number;
    }
  }

  Pitching: {
    InningsPitched: float;
    Runs: number;
    Hits: number;
    Doubles: number;
    Triples: number;
    HomeRuns: number;
    Walks: number;
    Strikeouts: number;
    StrikeoutsPerNine: number;
    PitchesPerStart: float;
    WalksPlusHitsPerInningsPitched: float;
    EarnedRunAverage: float;
    Saves: number;
    StolenBases: number;

    LeadingWins: {
      playerId: string;
      value: number;
    }
    LeadingEarnedRunAverage: {
      playerId: string;
      value: float;
    }
    LeadingStrikeouts: {
      playerId: string;
      value: number;
    }
    LeadingSaves: {
      playerId: string;
      value: number;
    }
    LeadingHolds: {
      playerId: string;
      value: number;
    }
  }

  Fielding: {
    Errors: number;
    FieldingPercentage: float;
  }
}
```

## Player Partitioning

Player partitioning has the following partition key format: `player#{playerId}` where `playerId` is the id for the player. The player dynamo record has a sort key of `team#{teamId}` where `teamId` is the id of the team the player plays for.

The saved properties vary based on the position of the player. The following properties could also be present on a hitting player:

```
{
  GamesPlayed: number;
  GamesStarted: number;
  AtBats: number;
  Runs: number;
  Hits: number;
  Doubles: number;
  Triples: number;
  HomeRuns: number;
  RunsBattedIn: number;
  Walks: number;
  HitByPitch: number;
  Strikeouts: number;
  StolenBases: number;
  CaughtStealing: number;
  Average: float;
  OnBasePercentage: float;
  Slugging: float;
  OnBaseAndSlugging: float;
  WinsAboveReplacement: float;
  FullInnings: float;
  Pickoffs: number;
  Assists: number;
  DoublePlays: number;
  FieldingPercentage: float;
  Errors: number;

  Height: {
    feet: number;
    inches: number;
  }
  Weight: number;
  DateOfBirth: string;
  BatSide: string;
  ThrowSide: string;
  Birthplace: string;
  Position: string;
  Number: string;
  Name: {
    first: string;
    last: string;
  }
}
```

The following properties could also be present on a pitching player:

```
{
  GamesPlayed: number;
  GamesStarted: number;
  Wins: number;
  Losses: number;
  WinPercentage: float;
  WinsAboveReplacement: float;
  EarnedRunaAverage: float;
  InningsPitched: float;
  Strikeouts: number;
  Walks: number;
  StrikeoutToWalkRatio: float;
  Hits: number;
  Runs: number;
  EarnedRuns: number;
  Saves: number;
  Holds: number;
  BlownSaves: number;

  PitchesPerStart: float;
  PitchersPerInning: float;
  StrikeoutsPerNine: float;
  CompleteGames: number;
  Shutouts: number;
  AverageGameScore: float;
  GroundBalls: number;
  FlyBalls: number;
  GroundToFlyRation: float;
  InheritedRuns: number;
  InheritedRunsScored: number;
  Balks: number;
  StolenBases: number;
  CaughtStealing: number;
  Run support: float;

  OpponentBattingAverage: float;
  OpponentOnBasePercentage: float;
  OpponentSluggingPercentage: float;
  OpponentDoubles: number;
  OpponentTriples: number;
  OpponentHomeRuns: number;
  OpponentTotalBases: number;
  OpponentRunsBattedIn: number;
  BattersHits: number;
  IntentionalWalks: number;

  Height: {
    feet: number;
    inches: number;
  }
  Weight: number;
  DateOfBirth: string;
  BatSide: string;
  ThrowSide: string;
  Birthplace: string;
  Position: string;
  Number: string;
  Name: {
    first: string;
    last: string;
  }
}
```