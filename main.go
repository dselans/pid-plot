package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
	"github.com/wcharczuk/go-chart"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	version = "No version specified"

	pidArg = kingpin.Arg("pid", "PID to watch").Required().Int32()
	debug  = kingpin.Flag("debug", "Enable debug output").Short('d').Bool()
)

type view struct {
	Slice *time.Duration
	Error error
}

func init() {
	log.SetLevel(log.InfoLevel)

	// Parse CLI stuff
	kingpin.Version(version)
	kingpin.CommandLine.HelpFlag.Short('h')
	kingpin.CommandLine.VersionFlag.Short('v')
	kingpin.Parse()

	if *debug {
		log.SetLevel(log.DebugLevel)
	}
}

// Redirects to 30 min view
func BaseHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/view/30m", http.StatusPermanentRedirect)
}

// Displays specific time slice
func SliceHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	slice := vars["slice"]

	t, err := template.ParseFiles("index.template")
	if err != nil {
		log.Fatalf("Unable to find required template: %v", err)
	}

	d, err := time.ParseDuration(slice)
	if err != nil {
		t.Execute(w, &view{Slice: nil, Error: fmt.Errorf("Unable to parse duration '%v': %v", slice, err)})
		return
	}

	t.Execute(w, &view{Slice: &d})
}

// Generates an image
func ImageHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	metricType := vars["type"]

	switch metricType {
	case "mem":
		metricType = "Memory Usage (KB)"
	case "cpu":
		metricType = "CPU %"
	}

	graph := chart.Chart{
		XAxis: chart.XAxis{
			Name:      "Time Period",
			NameStyle: chart.StyleShow(),
			Style:     chart.StyleShow(),
		},
		YAxis: chart.YAxis{
			Name:      metricType,
			NameStyle: chart.StyleShow(),
			Style:     chart.StyleShow(),
		},
		Series: []chart.Series{
			chart.ContinuousSeries{
				Style: chart.Style{
					Show:        true,
					StrokeColor: chart.GetDefaultColor(0).WithAlpha(64),
					FillColor:   chart.GetDefaultColor(0).WithAlpha(64),
				},
				XValues: []float64{1.0, 2.0, 3.0, 4.0, 5.0},
				YValues: []float64{1.0, 2.0, 3.0, 4.0, 5.0},
			},
		},
	}

	w.Header().Set("Content-Type", "image/png")
	graph.Render(chart.PNG, w)
}

func main() {
	p, err := NewPID(*pidArg)
	if err != nil {
		log.Fatalf("PID watcher error: %v", err)
	}

	go p.Watch()

	time.Sleep(5 * time.Second)

	stats := p.GetStats()
	log.Infof("Our stats here: %v", stats)

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", BaseHandler)
	router.HandleFunc("/view/{slice}/{type}_metrics.png", ImageHandler)
	router.HandleFunc("/view/{slice}", SliceHandler)
	log.Fatal(http.ListenAndServe(":8080", router))
}
