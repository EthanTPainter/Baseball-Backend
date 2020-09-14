import os
import sys
import json
import logging
import asyncio

from dynamoClient import getTeamRecord
from statNavigator import getTeamRecord, getLeadingBattingStats, getLeadingPitchingStats, getTeamBattingStats, getTeamPitchingStats, getTeamFieldingStats

# Team abbreviations
teamAbrs = ['nyy']
# teamAbrs = [
#   'nyy', 'tb', 'tor', 'bal', 'bos',
#   'min', 'cle', 'chw', 'kc', 'det',
#   'oak', 'hou', 'tex', 'laa', 'sea',
#   'atl', 'mia', 'nym', 'phi', 'wsh',
#   'chc', 'stl', 'mil', 'cin', 'pit',
#   'lad', 'sd', 'col', 'ari', 'sf'
# ]
teamNames = [
  'Yankees', 'Rays', 'Blue Jays', 'Orioles', 'Red Sox',
  'Twins', 'Indians', 'White Sox', 'Royals', 'Tigers',
  'Atheltics', 'Astros', 'Rangers', 'Angels', 'Mariners',
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
    # So convert stats to other types during new record creation
    teamRecord = getTeamRecord(abr)
    battingLeaders = getLeadingBattingStats(abr)
    pitchingLeaders = getLeadingPitchingStats(abr)
    teamBatting = getTeamBattingStats(abr)
    teamPitching = getTeamPitchingStats(abr)
    teamFielding = getTeamFieldingStats(abr)

    newRecord = {
      "Wins": int(teamRecord["W"]),
      "Losses": int(teamRecord["L"]),
      "divisionRank": teamRecord["divisionalRank"],
      "GamesPlayed": int(teamRecord["W"]) + int(teamRecord["L"]),
      "Hitting": {
        "atBats": int(teamBatting["AB"]),
        "Runs": int(teamBatting["R"]),
        "Hits": int(teamBatting["H"]),
        "Doubles": int(teamBatting["2B"]),
        "Triples": int(teamBatting["3B"]),
        "HomeRuns": int(teamBatting["HR"]),
        "RunsBattedIn": int(teamBatting["RBI"]),
        "TotalBases": int(teamBatting["TB"]),
        "Walks": int(teamBatting["BB"]),
        "Strikeouts": int(teamBatting["SO"]),
        "Average": float(teamBatting["AVG"]),
        "Slugging": float(teamBatting["SLG"]),
        "OnBasePercentage": float(teamBatting["OBP"]),
        "OnBaseAndSlugging": float(teamBatting["OPS"]),
        "LeadingBattingAverage": {
          "playerId": battingLeaders["AVG"]["id"],
          "value": float(battingLeaders["AVG"]["val"]),
        },
        "LeadingHomeRuns": {
          "playerId": battingLeaders["HR"]["id"],
          "value": int(battingLeaders["HR"]["val"])
        },
        "LeadingRunsBattedIn": {
          "playerId": battingLeaders["RBI"]["id"],
          "value": int(battingLeaders["RBI"]["val"])
        },
        "LeadingOnBasePercentage": {
          "playerId": battingLeaders["OBP"]["id"],
          "value": float(battingLeaders["OBP"]["val"])
        },
        "LeadingHits": {
          "playerId": battingLeaders["H"]["id"],
          "value": int(battingLeaders["H"]["val"])
        }
      },
      "Pitching": {
        "InningsPitched": float(teamPitching["IP"]),
        "QualityStarts": int(teamPitching["QS"]),
        "Wins": int(teamPitching["W"]),
        "Losses": int(teamPitching["L"]),
        "Saves": int(teamPitching["SV"]),
        "Holds": int(teamPitching["HLD"]),
        "Hits": int(teamPitching["H"]),
        "HomeRuns": int(teamPitching["HR"]),
        "Walks": int(teamPitching["BB"]),
        "Strikeouts": int(teamPitching["K"]),
        "StrikeoutsPerNine": float(teamPitching["K/9"]),
        "PitchesPerStart": float(teamPitching["P/S"]),
        "WalksPlusHitsPerInningsPitched": float(teamPitching["WHIP"]),
        "EarnedRunAverage": float(teamPitching["ERA"]),
        "LeadingWins": {
          "playerId": pitchingLeaders["W"]["id"],
          "value": int(pitchingLeaders["W"]["val"]),
        },
        "LeadingEarnedRunAverage": {
          "playerId": pitchingLeaders["ERA"]["id"],
          "value": float(pitchingLeaders["ERA"]["val"])
        },
        "LeadingStrikeouts": {
          "playerId": pitchingLeaders["K"]["id"],
          "value": int(pitchingLeaders["K"]["val"])
        },
        "LeadingSaves": {
          "playerId": pitchingLeaders["SV"]["id"],
          "value": int(pitchingLeaders["SV"]["val"])
        },
        "LeadingHolds": {
          "playerId": pitchingLeaders["HLD"]["id"],
          "value": int(pitchingLeaders["HLD"]["val"])
        }
      },
      "Fielding": {
        "Errors": teamFielding["E"],
        "FieldingPercentage": teamFielding["FP"]
      }
    }
    print(newRecord)

    # Update team record with retrieved statistics
    # updateTeamRecord(teamId, tableName, newRecord)


  return {
    'message': "Successfully scraped data"
  }

if __name__ == "__main__":
  print("Executing as main program")
  asyncio.run(main("e", "t"))