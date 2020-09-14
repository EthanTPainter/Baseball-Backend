from bs4 import BeautifulSoup
import requests

def getTeamSchedulePage(teamAbbreviation):
  baseUrl = f"https://espn.com/mlb/team/_/name/{teamAbbreviation}"
  page = requests.get(baseUrl)
  soup = BeautifulSoup(page.content, 'html.parser')
  teamScheduleHTML = soup.select("section .club-schedule > ul > ul:nth-of-type(1) > li")
  return teamSchedule

def getTeamBattingStatsPage(teamAbbreviation):
  baseUrl = f"https://espn.com/mlb/team/stats/_/name/{teamAbbreviation}"
  page = requests.get(baseUrl)
  soup = BeautifulSoup(page.content, 'html.parser')
  battingStatsHTML = soup.select("tbody .TABLE__TBODY > tr:last-child")
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
  pitchingStatsHTML = soup.select("tbody .TABLE__TBODY > tr:last-child")
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
  fieldingStats = soup.select("tbody .TABLE__TBODY > tr:last-child")
  return fieldingStats
