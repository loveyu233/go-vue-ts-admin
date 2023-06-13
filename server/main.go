package main

import (
	"log"
	"server/core"
)

func main() {
	log.SetFlags(log.Lshortfile | log.Ldate)
	core.InitCore()
}
