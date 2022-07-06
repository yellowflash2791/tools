#./meg_non_dir.sh hosts "directory which you wanna check"

while read -r line1;
do

line2=$2
echo $line1$line2 >> new_file.txt

done < $1

cat new_file.txt| /var/root/go/bin/httpx -status-code -content-length|grep 200|tee 200.txt
cut -d' ' -f1 200.txt > gowit.txt


while read -r line4;
do

/var/root/go/bin/gowitness single $line4

done < gowit.txt


rm 200.txt
#rm gowit.txt
rm new_file.txt



