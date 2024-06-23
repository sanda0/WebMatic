package webmaticlib

import (
	"fmt"

	"github.com/go-rod/rod"
)

func ReadJson(content []byte) ([]Block, error) {

	return nil, nil
}

type ProjectRunner struct {
	Project Project
	Browser *rod.Browser
}

func NewProjectRunner(project Project, browser *rod.Browser) ProjectRunner {
	return ProjectRunner{
		Project: project,
		Browser: browser,
	}
}

func (r ProjectRunner) Start() error {
	fmt.Println(r.Project)
	var page *rod.Page
	for _, b := range r.Project.Blocks {
		switch b.Type {
		case "open":
			fmt.Println(b.Data["url"])
			page = r.Browser.MustPage(fmt.Sprintf("%v", b.Data["url"]))

		case "write":
			if page != nil {
				page.MustElement(fmt.Sprintf("%v", b.Data["target"])).MustInput(fmt.Sprintf("%v", b.Data["text"])).WaitLoad()
			}

		case "click":
			if page != nil {
				page.MustElement(fmt.Sprintf("%v", b.Data["target"])).MustClick().WaitLoad()
			}

		case "navigate":
			if page != nil {
				page.MustNavigate(fmt.Sprintf("%v", b.Data["url"])).WaitLoad()
			}
		}

	}
	return nil
}
