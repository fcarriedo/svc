package main

import (
	"bitbucket.org/kardianos/service"
	"encoding/xml"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

var log service.Logger

var svcDescriptor = `
<service>
  <id>nginx</id>
  <name>nginx</name>
  <desc>nginx awesomeness</desc>
  <exec>C:/Apps/nginx/nginx.exe</exec>
  <args>-p C:/Apps/nginx</args>
  <stopexec>C:/Apps/nginx/nginx.exe</stopexec>
  <stopargs>-p C:/Apps/nginx -s stop</stopargs>
</service>
`

type svc struct {
	XMLName  xml.Name `xml:"service"`
	Id       string   `xml:"id"`
	Name     string   `xml:"name"`
	Desc     string   `xml:"desc"`
	Exec     string   `xml:"exec"`
	Args     string   `xml:"args"`
	StopExec string   `xml:"stopexec"`
	StopArgs string   `xml:"stopargs"`
}

func main() {
	asvc, err := loadCfg()
	if err != nil {
		fmt.Println("Error while loading config: %s", err)
		return
	}

	s, err := service.NewService(asvc.Id, asvc.Name, asvc.Desc)
	log = s

	if err != nil {
		fmt.Printf("%s unable to start: %s", asvc.Name, err)
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
			fmt.Printf("Service \"%s\" installed.\n", asvc.Name)
		case "remove":
			err = s.Remove()
			if err != nil {
				fmt.Printf("Failed to remove: %s\n", err)
				return
			}
			fmt.Printf("Service \"%s\" removed.\n", asvc.Name)
		case "run":
			doStart(asvc)
		case "start":
			err = s.Start()
			if err != nil {
				fmt.Printf("Failed to start: %s\n", err)
				return
			}
			fmt.Printf("Service \"%s\" started.\n", asvc.Name)
		case "stop":
			err = s.Stop()
			if err != nil {
				fmt.Printf("Failed to stop: %s\n", err)
				return
			}
			fmt.Printf("Service \"%s\" stopped.\n", asvc.Name)
		}
		return
	}
	err = s.Run(func() error {
		go doStart(asvc)
		return nil
	}, func() error {
		if asvc.StopExec != "" {
			doStop(asvc)
		}
		return nil
	})
	if err != nil {
		s.Error(err.Error())
	}
}

func doStart(msvc svc) {
	cmd := exec.Command(msvc.Exec, strings.Split(msvc.Args, " ")...)
	err := cmd.Start()
	if err != nil {
		log.Error(err.Error())
	}
}

func doStop(msvc svc) {
	cmd := exec.Command(msvc.StopExec, strings.Split(msvc.StopArgs, " ")...)
	err := cmd.Start()
	if err != nil {
		log.Error(err.Error())
	}
}

func loadCfg() (svc, error) {
	var s svc
	err := xml.NewDecoder(strings.NewReader(svcDescriptor)).Decode(&s)
	if err != nil {
		log.Error(err.Error())
		return s, err
	}
	return s, nil
}
