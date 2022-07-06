echo './all_check.sh output.txt http://burpcollab hosts'

cookie="acs_ngn=tid%3D3c138bab-7cec-4bef-8b5b-cdb99e6f91bb%26eid%3DAAAA033685054%26e%3D1%26a%3DdGVzdCB0ZXN0%26u%3D129b946f-0a8b-4918-9a2d-1352fedfe5a4%26t%3D1617347521%26h%3D9c7bf4b48da26b3ea15c9172b5fa2ebf;"

while read -r line;
         do 
            /var/root/go/bin/ssti_check $line $cookie
            /var/root/go/bin/error_based_sqli $line $cookie
#            /var/root/go/bin/sql_header $line $cookie
            /var/root/go/bin/ssrf_para $line $2 $cookie
#            /var/root/go/bin/time_sql $line $cookie
#            /var/root/go/bin/command_injector $line $cookie
            /var/root/go/bin/lfi $line $cookie
            /var/root/go/bin/ssrf_header $line $2 $cookie
            /var/root/go/bin/xss_check $line $cookie
            /var/root/go/bin/url_path_based_sqli $line $cookie
            /var/root/go/bin/directory_traversal $line $cookie
            /var/root/go/bin/jsonp $line $cookie

         done < $1


