package main

import (
	"errors"
	"log"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/gotk3/gotk3/gtk"
)

type Watcher struct {
	watcher *fsnotify.Watcher
	dir     string
	isWatch bool
	isDone  chan bool
}

func (w *Watcher) startWatch(tv *gtk.TextView) {
	if w.isWatch {
		if err := setText(tv, "Watcher allready running\n"); err != nil {
			log.Fatal(err)
		}
		return
	}
	w.isWatch = true
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	w.watcher = watcher
	err = watcher.Add(w.dir)
	if err != nil {
		log.Fatal(err)
	}
	go func(watcher *Watcher, tv *gtk.TextView) {
		for {
			select {
			case event, ok := <-w.watcher.Events:
				if !ok {
					return
				}
				if err := setText(tv, event.String()+"\n"); err != nil {
					log.Fatal(err)
				}
			case _, ok := <-w.watcher.Errors:
				if !ok {
					return
				}
			case <-w.isDone:
				break
			}
		}
	}(w, tv)
	if err := setText(tv, "Start watching\n"); err != nil {
		log.Fatal(err)
	}
}

func (w *Watcher) stopWatch(tv *gtk.TextView) {
	if !w.isWatch {
		if err := setText(tv, "Watcher is not running\n"); err != nil {
			log.Fatal(err)
		}
		return
	}
	if err := setText(tv, "Stop watching\n"); err != nil {
		log.Fatal(err)
	}
	w.isWatch = false
	w.isDone <- true
	w.watcher.Close()
}

func (w *Watcher) choiceDir(tv *gtk.TextView, win *gtk.Window) {
	fileChooserDlg, err := gtk.FileChooserNativeDialogNew("Open", win, gtk.FILE_CHOOSER_ACTION_SELECT_FOLDER, "_Open", "_Cancel")
	if err != nil {
		log.Fatal("Unable to create fileChooserDlg:", err)
	}
	response := fileChooserDlg.NativeDialog.Run()
	if gtk.ResponseType(response) == gtk.RESPONSE_ACCEPT {
		fileChooser := fileChooserDlg
		w.dir = fileChooser.GetFilename()
		if err := setText(tv, "Choice: "+w.dir+"\n"); err != nil {
			log.Fatal(err)
		}
	} else {
		cancelDlg := gtk.MessageDialogNew(win, gtk.DIALOG_MODAL, gtk.MESSAGE_INFO, gtk.BUTTONS_OK, "%s", "No file was selected")
		cancelDlg.Run()
		cancelDlg.Destroy()
	}
}

func setText(textView *gtk.TextView, text string) error {
	buffer, err := textView.GetBuffer()
	if err != nil {
		return errors.New("Unable to get buffer: " + err.Error())
	}
	start, end := buffer.GetBounds()
	oldText, err := buffer.GetText(start, end, true)
	if err != nil {
		return errors.New("Unable to get text: " + err.Error())
	}
	buffer.SetText(time.Now().Format("2006-01-02 15:04:05 ") + text + oldText)
	return nil
}
