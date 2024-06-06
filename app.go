package main

import (
	"context"
	"fmt"
	"log"

	"github.com/sanda0/webmatic/webmaticlib"
)

type Response struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
}

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) FileUpload(fileName string) int {
	return 100
}

func (a *App) SaveMatic(name string, autor string) Response {
	fileName, err := webmaticlib.SaveMatic(name, autor)
	if err != nil {
		log.Println(err.Error())
		return Response{
			Status: 500,
			Data:   err.Error(),
		}
	}
	return Response{Status: 200, Data: fileName}
}
