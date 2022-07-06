while read -r line
      do

       /var/root/go/bin/4xx_bypass $line      

      done < $1
