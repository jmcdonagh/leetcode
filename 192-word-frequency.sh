#!/usr/bin/env bash

gawk -F' ' 'BEGIN {PROCINFO["sorted_in"]="@val_num_desc"} {for(i=1;i<=NF;i++) a[$i]++} END {for(k in a) print k,a[k]}' words.txt
