package main

import (
	"log"
	"path/filepath"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

type MyMainWindow struct {
	*walk.MainWindow
	textEdit *walk.TextEdit
	dirPath  string
	watcher  *fsnotify.Watcher
	isWatch  bool
	isDone   chan bool
}

func main() {
	mw := new(MyMainWindow)
	*mw = MyMainWindow{
		dirPath: ".",
		isDone:  make(chan bool),
	}

	if _, err := (MainWindow{
		AssignTo: &mw.MainWindow,
		Title:    "Folder watchdog",
		MinSize:  Size{Width: 300, Height: 200},
		Size:     Size{Width: 600, Height: 200},
		Layout:   VBox{},
		Children: []Widget{
			TextEdit{
				AssignTo: &mw.textEdit,
				ReadOnly: true,
			},
			Composite{
				Layout: HBox{MarginsZero: true},
				Children: []Widget{
					PushButton{
						Text:      "Open",
						OnClicked: mw.openDir,
					},
					HSpacer{},
					PushButton{
						Text:      "Start",
						OnClicked: mw.startWatch,
					},
					PushButton{
						Text:      "Stop",
						OnClicked: mw.stopWatch,
					},
				},
			},
		},
	}.Run()); err != nil {
		log.Fatal(err)
	}
}

func (mw *MyMainWindow) openDir() {
	dlg := new(walk.FileDialog)
	if ok, err := dlg.ShowBrowseFolder(mw); err != nil {
		log.Fatal(err)
	} else if !ok {
		return
	}

	mw.dirPath = dlg.FilePath
	setText(mw.textEdit, mw.dirPath)
}

func setText(te *walk.TextEdit, text string) {
	te.AppendText(time.Now().Format("2006-01-02 15:04:05 ") + text + "\r\n")
}

func (mw *MyMainWindow) startWatch() {
	if mw.isWatch {
		setText(mw.textEdit, "Watcher allready running")
		return
	}
	mw.isWatch = true
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	mw.watcher = watcher
	err = watcher.Add(mw.dirPath)
	if err != nil {
		log.Fatal(err)
	}
	go func(m *MyMainWindow) {
		for {
			select {
			case event, ok := <-m.watcher.Events:
				if !ok {
					return
				}
				setText(m.textEdit, event.Op.String()+" <==> "+filepath.Clean(event.Name))
				walk.MsgBox(m, "New action", "Some changes with "+filepath.Clean(event.Name),
					walk.MsgBoxIconInformation|walk.MsgBoxSetForeground)

			case _, ok := <-m.watcher.Errors:
				if !ok {
					return
				}
			case <-m.isDone:
				break
			}
		}
	}(mw)
	setText(mw.textEdit, "Start watching")
}

func (mw *MyMainWindow) stopWatch() {
	if !mw.isWatch {
		setText(mw.textEdit, "Watcher is not running")
		return
	}
	setText(mw.textEdit, "Stop watching")
	mw.isWatch = false
	mw.isDone <- true
	mw.watcher.Close()
}
