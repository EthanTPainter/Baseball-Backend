import boto3

def getPlayerRecord(id, tableName):
  dynamodb = boto3.resource("dynamodb")
  table = dynamodb.Table(tableName)
  record = table.get_item(
    Key = {
      "partitionKey": id
    },
    ReturnConsumedCapacity = "NONE"
  )
  return record

def updatePlayerRecord(id, tableName, newRecord):
  dynamodb = boto3.resource("dynamodb")
  table = dynamodb.Table(tableName)
  record = table.update_item(
    Key = {
      "partitionKey": id
    },
    ExpressionAttributeNames = {

    },
    ExpressionAttributeValues = {

    },
    UpdateExpression = "SET ",
    ReturnConsumedCapacity = "NONE"
  )
  return record