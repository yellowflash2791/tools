cat gau|grep -e .js$|tee jslinks
cat wayback|grep -e .js$ >> jslinks
sort jslinks|uniq > jslinks_main
rm jslinks
while IFS= read -r line
do
  echo "$line"|httpx -sr -timeout 10
done < jslinks_main

cd output
gf keywords|tee ../gf/gau_wayback_keyword.txt
