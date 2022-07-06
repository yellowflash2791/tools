while read -r line;
         do 
            /var/root/go/bin/lfi $line "test"
         done < $1
