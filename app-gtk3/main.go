package main

import (
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/gotk3/gotk3/gtk"
)

func main() {
	watcher := Watcher{
		watcher: new(fsnotify.Watcher),
		dir:     ".",             // default dir
		isWatch: false,           // check watchdog is running
		isDone:  make(chan bool), // break watchdog goroutine
	}

	gtk.Init(nil)

	b, err := gtk.BuilderNew()
	if err != nil {
		log.Fatal("Couldn't create the builder:", err)
	}

	err = b.AddFromString(gladeTemplate)
	if err != nil {
		log.Fatal("Couldn't add UI XML to builder:", err)

	}

	obj, err := b.GetObject("window_main")
	if err != nil {
		log.Fatal("Couldn't window object:", err)
	}

	win := obj.(*gtk.Window)
	win.Connect("destroy", func() {
		gtk.MainQuit()
	})

	obj, _ = b.GetObject("textView")
	textView := obj.(*gtk.TextView)

	obj, _ = b.GetObject("startButton")
	startButton := obj.(*gtk.Button)
	startButton.Connect("clicked", func() {
		watcher.startWatch(textView)
	})

	obj, _ = b.GetObject("stopButton")
	stopButton := obj.(*gtk.Button)
	stopButton.Connect("clicked", func() {
		watcher.stopWatch(textView)
	})

	obj, _ = b.GetObject("openButton")
	openButton := obj.(*gtk.Button)
	openButton.Connect("clicked", func() {
		watcher.choiceDir(textView, win)
	})

	win.ShowAll()
	gtk.Main()
}
