package main

import (
	"bitbucket.org/kardianos/service"
	"encoding/xml"
	"fmt"
	"os"
	"os/exec"
	"io/ioutil"
	"strings"
)

const confFile = "svc.xml"

var log service.Logger

type svcCmd struct {
	XMLName     xml.Name `xml:"service"`
	Name        string   `xml:"id"`
	DisplayName string   `xml:"name"`
	Desc        string   `xml:"description"`
	Exec        string   `xml:"executable"`
	Args        string
	Stopargs    string
}

func main() {
	cmd, err := loadConfig()
	if err != nil {
    fmt.Printf("Error loading conf file: %s\n", err)
		return
	}

	s, err := service.NewService(cmd.Name, cmd.DisplayName, cmd.Desc)
	log = s

	if err != nil {
		fmt.Printf("%s unable to start: %s", cmd.DisplayName, err)
		return
	}

	if len(os.Args) > 1 {
		var err error
		verb := os.Args[1]
		switch verb {
		case "install":
			err = s.Install()
			if err != nil {
				fmt.Printf("Failed to install: %s\n", err)
				return
			}
			fmt.Printf("Service \"%s\" installed.\n", cmd.DisplayName)
		case "remove":
			err = s.Remove()
			if err != nil {
				fmt.Printf("Failed to remove: %s\n", err)
				return
			}
			fmt.Printf("Service \"%s\" removed.\n", cmd.DisplayName)
		case "run":
			start(cmd)
		case "start":
			err = s.Start()
			if err != nil {
				fmt.Printf("Failed to start: %s\n", err)
				return
			}
			fmt.Printf("Service \"%s\" started.\n", cmd.DisplayName)
		case "stop":
			err = s.Stop()
			if err != nil {
				fmt.Printf("Failed to stop: %s\n", err)
				return
			}
			fmt.Printf("Service \"%s\" stopped.\n", cmd.DisplayName)
		}
		return
	}

	err = s.Run(func() error {
		// start
		go start(cmd)
		return nil
	}, func() error {
		go stop(cmd)
		return nil
	})

	if err != nil {
		s.Error(err.Error())
	}
}

func loadConfig() (svcCmd, error) {
	var cfg svcCmd
	xmlBytes, err := ioutil.ReadFile(confFile)
	if err != nil {
		return cfg, err
	}

	err = xml.Unmarshal(xmlBytes, &cfg)
	if err != nil {
		return cfg, err
	}

	return cfg, nil
}

func start(cmd svcCmd) {
	execCmd(cmd.Exec, strings.Split(cmd.Args, " "))
}

func stop(cmd svcCmd) {
	execCmd(cmd.Exec, strings.Split(cmd.Stopargs, " "))
}

func execCmd(svcExe string, args []string) {
	cmd := exec.Command(svcExe, args...)
	err := cmd.Run()
	if err != nil {
		log.Error(err.Error())
	}
}
