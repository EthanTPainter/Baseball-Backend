from bs4 import BeautifulSoup
import requests

def getTeamRoster(teamAbbreviation):
  baseUrl = f"https://www.espn.com/mlb/team/roster/_/name/${teamAbbreviation}"
  page = requests.get(baseUrl)
  soup = BeautifulSoup(page.content, "html.parser")
  rosterHTML = soup.select("")
  return rosterHTML

