package main

import (
	"./zproto"
	"fmt"
)

func main() {
	fmt.Printf("===begin \n")
	pt := zproto.ReadFile("test.zp")

	zproto.MatchMessage(pt)

	zproto.MatchTest(" #ss \n#bb\n")

}
