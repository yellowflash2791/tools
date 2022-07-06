while read -r line;
         do 
            /var/root/go/bin/subdomain_takeover $line test
         done < $1
