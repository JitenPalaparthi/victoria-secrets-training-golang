package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/golang/glog"
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage: example -stderrthreshold=[INFO|WARNING|FATAL] -log_dir=[string]\n")
	flag.PrintDefaults()
	os.Exit(2)
}

func init() {
	flag.Usage = usage
	// NOTE: This next line is key you have to call flag.Parse() for the command line
	// options or "flags" that are defined in the glog module to be picked up.
	flag.Parse()
}

func main() {
	//flag.Parse()

	//glog.V(2).Infoln("Welcome to Vitoria Secrets & Co")
	glog.Infoln("Welcome to Victoria Secrets & Co")
	glog.Warningln("No warning but sample warning")
	glog.Errorln("No error but error log")
	//glog.Fatalln("Not a fatal but to test only") // commented becase it crashes this application

}
