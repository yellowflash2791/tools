while read -r line; do python3 /Users/arnold/Desktop/tools/dirsearch-master/dirsearch.py -u $line -e html,php,jsp,asp,json,js; done < hosts
