package webmaticlib

type Block struct {
	Type string                 `json:"type"`
	Data map[string]interface{} `json:"data"`
}

type BlockContainer struct {
	Headless bool    `json:"headless"`
	Blocks   []Block `json:"blocks"`
}
