while read -r line1;
do
 python3 /Users/arnold/Desktop/tools/dirsearch-master/dirsearch.py -e * --plain-text-report=PLAINTEXTOUTPUTFILE  -u $line1
 cat PLAINTEXTOUTPUTFILE |grep '403\|401'|grep -Eo "(http|https)://.+" >> for_url.txt 
 cat PLAINTEXTOUTPUTFILE |grep '200\|201'|grep -Eo "(http|https)://.+" >> accessible_url.txt
done < $1

while read -r line; 
do 
  
  /var/root/go/bin/4xx_bypass $line |tee bypassurl  

done < for_url.txt

