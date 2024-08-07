package main

import (
	"context"
	"fmt"
	"log"

	"github.com/sanda0/webmatic/webmaticlib"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Response struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
}

// App struct
type App struct {
	ctx context.Context
	DB  *gorm.DB
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx

	var err error
	a.DB, err = gorm.Open(sqlite.Open("data.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	a.DB.AutoMigrate(&webmaticlib.Project{})
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) FileUpload(fileName string) int {
	return 100
}

func (a *App) SaveMatic(name string, autor string) Response {
	fileName, err := webmaticlib.SaveMatic(a.DB, name, autor)
	if err != nil {
		log.Println(err.Error())
		return Response{
			Status: 500,
			Data:   err.Error(),
		}
	}
	return Response{Status: 200, Data: fileName}
}

func (a *App) GetAllMatics() Response {
	matics, err := webmaticlib.GetAllMatics(a.DB)
	if err != nil {
		log.Println(err.Error())
		return Response{
			Status: 500,
			Data:   err.Error(),
		}
	}
	return Response{Status: 200, Data: matics}
}

func (a *App) GetMaticById(id uint) Response {
	matic, err := webmaticlib.GetMaticById(a.DB, id)
	if err != nil {
		return Response{
			Status: 500,
			Data:   err.Error(),
		}
	}
	return Response{Status: 200, Data: matic}
}

func (a *App) SaveXML(id uint, data string) Response {
	err := webmaticlib.SaveXML(a.DB, id, data)
	if err != nil {
		return Response{
			Status: 500,
			Data:   err.Error(),
		}
	}
	return Response{Status: 200, Data: "xml saved"}
}

func (a *App) RunMatic(jsonStr string) {

	// fmt.Println(jsonStr)
	webmaticlib.RunMatic(jsonStr)

}
