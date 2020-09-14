from webNavigator import getTeamRecordPage, getLeadingBattingStatsPage, getLeadingPitchingStatsPage, getTeamBattingStatsPage, getTeamPitchingStatsPage, getTeamFieldingStatsPage

def getTeamRecord(teamAbr):
  teamRecordHTML = getTeamRecordPage(teamAbr)

  # Get Soup HTML for record
  record = teamRecordHTML.select("li:first-child")[0].get_text()
  regionalRank = teamRecordHTML.select("li:last-child")[0].get_text()

  # Create json blob as return value
  teamRecord = {
    "W": record.split("-")[0],
    "L": record.split("-")[1],
    "divisionalRank": regionalRank
  }
  return teamRecord

def getLeadingBattingStats(teamAbr):
  leadBattingStatsHTML = getLeadingBattingStatsPage(teamAbr)

  # Get Soup HTML for different stat blocks
  battingAvgHTML = leadBattingStatsHTML[0].select("section > a:first-child")[0]
  homeRunsHTML = leadBattingStatsHTML[0].select("section > a:nth-of-type(2)")[0]
  runsBattedInHTML = leadBattingStatsHTML[0].select("section > a:nth-of-type(3)")[0]
  onBasePercentageHTML = leadBattingStatsHTML[0].select("section > a:nth-of-type(4)")[0]
  hitsHTML = leadBattingStatsHTML[0].select("section > a:last-child")[0]

  # Get player id from each blocks
  battingAvgPlayerId = battingAvgHTML.attrs["href"].split("id/")[1]
  homeRunsPlayedId = homeRunsHTML.attrs["href"].split("id/")[1]
  runsBattedInPlayedId = homeRunsHTML.attrs["href"].split("id/")[1]
  onBasePercentagePlayedId = onBasePercentageHTML.attrs["href"].split("id/")[1]
  hitsPlayedId = hitsHTML.attrs["href"].split("id/")[1]

  # Get stat value from each block
  battingAvg = battingAvgHTML.select("div > div > div > div")[0].get_text()
  homeRuns = homeRunsHTML.select("div > div > div > div")[0].get_text()
  runsBattedIn = runsBattedInHTML.select("div > div > div> div")[0].get_text()
  onBasePercentage = onBasePercentageHTML.select("div > div > div > div")[0].get_text()
  hits = hitsHTML.select("div > div > div > div")[0].get_text()

  # Create JSON blob as return value
  leads = {
    "AVG": {
      "id": battingAvgPlayerId,
      "val": battingAvg,
    },
    "HRs": {
      "id": homeRunsPlayedId,
      "val": homeRuns
    },
    "RBIs": {
      "id": runsBattedInPlayedId,
      "val": runsBattedIn
    },
    "OBP": {
      "id": onBasePercentagePlayedId,
      "val": onBasePercentage
    },
    "H": {
      "id": hitsPlayedId,
      "val": hits
    }
  }

  return leads

def getLeadingPitchingStats(teamAbr):
  leadPitchingStatsHTML = getLeadingPitchingStatsPage(teamAbr)

  # Get Soup HTML for different stat blocks
  winsHTML = leadPitchingStatsHTML[0].select("section > a:first-child")[0]
  earnedRunAverageHTML = leadPitchingStatsHTML[0].select("section > a:nth-of-type(2)")[0]
  strikeoutsHTML = leadPitchingStatsHTML[0].select("section > a:nth-of-type(3)")[0]
  savesHTML = leadPitchingStatsHTML[0].select("section > a:nth-of-type(4)")[0]
  holdsHTML = leadPitchingStatsHTML[0].select("section > a:last-child")[0]

  # Get player id from each blocks
  winsPlayedId = winsHTML.attrs["href"].split("id/")[1]
  earnedRunAveragePlayedId = earnedRunAverageHTML.attrs["href"].split("id/")[1]
  strikeoutsPlayedId = strikeoutsHTML.attrs["href"].split("id/")[1]
  savesPlayedId = savesHTML.attrs["href"].split("id/")[1]
  holdsPlayedId = holdsHTML.attrs["href"].split("id/")[1]

  # Get stat value from each block
  wins = winsHTML.select("div > div > div > div")[0].get_text()
  earnedRunAverage = earnedRunAverageHTML.select("div > div > div > div")[0].get_text()
  strikeouts = strikeoutsHTML.select("div > div > div> div")[0].get_text()
  saves = savesHTML.select("div > div > div > div")[0].get_text()
  holds = holdsHTML.select("div > div > div > div")[0].get_text()

  # Create json blob as return value
  leads = {
    "Ws": {
      "id": winsPlayedId,
      "val": wins
    },
    "ERA": {
      "id": earnedRunAveragePlayedId,
      "val": earnedRunAverage
    },
    "Ks": {
      "id": strikeoutsPlayedId,
      "val": strikeouts
    },
    "SVs": {
      "id": savesPlayedId,
      "val": saves
    },
    "HLD": {
      "id": holdsPlayedId,
      "val": holds
    }
  }

  return leads

def getTeamBattingStats(teamAbr):
  teamBattingStatsHTML = getTeamBattingStatsPage(teamAbr)

  # Parse table team totals
  gamesPlayed = teamBattingStatsHTML.select("td:first-child > span")[0].get_text()
  atBats = teamBattingStatsHTML.select("td:nth-of-type(2) > span")[0].get_text()
  runs = teamBattingStatsHTML.select("td:nth-of-type(3) > span")[0].get_text()
  hits = teamBattingStatsHTML.select("td:nth-of-type(4) > span")[0].get_text()
  doubles = teamBattingStatsHTML.select("td:nth-of-type(5) > span")[0].get_text()
  triples = teamBattingStatsHTML.select("td:nth-of-type(6) > span")[0].get_text()
  homeRuns = teamBattingStatsHTML.select("td:nth-of-type(7) > span")[0].get_text()
  runsBattedIn = teamBattingStatsHTML.select("td:nth-of-type(8) > span")[0].get_text()
  totalBases = teamBattingStatsHTML.select("td:nth-of-type(9) > span")[0].get_text()
  walks = teamBattingStatsHTML.select("td:nth-of-type(10) > span")[0].get_text()
  strikeouts = teamBattingStatsHTML.select("td:nth-of-type(11) > span")[0].get_text()
  stolenBases = teamBattingStatsHTML.select("td:nth-of-type(12) > span")[0].get_text()
  average = teamBattingStatsHTML.select("td:nth-of-type(13) > span")[0].get_text()
  onBasePercentage = teamBattingStatsHTML.select("td:nth-of-type(14) > span")[0].get_text()
  sluggingPercentage = teamBattingStatsHTML.select("td:nth-of-type(15) > span")[0].get_text()
  onBasePercentagePlusSlugging = teamBattingStatsHTML.select("td:nth-of-type(16) > span")[0].get_text()

  # Create json blob as return value
  teamHitting = {
    "GP": gamesPlayed,
    "AB": atBats,
    "R": runs,
    "H": hits,
    "2B": doubles,
    "3B": triples,
    "HR": homeRuns,
    "RBI": runsBattedIn,
    "TB": totalBases,
    "BB": walks,
    "SO": strikeouts,
    "SB": stolenBases,
    "AVG": average,
    "OBP": onBasePercentage,
    "SLG": sluggingPercentage,
    "OPS": onBasePercentagePlusSlugging
  }
  return teamHitting

def getTeamPitchingStats(teamAbr):
  teamPitchingStatsHTML = getTeamPitchingStatsPage(teamAbr)

  # Parse table team totals
  gamesPlayed = teamPitchingStatsHTML.select("td:first-child > span")[0].get_text()
  gamesStarted = teamPitchingStatsHTML.select("td:nth-of-type(2) > span")[0].get_text()
  qualityStarts = teamPitchingStatsHTML.select("td:nth-of-type(3) > span")[0].get_text()
  wins = teamPitchingStatsHTML.select("td:nth-of-type(4) > span")[0].get_text()
  losses = teamPitchingStatsHTML.select("td:nth-of-type(5) > span")[0].get_text()
  saves = teamPitchingStatsHTML.select("td:nth-of-type(6) > span")[0].get_text()
  holds = teamPitchingStatsHTML.select("td:nth-of-type(7) > span")[0].get_text()
  inningsPitched = teamPitchingStatsHTML.select("td:nth-of-type(8) > span")[0].get_text()
  hits = teamPitchingStatsHTML.select("td:nth-of-type(9) > span")[0].get_text()
  earnedRuns = teamPitchingStatsHTML.select("td:nth-of-type(10) > span")[0].get_text()
  homeRuns = teamPitchingStatsHTML.select("td:nth-of-type(11) > span")[0].get_text()
  walks = teamPitchingStatsHTML.select("td:nth-of-type(12) > span")[0].get_text()
  strikeouts = teamPitchingStatsHTML.select("td:nth-of-type(13) > span")[0].get_text()
  strikeoutsPerNine = teamPitchingStatsHTML.select("td:nth-of-type(14) > span")[0].get_text()
  pitchesPerStart = teamPitchingStatsHTML.select("td:nth-of-type(15) > span")[0].get_text()
  # Skip 16 since WAR doesn't have a team total
  walksPlusHitsPerInningPitched = teamPitchingStatsHTML.select("td:nth-of-type(17) > span")[0].get_text()
  earnedRunAverage = teamPitchingStatsHTML.select("td:nth-of-type(18) > span")[0].get_text()

  # Create json blob as return value
  teamPitching = {
    "GP": gamesPlayed,
    "GS": gamesStarted,
    "QS": qualityStarts,
    "W": wins,
    "L": losses,
    "SV": saves,
    "HLD": holds,
    "IP": inningsPitched,
    "H": hits,
    "ER": earnedRuns,
    "HR": homeRuns,
    "BB": walks,
    "K": strikeouts,
    "K/9": strikeoutsPerNine,
    "P/S": pitchesPerStart,
    "WHIP": walksPlusHitsPerInningPitched,
    "ERA": earnedRunAverage
  }
  return teamPitching

def getTeamFieldingStats(teamAbr):
  teamFieldingStatsHTML = getTeamFieldingStatsPage(teamAbr)

  # Parse table team totals
  gamesPlayed = teamFieldingStatsHTML.select("td:first-child > span")[0].get_text()
  fieldingPercentage = teamFieldingStatsHTML.select("td:nth-of-type(7) > span")[0].get_text()
  errors = teamFieldingStatsHTML.select("td:nth-of-type(8) > span")[0].get_text()

  # Create json blob as return value
  teamFielding = {
    "GP": gamesPlayed,
    "FP": fieldingPercentage,
    "E": errors
  }
  return teamFielding