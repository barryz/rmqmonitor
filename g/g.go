package g

import (
	"log"
	"runtime"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU() * 2)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}
