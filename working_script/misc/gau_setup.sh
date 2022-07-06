mkdir gau_process
cp gau gau_process/
cd gau_process/
grep -Er "\?\w+=|\?\w+\&\w+=" gau|tee para_gau
sed -i -e 's/gau://g' para_gau
nullify_para para_gau|tee null_gau
sed -i '' '/.woff/d;/.png/d;/.gif/d;/.jpg/d;/.jpeg/d;/.css/d;/.js/d;' null_gau
python /Volumes/My\ Passport/Bugbounty/scripts/working_script/misc/remove_duplicate_para.py null_gau
