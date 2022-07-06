mkdir wayback_process
cp wayback wayback_process/
cd wayback_process/
grep -Er "\?\w+=|\?\w+\&\w+=" wayback|tee para_wayback
sed -i -e 's/wayback://g' para_wayback
nullify_para para_wayback|tee null_wayback
sed -i '' '/.woff/d;/.png/d;/.gif/d;/.jpg/d;/.jpeg/d;/.css/d;/.js/d;' null_wayback
python /Volumes/My\ Passport/Bugbounty/scripts/working_script/misc/remove_duplicate_para.py null_wayback
