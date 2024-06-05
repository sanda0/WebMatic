package webmaticlib

type Block struct {
	Type string                 `json:"type"`
	Data map[string]interface{} `json:"data"`
}

type Project struct {
	Name        string  `json:"name"`
	Author      string  `json:"author"`
	Description string  `json:"description"`
	CreateAt    string  `json:"createAt"`
	Headless    bool    `json:"headless"`
	Blocks      []Block `json:"blocks"`
}
