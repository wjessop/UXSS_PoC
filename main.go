package main

import (
	"flag"
)

var (
	addr           = flag.String("addr", ":9090", "http service address")
	non_target_url = flag.String("non_target_url", "https://www.google.com/images/srpr/logo11w.png", "A URL that we can request, necessary part of the exploit but it doesn't matter what it is.")
	target_url     = flag.String("target_url", "", "URL on target domain that we want to run our code on.")
	use_sleep      = flag.String("use_sleep", "yes", "Use sleep method of exploitation instead of alert()")
)

func main() {
	flag.Parse()
	serveHTTP()
}
