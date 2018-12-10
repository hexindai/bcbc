# sort bin.csv

BEGIN {
    FS=","
}

NR == 1 {
    print > to; next
}

NR > 1 {
    lines[$1 ":" $4]=$0
}

END {
    PROCINFO["sorted_in"] = "sort_bin"
    for (l in lines) {
        print lines[l] > to
    }
}

function sort_bin(k1, v1, k2, v2,    a1, a2) {
    
    split(k1, a1, ":")
    split(k2, a2, ":")

    k1 = a1[1] + 0 ""
    k2 = a2[1] + 0 ""

    if (k1 < k2) {
        return -1
    }
    if (k1 == k2) {
        return (a1[2] - a2[2])
    }
    return 1
}
