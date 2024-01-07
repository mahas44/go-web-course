package main

import (
	fs "serving-file/file-server"
)

var tobyPic = "toby.jpg"

func main() {
	// IOCopy()
	// ServeContent()
	// ServeFile()

	// fs.FileServer()
	fs.HttpStripPrefix()
}
