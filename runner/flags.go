package main

import "flag"

func InitFlags() {
	flag.Parse()
}

var (
	DockerContext = flag.String("dockerContext", ".", "Relative path to the context of the docker build. This directory should contain your Dockerfile")
	Dockerfile    = flag.String("dockerfile", "Dockerfile", "Relative path to the directory containing your Dockerfile")
)
