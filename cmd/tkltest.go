package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"github.com/konveyor/tackle2-addon/command"
)

// tkltest application analyzer.
type Tkltest struct {
	application *api.Application
	*Data
}

// Run tkltest add on
func (r *Tkltest) Run() error {

	output := r.output()
	cmd := command.Command{Path: "/usr/bin/rm"}
	cmd.Options.Add("-rf", output)
	err = cmd.Run()
	if err != nil {
		return
	}
	err = os.MkdirAll(output, 0777)
	if err != nil {
		err = liberr.Wrap(
			err,
			"path",
			output)
		return
	}
	addon.Activity("[TCA] created: %s.", output)
	cmd = command.Command{Path: "/opt/tca"}
	cmd.Options, err = r.options()
	if err != nil {
		return
	}
	err = cmd.Run()
	if err != nil {
		r.reportLog()
	}

	return
}

// output returns output directory.
func (r *Tkltest) output() string {
	return pathlib.Join(
		r.application.Bucket,
		r.Output)
}


// reportLog reports the log content.
func (r *Tkltest) reportLog() {
	logPath := path.Join(
		HomeDir,
		".mta",
		"log",
		"mta.log")
	f, err := os.Open(logPath)
	if err != nil {
		return
	}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		addon.Activity(">> %s\n", scanner.Text())
	}
	_ = f.Close()
}

func (r *Tkltest) options() (options command.Options, err error) {
	options = command.Options{
		"-input_json",
		"-output_json"
	}
	return
}
