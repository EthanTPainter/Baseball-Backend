import os
import sys
import json
import logging
import asyncio

from dynamoClient import getTeamRecord, updateTeamRecord
from statNavigator import getTeamRecord, getLeadingBattingStats, getLeadingPitchingStats, getTeamBattingStats, getTeamPitchingStats, getTeamFieldingStats

# Team abbreviations
teamAbrs = [
  'nyy', 'tb', 'tor', 'bal', 'bos',
  'min', 'cle', 'chw', 'kc', 'det',
  'oak', 'hou', 'tex', 'laa', 'sea',
  'atl', 'mia', 'nym', 'phi', 'wsh',
  'chc', 'stl', 'mil', 'cin', 'pit',
  'lad', 'sd', 'col', 'ari', 'sf'
]
teamNames = [
  'Yankees', 'Rays', 'Blue Jays', 'Orioles', 'Red Sox',
  'Twins', 'Indians', 'White Sox', 'Royals', 'Tigers',
  'Athletics', 'Astros', 'Rangers', 'Angels', 'Mariners',
  'Braves', 'Marlins', 'Mets', 'Phillies', 'Nationals',
  'Cubs', 'Cardinals', 'Brewers', 'Reds', 'Pirates',
  'Dodgers', 'Padres', 'Rockies', 'Diamonbacks', 'Giants'
]

# Python Lambda Docs: https://docs.aws.amazon.com/lambda/latest/dg/python-handler.html
def main(event, context):
  for abr in teamAbrs:
    # Create teamId from abr and name
    index = teamAbrs.index(abr)
    name = teamNames[index]
    teamId = "team#" + abr + name

    # tableName = os.environ["TABLE_NAME"]
    tableName = "test-table"

    # team statistics
    # All stats recorded as strings in each json blob
    teamRecord = getTeamRecord(abr)
    battingLeaders = getLeadingBattingStats(abr)
    pitchingLeaders = getLeadingPitchingStats(abr)
    teamBatting = getTeamBattingStats(abr)
    teamPitching = getTeamPitchingStats(abr)
    teamFielding = getTeamFieldingStats(abr)

    newRecord = {
      "Wins": teamRecord["W"],
      "Losses": teamRecord["L"],
      "DivisionRank": teamRecord["divisionalRank"],
      "GamesPlayed": teamRecord["W"] + teamRecord["L"],
      "Hitting": {
        "atBats": teamBatting["AB"],
        "Runs": teamBatting["R"],
        "Hits": teamBatting["H"],
        "Doubles": teamBatting["2B"],
        "Triples": teamBatting["3B"],
        "HomeRuns": teamBatting["HR"],
        "RunsBattedIn": teamBatting["RBI"],
        "TotalBases": teamBatting["TB"],
        "Walks": teamBatting["BB"],
        "Strikeouts": teamBatting["SO"],
        "Average": teamBatting["AVG"],
        "Slugging": teamBatting["SLG"],
        "OnBasePercentage": teamBatting["OBP"],
        "OnBaseAndSlugging": teamBatting["OPS"],
        "LeadingBattingAverage": {
          "playerId": battingLeaders["AVG"]["id"],
          "value": battingLeaders["AVG"]["val"],
        },
        "LeadingHomeRuns": {
          "playerId": battingLeaders["HR"]["id"],
          "value": battingLeaders["HR"]["val"]
        },
        "LeadingRunsBattedIn": {
          "playerId": battingLeaders["RBI"]["id"],
          "value": battingLeaders["RBI"]["val"]
        },
        "LeadingOnBasePercentage": {
          "playerId": battingLeaders["OBP"]["id"],
          "value": battingLeaders["OBP"]["val"]
        },
        "LeadingHits": {
          "playerId": battingLeaders["H"]["id"],
          "value": battingLeaders["H"]["val"]
        }
      },
      "Pitching": {
        "InningsPitched": teamPitching["IP"],
        "QualityStarts": teamPitching["QS"],
        "Wins": teamPitching["W"],
        "Losses": teamPitching["L"],
        "Saves": teamPitching["SV"],
        "Holds": teamPitching["HLD"],
        "Hits": teamPitching["H"],
        "HomeRuns": teamPitching["HR"],
        "Walks": teamPitching["BB"],
        "Strikeouts": teamPitching["K"],
        "StrikeoutsPerNine": teamPitching["K/9"],
        "PitchesPerStart": teamPitching["P/S"],
        "WalksPlusHitsPerInningsPitched": teamPitching["WHIP"],
        "EarnedRunAverage": teamPitching["ERA"],
        "LeadingWins": {
          "playerId": pitchingLeaders["W"]["id"],
          "value": pitchingLeaders["W"]["val"],
        },
        "LeadingEarnedRunAverage": {
          "playerId": pitchingLeaders["ERA"]["id"],
          "value": pitchingLeaders["ERA"]["val"]
        },
        "LeadingStrikeouts": {
          "playerId": pitchingLeaders["K"]["id"],
          "value": pitchingLeaders["K"]["val"]
        },
        "LeadingSaves": {
          "playerId": pitchingLeaders["SV"]["id"],
          "value": pitchingLeaders["SV"]["val"]
        },
        "LeadingHolds": {
          "playerId": pitchingLeaders["HLD"]["id"],
          "value": pitchingLeaders["HLD"]["val"]
        }
      },
      "Fielding": {
        "Errors": teamFielding["E"],
        "FieldingPercentage": teamFielding["FP"]
      }
    }

    # Update team record with retrieved statistics
    updateTeamRecord(teamId, tableName, newRecord)


  return {
    'message': "Successfully scraped data"
  }

if __name__ == "__main__":
  print("Executing as main program")
  asyncio.run(main("e", "t"))