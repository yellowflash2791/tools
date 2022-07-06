
#/var/root/go/bin/assetfinder --subs-only $1 > assetfinder.txt;
#echo "assetfinder hosts: `cat assetfinder.txt|wc -l`"
#amass enum -d $1 -o amass.txt;
#echo "amass hosts: `cat amass.txt|wc -l`" ;
#subfinder_work -d $1 -v -o subfinder.txt
#echo "subfinder hosts: `cat subfinder.txt|wc -l`" ;
#sort assetfinder.txt amass.txt subfinder.txt|uniq > domains;
#echo "total hosts found: `cat domains|wc -l`";
#cat domains|/var/root/go/bin/httprobe > hosts;
echo "total http hosts found: `cat hosts|wc -l`";
cat hosts|/var/root/go/bin/waybackurls > wayback
cat hosts|/var/root/go/bin/gau > gau
