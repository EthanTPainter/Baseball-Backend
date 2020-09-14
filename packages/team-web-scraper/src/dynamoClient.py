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
  record = table.get_item(
    Key = {
      "partitionKey": id
    },
    UpdateExpression = "",
    ReturnConsumedCapacity = "NONE"
  )
  return record