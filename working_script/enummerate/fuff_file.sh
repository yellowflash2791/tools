#tool is for directiry listing using fuff ./fuff_file.sh hosts


mkdir report
while read -r line1;
do

line2="/FUZZ"
l3=".html"
line3=$(echo $line1|sed 's/http:\/\///g;s/https:\/\///g;s/\./_/g')

/var/root/go/bin/ffuf -w /Users/arnold/Desktop/tools/wordlist/fuzz.txt -u  $line1$line2 -mc 200 -od new_dir_all -of html -o report/$line3$l3 -debug-log debug.log

done < $1
