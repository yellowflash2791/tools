import re

test="""

APP_URL=https://tagview-slack-bot.herokuapp.com/
BOT_TOKEN=xoxb-xxxx
PHABRICATOR_URL=http://phabricator.tagview.com.br/
PHABRICATOR_DB_URL=mysql://user:pass@host/db?connectionLimit=10
PHABRICATOR_S3_BUCKET=happy_bucket
AWS_KEY_ID=AKIAHDGXYSA2DF5IKW6A
AWS_SECRET_ACCESS_KEY=s3cr3t
SLACK_WEBHOOK_URL=http://hooks.slack.com/services


"""

r=re.search("((A3T[A-Z0-9]|AKIA|AGPA|AIDA|AROA|AIPA|ANPA|ANVA|ASIA)[A-Z0-9]{16})",test)
if r:
    print r.group(0)
