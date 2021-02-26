#!/usr/bin/env bash

# http://www.ruanyifeng.com/blog/2018/11/awk.html
# 得到如F结果:域名的出现次数,域名

awk -F / '{print $3}' file1.txt | sort -r | uniq -c
echo "-----"

cat file1.txt|awk 'NR==1'|awk '{print $NF}'

head -1 file2|awk '{print $(NF-1)}'


sed -n '/2020-12-15 15:15:15/,/2020-12-17 17:17:17/p' time.txt
