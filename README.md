# gpsdtestclient
Client that dumps data from gpsd

Based on the go_gpsd library https://github.com/stratoberry/go-gpsd.

gpsd must be running and listening on port `localhost:2947`.

Currently prints "TPV" (time-position-velocity) and SKY (Sattelites) events.

See https://gpsd.gitlab.io/gpsd/gpsd_json.html for more info.

