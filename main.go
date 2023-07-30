package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/sanda0/webmatic/webmaticlib"
)

func main() {

	jsonFile, err := os.Open("input.json")
	if err != nil {
		log.Fatal(err)
	}

	defer jsonFile.Close()

	fmt.Println(jsonFile)

	project := webmaticlib.Project{}
	byteVal, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteVal, &project)

	fmt.Println(project)

	l := launcher.New().Headless(project.Headless)
	u := l.MustLaunch()
	browser := rod.New().ControlURL(u).MustConnect()
	projectRunner := webmaticlib.ProjectRunner{

		Project: project,
		Browser: browser,
	}

	projectRunner.Start()

	time.Sleep(time.Hour)

}
