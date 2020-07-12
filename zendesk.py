import requests

APIURL = "http://www.zendesk.com/api/v2/tickets.json"

response = requests.get(APIURL)

print(response)