package gui

import (
	"io/ioutil"
	"os"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var server_output_msg string

func Run_gui() {
	server_output_msg = "[START OUTPUT LOGS]"
	// start the fyne gui
	a := app.New()

	/*
		the layout will be this:
			main_page -> if the server is up and output log from the server
			edit      -> edit the html and add new pages
			quit      -> quit the program
	*/

	// window
	w := a.NewWindow("Simple Host")

	// window lbl
	hello := widget.NewLabel("Welcome to Simple Host! the server is running on http://127.0.0.1:8080 :)")
	logs := widget.NewLabel(server_output_msg)
	go Server_logs(logs)
	// content of window
	w.SetContent(container.NewVBox(
		hello,
		widget.NewButton("Edit!", func() {
			// load the html edit page
			edit := Edit(a)

			w.Hide()
			edit.Show()
			w.Show()
		}),
		widget.NewButton("Quit", func() {
			// Quit
			quit := Quit(a)

			w.Hide()
			quit.Show()
			w.Show()
		}),
		logs,
	))

	w.ShowAndRun()
	// if quit button is not pressed do it in the background
	os.Exit(0)

}

// util
func Server_logs(lbl *widget.Label) {
	for {
		// update it

		// wait 5 secods
		server_output_msg = "[WOW]"
		lbl.SetText(server_output_msg)
	}
}

/*
adding new buttons
1) create the button to activate it in the Run_gui function
2) create the function
3) create the gui app like Quit()
4) call it at the start of Run_gui
5) activate it when it's called
[INFO] Rember have a back button as we don't want it to be destoryed (as if they clicked X)
*/

// edit func
func Edit_file(a fyne.App, name string) fyne.Window {
	w := a.NewWindow("EDITOR")

	edit_box := widget.NewMultiLineEntry()
	edit_box.SetMinRowsVisible(10)

	if _, err := os.Stat(name); err == nil {
		read, _ := ioutil.ReadFile(name)
		edit_box.SetText(string(read[:]))
	} else {
		// file not real, create it and add bioler plate
		os.Create(name)
		edit_box.SetText(`
<!DOCTYPE html>
<html>
	<head>
		<title>EDIT ME</title>
	</head>
	<body>
		<p>EDIT ME</p>
	</body>
</html>
		
		`)
	}

	w.SetContent(container.NewVBox(
		edit_box,
		widget.NewButton("Save", func() {
			// get content
			name_con := edit_box.Text
			// save content to file
			os.WriteFile(name, []byte(name_con), 0644)
		}),
	))

	return w
}

func Edit(a fyne.App) fyne.Window {
	w := a.NewWindow("Edit")
	up := widget.NewLabel("Enter a file name to edit, if it doens't exisit we will make it :)")
	text_box := widget.NewEntry()
	text_box.SetPlaceHolder("File Name")
	w.SetContent(container.NewVBox(
		up,
		text_box,
		widget.NewButton("Edit", func() {
			// get the file name
			file_name := text_box.Text
			up.SetText("Opening: " + file_name + "...")
			if strings.Contains(file_name, ".html") {
				// html file
				file_name = "templates/html/" + file_name
				app := Edit_file(a, file_name)

				w.Hide()
				app.Show()
				w.Show()

			} else if strings.Contains(file_name, ".css") {
				// css file
				file_name = "templates/css/" + file_name
				app := Edit_file(a, file_name)

				w.Hide()
				app.Show()
				w.Show()
			} else {
				// not editable
				up.SetText("[!!] The file " + file_name + " Cannot be opend as it's not a .html file or a .css file\n. Try a diffrent name")
			}

		}),
	))

	return w

}

// quit func setup
func Quit(a fyne.App) fyne.Window {
	w := a.NewWindow("Quit")
	w.SetContent(container.NewVBox(
		widget.NewLabel("Are you sure?"),
		widget.NewButton("Yes", func() {
			os.Exit(0)
		}),
		widget.NewButton("No", func() {
			w.Hide()
		}),
	))

	return w
}
