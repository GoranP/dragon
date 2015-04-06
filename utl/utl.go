package utl

import (
	"log"
	"time"
)

func TimeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.SetFlags(log.Ldate | log.Lmicroseconds)
	log.Printf("%s \t %s \n", name, elapsed)
}
