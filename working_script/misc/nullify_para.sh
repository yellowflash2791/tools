while read -r line;
         do 
            /var/root/go/bin/para_fuzz $line "FUZZ" 1
         done         
      done < $1
