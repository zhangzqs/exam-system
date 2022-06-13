import json

import requests


def upload(que):
    repo = requests.post(
        url='http://localhost:8080/questions/',
        headers={
            'Authorization': 'Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9'
                             '.eyJleHAiOjE2NTYwMDUzOTAsInVpZCI6MX0'
                             '.9963c4czMbv1sxr58lmZ0yTIAN-6NkvnUzFvdYY4Vbg',
            'Content-Type': 'application/json'
        },
        data=json.dumps(que))
    print(repo.request.body)
    print(repo.content.decode('utf-8'))


with open('output.json', encoding='utf-8') as f:
    for que in json.load(f):
        upload(que)
