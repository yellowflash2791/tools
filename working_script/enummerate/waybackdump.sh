sed '/.jpg/d;/.png/d;/.gif/d;/.css/d' -i wayback > wayback1
sort wayback1|uniq > wayback2
cat wayback2|/var/root/go/bin/httpx -sr
