// Generates a globally unique id
// https://github.com/rs/xid#usage

package main

import "github.com/rs/xid"

func main() {
	guid := xid.New()

	println(guid.String())

}
