import requests
from bs4 import BeautifulSoup

def scrapeTeamSchedule(teamAbr, teamName):
  # Visit team web page
  baseUrl = f"https://espn.com/mlb/team/_/name/{teamAbr}"
  page = requests.get(baseUrl)
  soup = BeautifulSoup(page.content, 'html.parser')
  teamSchedule = soup.select("section .club-schedule > ul > ul:nth-of-type(1) > li")

  # Includes all previous played games and next upcoming game
  currentGames = len(teamSchedule)
  trackedGames = 0
  recordedGames = []

  while trackedGames < currentGames:
    # Get Game Id from href (if the game has already been played)
    gameHref = teamSchedule[trackedGames].a["href"];
    gameId = None
    if ("gameId" in gameHref):
      gameId = gameHref[gameHref.index("gameId") + len("gameId="):]

    opponentHTML = teamSchedule[trackedGames].select("a > div:nth-of-type(2)")
    resultHTML = teamSchedule[trackedGames].select("a > div:nth-of-type(3) > div")
    gameResult = resultHTML[0].string.upper()
    opponent = opponentHTML[0].string.replace("vs", "").replace("@", "")

    # Check if game is suspended or postponed
    if gameResult == "PPD" or gameResult == "SUSP":
      game = {
        "opponent": opponent,
        "result": "PPD"
      }
      recordedGames.append(game)

    # Check if game is won or lost
    if gameResult == "W" or gameResult == "L":
      game = {
        "gameId": gameId,
        "opponent": opponent,
        "result": resultHTML[1].string,
        "winner": teamName if gameResult == "W" else opponent
      }
      recordedGames.append(game)
    trackedGames += 1
  return recordedGames
