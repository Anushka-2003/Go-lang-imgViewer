package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	// "fyne.io/fyne/v2/container"
	// "fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2/canvas"
	// "fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/container"
)

func main() {
	a := app.New()
	w := a.NewWindow("Pictures")
	w.Resize(fyne.NewSize(800, 600))
	root_src := "C:\\Users\\Admin\\Pictures\\Screenshots";
	files, err := ioutil.ReadDir(root_src);

	if err != nil{
		log.Fatal(err)
	}

	var picsArr []string;
	for _, file := range files{
		fmt.Println(file.Name(), file.IsDir())
		if !file.IsDir(){
			extension := strings.Split(file.Name(), ".")[1];
			if extension == "png" || extension == "jpeg"{
				picsArr = append(picsArr, root_src+"\\"+file.Name())
			}
		}
	}

	// image := canvas.NewImageFromFile(picsArr[0])
	tabs := container.NewAppTabs(
		
		container.NewTabItem("Image", canvas.NewImageFromFile(picsArr[0])),
	)
	for i:= 1; i < len(picsArr); i++{
		image := canvas.NewImageFromFile(picsArr[i]);
		image.FillMode = canvas.ImageFillOriginal
		tabs.Append(container.NewTabItem("Image", image))
	}

	tabs.SetTabLocation((container.TabLocationLeading))
	w.SetContent(tabs)
	w.ShowAndRun()
}