package main

import (
	"flag"
)

var (
	addr           = flag.String("addr", ":9090", "http service address")
	non_target_url = flag.String("non_target_url", "https://www.google.com/images/srpr/logo11w.png", "A URL that we can request, necessary part of the exploit but it doesn't matter what it is.")
	protected_url  = flag.String("protected_url", "", "URL on target domain that we want to run our code on.")
)

func main() {
	flag.Parse()
	serveHTTP()
}
