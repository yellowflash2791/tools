while read -r line;
         do 
            /var/root/go/bin/s3-bucket-check $line 
         done < $1
