import logging

# Python Lambda Docs: https://docs.aws.amazon.com/lambda/latest/dg/python-handler.html
def main(event, context):
  logging.info("Python function processed a request")

  teamInfo = [
    # AL EAST
    { abr: 'nyy', nickname: 'Yankees' },
    { abr: 'tb', nickname: 'Rays' },
    { abr: 'tor', nickname: 'Blue Jays' },
    { abr: 'bal', nickname: 'Orioles' },
    { abr: 'bos', nickname: 'Red Sox' },
    # AL CENTRAL
    { abr: 'min', nickname: 'Twins' },
    { abr: 'cle', nickname: 'Indians' },
    { abr: 'chw', nickname: 'White Sox' },
    { abr: 'kc', nickname: 'Royals' },
    { abr: 'det', nickname: 'Tigers' },
    # AL WEST
    { abr: 'oak', nickname: 'Athletics' },
    { abr: 'hou', nickname: 'Astros' },
    { abr: 'tex', nickname: 'Rangers' },
    { abr: 'laa', nickname: 'Angels' },
    { abr: 'sea', nickname: 'Mariners' },
    # NL EAST
    { abr: 'atl', nickname: 'Braves' },
    { abr: 'mia', nickname: 'Marlins' },
    { abr: 'nym', nickname: 'Mets' },
    { abr: 'phi', nickname: 'Phillies' },
    { abr: 'wsh', nickname: 'Nationals' },
    # NL CENTRAL
    { abr: 'chc', nickname: 'Cubs' },
    { abr: 'stl', nickname: 'Cardinals' },
    { abr: 'mil', nickname: 'Brewers' },
    { abr: 'cin', nickname: 'Reds' },
    { abr: 'pit', nickname: 'Pirates' },
    # NL WEST
    { abr: 'lad', nickname: 'Dodges' },
    { abr: 'sd', nickname: 'Padres' },
    { abr: 'col', nickname: 'Rockies' },
    { abr: 'ari', nickname: 'Diamonbacks' },
    { abr: 'sf', nickname: 'Giants' },
  ]
  
  for teamAbr in teamAbbreviations:

  return {
    'message': "Successfully scraped data"
  }
