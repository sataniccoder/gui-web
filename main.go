package main

import (
	gui "gui_host/mod/gui"
	serv "gui_host/mod/serv"
)

func main() {

	// start the server
	go serv.Serv()
	// start the gui
	gui.Run_gui()
}
