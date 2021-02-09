package main

import (
	"fmt"

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
	}

	gps.AddFilter("SKY", skyfilter)

	done := gps.Watch()
	<-done
}
