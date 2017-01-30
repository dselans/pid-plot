pid-plot
========
Basic/poor-man's utility for watching a given pid, graphing its resource usage and making the graphs available via the built-in webserver.

Note: This is a result of being too lazy to instrument datadog, go-metrics, new relic, etc., when all I wanted is to see is if memory consumption on a given process increases. Ie. poor man's mem leak detection.

## Usage
```
$ ./pid-plot -p <pid>
INFO[0000] Server listening on :8080
```

View metrics on http://localhost:8080/

You can drill down to a specific timeslice by passing a time.Duration compatible duration string (1h30m, 1ms and so on) as the path param: http://localhost:8080/view/{durationParam}.

PID resources are collected every 1s. Once the server is stopped, metrics collection stops.

Here's a screenshot too!

![screenshot](/img/screenshot.png?raw=true)

That's all folks!
