import boto3

def getTeamRecord(id, tableName):
  dynamodb = boto3.resource("dynamodb")
  table = dynamodb.Table(tableName)
  record = table.get_item(
    Key = {
      "partitionKey": id
    },
    ReturnConsumedCapacity = "NONE"
  )
  return record

def updateTeamRecord(id, tableName, newRecord):
  dynamodb = boto3.resource("dynamodb")
  table = dynamodb.Table(tableName)
  record = table.update_item(
    Key = {
      "partitionKey": id
    },
    ExpressionAttributeNames = {
      "#W": "Wins",
      "#L": "Losses",
      "#DR": "DivisionRank",
      "#G": "GamesPlayed",
      "#H": "Hitting",
      "#P": "Pitching",
      "#F": "Fielding"
    },
    ExpressionAttributeValues = {
      ":w": newRecord["Wins"],
      ":l": newRecord["Losses"],
      ":dr": newRecord["DivisionRank"],
      ":g": newRecord["GamesPlayed"],
      ":h": newRecord["Hitting"],
      ":p": newRecord["Pitching"],
      ":f": newRecord["Fielding"]
    },
    UpdateExpression = "SET #W = :w, #L = :l, #DR = :dr, #G = :g, #H = :h, #P = :p, #F = :f",
    ReturnConsumedCapacity = "NONE"
  )
  return record