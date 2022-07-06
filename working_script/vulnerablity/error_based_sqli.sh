while read -r line;
         do 
            /var/root/go/bin/error_based_sqli $line test
         done < $1
