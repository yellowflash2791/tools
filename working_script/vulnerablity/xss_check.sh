while read -r line;
         do 
            /var/root/go/bin/xss_check $line test
         done < $1
