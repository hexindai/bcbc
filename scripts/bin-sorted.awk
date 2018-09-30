# sort bin.csv

BEGIN {
    FS=","
}

NR == 1 {
    print > FILENAME; next
}

NR > 1 {
    lines[$1 ":" $4]=$0
}

END {
    PROCINFO["sorted_in"] = "@ind_str_asc"
    for (l in lines) {
        print lines[l] > FILENAME
    }
}