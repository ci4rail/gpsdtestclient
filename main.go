package main

import (
	"fmt"
	"os"
	"time"

	"github.com/stratoberry/go-gpsd"
)

var Version = "dev"

func main() {
	var gps *gpsd.Session
	var err error

	fmt.Println("Version", Version)

	if gps, err = gpsd.Dial(gpsd.DefaultAddress); err != nil {
		panic(fmt.Sprintf("Failed to connect to GPSD: %s", err))
	}

	gps.AddFilter("TPV", func(r interface{}) {
		tpv := r.(*gpsd.TPVReport)
		fmt.Printf("TPV: Mode=%v Time=%v Lat=%.5f Lon=%.5f Alt=%.1f Speed=%.3f\n", tpv.Mode, tpv.Time, tpv.Lat, tpv.Lon, tpv.Alt, tpv.Speed)
	})

	skyfilter := func(r interface{}) {
		sky := r.(*gpsd.SKYReport)

		fmt.Println("SKY", len(sky.Satellites), "satellites")

		for _, s := range sky.Satellites {
			dumpSat(s)
		}

	}

	gps.AddFilter("SKY", skyfilter)

	done := gps.Watch()

	if len(os.Args) > 1 {
		timeout, err := time.ParseDuration(os.Args[1])
		if err != nil {
			panic(fmt.Sprintf("Cannot parse time duration %s", err))
		}
		fmt.Printf("Exit after %v seconds\n", timeout)

		go func() {
			time.Sleep(timeout)
			done <- true
		}()
	}
	<-done
}

func dumpSat(s gpsd.Satellite) {
	fmt.Printf("  PRN=%v, ss=%.3f\n", s.PRN, s.Ss)
}
