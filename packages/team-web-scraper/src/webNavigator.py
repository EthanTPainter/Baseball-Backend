from bs4 import BeautifulSoup
import requests

def getTeamSchedulePage(teamAbbreviation):
  baseUrl = f"https://espn.com/mlb/team/_/name/{teamAbbreviation}"
  page = requests.get(baseUrl)
  soup = BeautifulSoup(page.content, 'html.parser')
  teamScheduleHTML = soup.select("section .club-schedule > ul > ul:nth-of-type(1) > li")
  return teamSchedule

def getTeamRecordPage(teamAbbreviation):
  baseUrl = f"https://espn.com/mlb/team/_/name/{teamAbbreviation}"
  page = requests.get(baseUrl)
  soup = BeautifulSoup(page.content, 'html.parser')
  teamRecordHTML = soup.select("ul.ClubhouseHeader__Record")[0]
  return teamRecordHTML

def getTeamBattingStatsPage(teamAbbreviation):
  baseUrl = f"https://espn.com/mlb/team/stats/_/name/{teamAbbreviation}"
  page = requests.get(baseUrl)
  soup = BeautifulSoup(page.content, 'html.parser')
  tableBodies = soup.select("tbody > tr:last-child")
  # There are 2 tables, each with 2 bodies. First is for general stats, Second is advanced
  # Possible TODO: add second table body to team stats
  battingStatsHTML = tableBodies[1]
  return battingStatsHTML

def getLeadingBattingStatsPage(teamAbbreviation):
  baseUrl = f"https://espn.com/mlb/team/stats/_/name/{teamAbbreviation}"
  page = requests.get(baseUrl)
  soup = BeautifulSoup(page.content, 'html.parser')
  battingStatsHTML = soup.select("section .StatLeaders")
  return battingStatsHTML

def getTeamPitchingStatsPage(teamAbbreviation):
  baseUrl = f"https://espn.com/mlb/team/stats/_/type/pitching/name/{teamAbbreviation}"
  page = requests.get(baseUrl)
  soup = BeautifulSoup(page.content, 'html.parser')
  tableBodies = soup.select("tbody > tr:last-child")
  # There are 2 tables, each with 2 bodies. First is for general stats, Second is advanced
  # Possible TODO: add second table body to team stats
  pitchingStatsHTML = tableBodies[1]
  return pitchingStatsHTML

def getLeadingPitchingStatsPage(teamAbbreviation):
  baseUrl = f"https://espn.com/mlb/team/stats/_/type/pitching/name/{teamAbbreviation}"
  page = requests.get(baseUrl)
  soup = BeautifulSoup(page.content, 'html.parser')
  pitchingStatsHTML = soup.select("section .StatLeaders")
  return pitchingStatsHTML

def getTeamFieldingStatsPage(teamAbbreviation):
  baseUrl = f"https://espn.com/mlb/team/stats/_/type/fielding/name/{teamAbbreviation}"
  page = requests.get(baseUrl)
  soup = BeautifulSoup(page.content, 'html.parser')
  tableBodies = soup.select("tbody > tr:last-child")
  fieldingStats = tableBodies[1]
  return fieldingStats
