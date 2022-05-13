package main

import "f3s.tech/buddy/elastic"

func main() {
	installer := elastic.NewInstaller("/home/fit/temp")
	if err := installer.InstallElasticsearch(); err != nil {
		panic(err)
	}
	if err := installer.InstallKibana(); err != nil {
		panic(err)
	}
}
