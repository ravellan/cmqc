package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/ibm-messaging/mq-golang/v5/ibmmq"
)

var (
	flagPrefix  = flag.String("prefix", "RC", "Prefix")
	flagCode    = flag.Int("code", 0, "Code")
	flagStrOnly = flag.Bool("string", false, "Output resulting string only")
)

func main() {
	flag.Parse()
	prefix := strings.ToUpper(*flagPrefix)
	mqiStr := ibmmq.MQItoString(prefix, *flagCode)
	if mqiStr == "" {
		mqiStr = "no matching string for prefix and code"
	}

	if *flagStrOnly {
		fmt.Println(mqiStr)
	} else {
		fmt.Printf("%s %d: %s\n",
			prefix,
			*flagCode,
			mqiStr,
		)
	}
}
