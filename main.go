package main

import (
	"fmt"
	"log"
	"os"
	"time"

	c "digger/clockUtils"
	dirUtils "digger/dirUtils"
	fileUtils "digger/fileUtils"
	"digger/logUtils"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/kardianos/service"
)

/*
func updateTime(clock *widget.Label) {
	formattedTime := time.Now().Format("Time: 03:04:05")
	clock.SetText(formattedTime)
}
*/

type MyDirList struct {
	selected int
}

func selectFirstItem(listWidget *widget.List) {
	if listWidget.Length() > 0 {
		listWidget.Select(0)
	}
}

func main() {
	a := app.New()
	setup()
	a.Settings().SetTheme(theme.DarkTheme())
	w := a.NewWindow("Digger")
	w.Resize(fyne.NewSize(800, 800))

	var myDirList MyDirList

	currDir := binding.NewString()
	visibleDirStrings := &[]string{}
	actualDirStrings := &[]string{}
	dirList := binding.BindStringList(
		visibleDirStrings,
	)

	adsBinding := binding.BindStringList(
		actualDirStrings,
	)

	currentDirectoryLabel := widget.NewLabelWithData(currDir)
	clock := widget.NewLabel("")

	dirListWidget := widget.NewListWithData(dirList,
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(i binding.DataItem, o fyne.CanvasObject) {
			o.(*widget.Label).Bind(i.(binding.String))
		})

	dirListWidget.OnSelected = func(id int) {
		myDirList.selected = id
	}

	cDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	currDir.Set(cDir)
	dirList.Set(nil)
	adsBinding.Set(nil)
	files := dirUtils.GetDirectoryList(cDir)
	for _, file := range files {
		isFile := fileUtils.IsFile(dirUtils.GetFullPath(file.Name()))
		marker := ""
		if isFile {
			marker = "FILE"
		}
		perm := fileUtils.GetPermissions(dirUtils.GetFullPath(file.Name()))
		dirList.Append(fmt.Sprintf("%s - %s - %s\n", dirUtils.GetFullPath(file.Name()), marker, perm))
		adsBinding.Append(fmt.Sprintf("%s", dirUtils.GetFullPath(file.Name())))
	}
	selectFirstItem(dirListWidget)

	openfileWithVSCode := widget.NewButton("Open with code", func() {
		file, err := adsBinding.GetValue(myDirList.selected)
		if logUtils.LogFatalError(err) {
			return
		}
		fileUtils.OpenFileWithProgram("code", file)
	})

	upDirectory := widget.NewButton("Up dir", func() {
		//current, err := adsBinding.GetValue(myDirList.selected)
		dirUtils.WalkUpDirectory()
		if err != nil {
			log.Fatal(err)
		}
		cDir, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
		currDir.Set(cDir)
		dirList.Set(nil)
		adsBinding.Set(nil)
		files := dirUtils.GetDirectoryList(cDir)
		for _, file := range files {
			isFile := fileUtils.IsFile(dirUtils.GetFullPath(file.Name()))
			marker := ""
			if isFile {
				marker = "FILE"
			}
			perm := fileUtils.GetPermissions(dirUtils.GetFullPath(file.Name()))
			dirList.Append(fmt.Sprintf("%s - %s - %s\n", dirUtils.GetFullPath(file.Name()), marker, perm))
			adsBinding.Append(fmt.Sprintf("%s", dirUtils.GetFullPath(file.Name())))
		}
		selectFirstItem(dirListWidget)
	})

	downDirectory := widget.NewButton("Down dir", func() {
		current, err := adsBinding.GetValue(myDirList.selected)
		dirUtils.WalkDownDirectory(current)
		if err != nil {
			log.Fatal(err)
		}
		cDir, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
		currDir.Set(cDir)
		dirList.Set(nil)
		adsBinding.Set(nil)
		files := dirUtils.GetDirectoryList(cDir)
		for _, file := range files {
			isFile := fileUtils.IsFile(dirUtils.GetFullPath(file.Name()))
			marker := ""
			if isFile {
				marker = "FILE"
			}
			perm := fileUtils.GetPermissions(dirUtils.GetFullPath(file.Name()))
			dirList.Append(fmt.Sprintf("%s - %s - %s\n", dirUtils.GetFullPath(file.Name()), marker, perm))
			adsBinding.Append(fmt.Sprintf("%s", dirUtils.GetFullPath(file.Name())))
		}
		selectFirstItem(dirListWidget)
	})

	c.UpdateTime(clock)
	header := container.NewHBox(clock, openfileWithVSCode, currentDirectoryLabel, upDirectory, downDirectory)
	content := container.NewBorder(header, nil, nil, nil, dirListWidget)
	w.SetContent(content)

	go func() {
		for range time.Tick(time.Second) {
			c.UpdateTime(clock)
		}
	}()

	go func() {
		for range time.Tick(time.Millisecond) {
			//current, _ := currDir.Get()
			current, _ := currDir.Get()
			currDir.Set(current)
		}
	}()

	w.ShowAndRun()

}

var logger service.Logger

type program struct{}

func (p *program) Start(s service.Service) error {
	go p.run()
	return nil
}
func (p *program) run() {
	return
}
func (p *program) Stop(s service.Service) error {
	return nil
}

func setup() {
	svcConfig := &service.Config{
		Name:        "Digger",
		DisplayName: "Digger",
		Description: "Directory explorer",
	}
	digger := &program{}
	s, err := service.New(digger, svcConfig)

	logger, err = s.Logger(nil)
	if err != nil {
		log.Fatal(err)
	}

	/*err = s.Install()
	if err != nil {
		fmt.Printf("installed %s", err.Error())
	}*/

	err = s.Start()
	if err != nil {
		logger.Error(err)
	}
}
