package main

import (
	log "github.com/Sirupsen/logrus"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	version = "No version specified"

	pidArg     = kingpin.Arg("pid", "PID to watch").Required().Int32()
	listenFlag = kingpin.Flag("listen", "Address to listen on").Default(":8080").Short('l').String()
	debugFlag  = kingpin.Flag("debug", "Enable debug output").Short('d').Bool()
)

func init() {
	log.SetLevel(log.InfoLevel)

	// Parse CLI stuff
	kingpin.Version(version)
	kingpin.CommandLine.HelpFlag.Short('h')
	kingpin.CommandLine.VersionFlag.Short('v')
	kingpin.Parse()

	if *debugFlag {
		log.SetLevel(log.DebugLevel)
	}
}

func main() {
	p, err := NewPID(*pidArg)
	if err != nil {
		log.Fatalf("PID watcher error: %v", err)
	}

	// Launch pid watcher
	go p.Watch()

	// Launch http server
	s := NewServer(*listenFlag, p)

	log.Infof("Starting server on '%v'", *listenFlag)
	log.Fatal(s.Run())
}
