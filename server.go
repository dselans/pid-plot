package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
	"github.com/wcharczuk/go-chart"
)

type Server struct {
	pid           *PID
	listenAddress string
}

type view struct {
	Slice *time.Duration
	Error error
}

func NewServer(address string, p *PID) *Server {
	return &Server{
		pid:           p,
		listenAddress: address,
	}
}

// Redirects to 30 min view
func (s *Server) BaseHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/view/30m", http.StatusPermanentRedirect)
}

// Displays specific time slice
func (s *Server) SliceHandler(w http.ResponseWriter, r *http.Request) {
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
func (s *Server) ImageHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	metricType := vars["type"]
	slice := vars["slice"]

	d, err := time.ParseDuration(slice)
	if err != nil {
		log.Errorf("Invalid duration '%v': %v", slice, err)
		return
	}

	switch metricType {
	case "mem":
		metricType = "Memory Usage (KB)"
	case "cpu":
		metricType = "CPU %"
	}

	dataset, err := GenerateXYDataset(d, s.pid.GetStats())
	if err != nil {
		log.Errorf("Unable to generate dataset for graph: %v", err)
		return
	}

	graph := chart.Chart{
		XAxis: chart.XAxis{
			Style:          chart.StyleShow(),
			ValueFormatter: chart.TimeMinuteValueFormatter,
		},
		YAxis: chart.YAxis{
			Name:      metricType,
			NameStyle: chart.StyleShow(),
			Style:     chart.StyleShow(),
		},
		Series: []chart.Series{
			chart.TimeSeries{
				Style: chart.Style{
					Show:        true,
					StrokeColor: chart.GetDefaultColor(0).WithAlpha(64),
					FillColor:   chart.GetDefaultColor(0).WithAlpha(64),
				},
				XValues: dataset.XValues,
				YValues: dataset.YValues,
			},
		},
	}

	w.Header().Set("Content-Type", "image/png")
	graph.Render(chart.PNG, w)
}

func (s *Server) Run() error {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", s.BaseHandler)
	router.HandleFunc("/view/{slice}/{type}_metrics.png", s.ImageHandler)
	router.HandleFunc("/view/{slice}", s.SliceHandler)

	return http.ListenAndServe(s.listenAddress, router)
}
