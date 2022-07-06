
cat hosts|sed 's~http[s]*://~~g' > hosts1;
sort hosts1 | uniq > uniqhost.txt
mkdir param
cd param/
cat ../uniqhost.txt|while read -r line;do python3 /Users/arnold/Desktop/tools/ParamSpider/paramspider.py -d $line --exclude woff,css,png,svg,jpg -o $line.txt;done;
ls output/*.txt > txtfile.txt;
#cat txtfile.txt| while read -r line1;do /var/root/go/bin/dalfox -b https://infosecflash.xss.ht file $line1 -o output/dalfox.txt;done;
