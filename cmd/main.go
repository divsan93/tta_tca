package main

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/FabianWe/goslugify"
	"github.com/konveyor/tackle2-addon/repository"
	"github.com/konveyor/tackle2-addon/ssh"
	hub "github.com/konveyor/tackle2-hub/addon"
)

var (
	// hub integration.
	addon = hub.Addon
	// HomeDir directory.
	HomeDir   = ""
	BinDir    = ""
	SourceDir = ""
	AppDir    = ""
	Dir       = ""
)

type SoftError = hub.SoftError

// addon data passed in secret
type Data struct {
	// Input directory within application bucket.
	Input string `json:"input" binding:"required"`
	// Output directory within application bucket.
	Output string `json:"output" binding:"required"`
	// Mode options.
	//Mode Mode `json:"mode"`
}


// main
func main() {

	addon.Run(func() error {

		Dir, _ = os.Getwd()
		HomeDir, _ = os.UserHomeDir()
		SourceDir = path.Join(Dir, "source")
		BinDir = path.Join(Dir, "dependencies")

		// Get the addon data associated with the task.
		d := &Data{}
		if err := addon.DataWith(d); err != nil {
			return &SoftError{Reason: err.Error()}
		}

		// Setup tkltest
		tkltest := Tkltest{}

		// Fetch application.
		addon.Activity("Fetching application.")
		application, err := addon.Task.Application()
		if err != nil {
			return err
		}
		// NOTE: We slugify to handle application names with spaces intelligently
		tkltest.appName = goslugify.GenerateSlug(application.Name)

		// SSH
		agent := ssh.Agent{}
		if err = agent.Start(); err != nil {
			return err
		}

		addon.Total(5)
		if application.Repository == nil {
			return &SoftError{Reason: "Application repository not defined."}
		}
		SourceDir = path.Join(
			Dir,
			strings.Split(
				path.Base(
					application.Repository.URL),
				".")[0])
		AppDir = path.Join(SourceDir, application.Repository.Path)
		repo, err := repository.New(SourceDir, application)
		if err != nil {
			return err
		}
		err = repo.Fetch()
		if err != nil {
			return err
		}
		addon.Increment()

		// Run tkltest.
		if err = tkltest.Run(); err != nil {
			return &SoftError{Reason: err.Error()}
		}
		addon.Increment()

		return
	})
}
