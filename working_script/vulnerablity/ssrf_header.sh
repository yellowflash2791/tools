while read -r line;
         do 
            /var/root/go/bin/ssrf_header $line $2 "test"
         done < $1
