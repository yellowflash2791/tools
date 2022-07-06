#tool is for directiry listing using fuff ./fuff_file.sh hosts


mkdir report_para
while read -r line1;
do

line2="?FUZZ="$2
l3=".html"
line3=$(echo $line1|sed 's/http:\/\///g;s/https:\/\///g;s/\./_/g')

/var/root/go/bin/ffuf -w /Users/arnold/Desktop/tools/wordlist/para.txt -u  $line1$line2 -od new_dir_all -mc 200 -of html -o report_para/$line3$l3 -debug-log debug.log

done < $1
