while read -r line; 

do

echo $line

/var/root/go/bin/meg -v -d 1000 \/$line

rm -r out/;

done ;

