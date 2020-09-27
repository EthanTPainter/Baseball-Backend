import os
import sys
import json
import logging
import asyncio

# Team abbreviations
teamAbrs = [
  'nyy', 'tb', 'tor', 'bal', 'bos',
  'min', 'cle', 'chw', 'kc', 'det',
  'oak', 'hou', 'tex', 'laa', 'sea',
  'atl', 'mia', 'nym', 'phi', 'wsh',
  'chc', 'stl', 'mil', 'cin', 'pit',
  'lad', 'sd', 'col', 'ari', 'sf'
]
# Team Names
teamNames = [
  'Yankees', 'Rays', 'Blue Jays', 'Orioles', 'Red Sox',
  'Twins', 'Indians', 'White Sox', 'Royals', 'Tigers',
  'Athletics', 'Astros', 'Rangers', 'Angels', 'Mariners',
  'Braves', 'Marlins', 'Mets', 'Phillies', 'Nationals',
  'Cubs', 'Cardinals', 'Brewers', 'Reds', 'Pirates',
  'Dodgers', 'Padres', 'Rockies', 'Diamondbacks', 'Giants'
]

# Python Lambda Docs: https://docs.aws.amazon.com/lambda/latest/dg/python-handler.html
def main(event, context):
  for abr in teamAbrs:
    # Create teamId from abrs and names
    index = teamAbrs.index(abr)
    name = teamNames[index]
    teamId = "team#" + abr + name

    # tableName = os.environ["TABLE_NAME"]
    tableName = "test-table"

    # player statistics
    

if __name__ == "__main__":
  print("Executing as main program")
  asyncio.run(main("e", "t"))