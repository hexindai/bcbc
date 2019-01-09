BEGIN {

  if (bin=="" || len=="" || binfile =="") {
    print "bin, len and binfile are required"; exit 1
  }

  if ((len - length(bin)) < 0) {
    print "len should be longer than bin's length"; exit 1
  }

  Service = "curl \
  -s \
  -S \
  -G \
  -d \"cardNo=" random_bank(bin, len) "&_input_charset=utf-8&cardBinCheck=true\" \
  https://ccdcapi.alipay.com/validateAndCacheCardInfo.json"

  # random_bank(bin, len)
  result = ""

  while ((Service |& getline) > 0) {
      result = result $0
  }
  close(Service)

  # print response result
  if (debug) {
    print "[DEBUG] " result
  }

  dumbParseJson(result, obj)

  if (obj["validated"]) {
    printf "%s,%s,%s,%s\n", bin, obj["bank"], obj["cardType"], len >> binfile
  } else {
    print "BAD CARD BIN" > "/dev/stderr"; exit 1
  }
}

function random_bank(bin, len,    i, j) {
  j = len - length(bin)
  for (i=0; i<j; i++) {
    srand(systime() + i)
    bin = bin int(rand() * 10)
  }
  return bin
}

function dumbParseJson(json, obj) {

  match(json, /^{"cardType":"(.+)","bank":"([^,]+)".*"validated":([^,]+).+$/, obj)

  if (obj[3] == "true") {
    obj["validated"] = 1
    obj["cardType"] = obj[1]
    obj["bank"] = obj[2]
  } else {
    obj["validated"] = 0
    delete obj[3]
  }
}