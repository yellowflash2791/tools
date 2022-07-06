import config
import json
import requests
import os
import time, threading


def notifySlack(message):
    if not config.SLACK_WEBHOOKURL:
        print('Please define Slack Webhook URL to enable notifications')
        exit()
    requests.post(config.SLACK_WEBHOOKURL, json={'text': ':new:'+message})


def foo():
  os.system("cat dir.txt|grep '(200 OK)' > new_dir.txt")
  f = open("new_dir.txt", "r")
  line=f.read()
  notifySlack(line)


while True:
  foo()
  time.sleep(14400)
