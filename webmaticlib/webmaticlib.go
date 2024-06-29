package webmaticlib

import (
	"fmt"

	"github.com/go-rod/rod"
)

func ReadJson(content []byte) ([]Block, error) {

	return nil, nil
}

type ProjectRunner struct {
	BlockContainer BlockContainer
	Browser        *rod.Browser
}

func NewProjectRunner(blockContainer BlockContainer, browser *rod.Browser) ProjectRunner {
	return ProjectRunner{
		BlockContainer: blockContainer,
		Browser:        browser,
	}
}

func (r ProjectRunner) Start() error {

	var page *rod.Page
	for _, b := range r.BlockContainer.Blocks {
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
				fmt.Printf("this is url : %v", b.Data["url"])
				page.MustNavigate(fmt.Sprintf("%v", b.Data["url"])).WaitLoad()
			}
		}

	}
	return nil
}
