# spark-monitor
lightweight daemon for monitoring and notification about spark problems


## Build (optionally, needed to update ca-certificates)

``` bash
./build.sh 0.1.1
```

## Run

``` bash
docker run -e "APP_COUNT=3" \
           -e "DRIVER_ADDRESS=spark:4040" \
           -e "PERIOD=60" \
           -e "LONG_EXECUTION=3600" \
           -e "PUBLISHER_ADDRESS=slack-publisher:8000" \
           -d -p cortwave/spark-monitor:0.1.1
```

* DRIVER_ADDRESS - address (host + port) of running spark driver
* APP_COUNT - apss count which should be run on spark driver
* PERIOD - period in sec to check spark state
* LONG_EXECUTION - stage execution time in seconds enough for notification
* PUBLISHER_ADDRESS - address of running publisher [slack-publisher](https://github.com/cortwave/slack-publisher) e.g.)
