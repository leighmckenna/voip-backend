import json
import boto3

def sms_mms_webhook_lambda(event, context):
    # TODO implement
    dynamodb = boto3.resource('dynamodb')
    
    table = dynamodb.Table('message_table')
    to_number = event["data"]["payload"]["to"][0]["phone_number"]
    message = event["data"]["payload"]["text"]
    
    response = table.get_item(
        Key={
            'number': to_number
        }
    )
    
    if response.get('Item') is not None:
        item = response['Item']
        item['message'] = message
        item.update()
    
    table.put_item(
        Item={
            'number': to_number,
            'message': message,
        }
    )
    
    return {
        'statusCode': 200,
    }
