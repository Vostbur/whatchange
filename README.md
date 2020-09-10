# whatchange
Tracking changes in a folder on Go with GUI.
Golang versions are based on package [fsnotify/fsnotify](https://github.com/fsnotify/fsnotify).
Python version uses [gorakhargosh/watchdog](https://github.com/gorakhargosh/watchdog).

## Python Tkinter version

![](/images/python-tkinter-watchdog.PNG)

## Golang GTK3 version

![](/images/go-gtk3-watchdog.PNG)

## Golang [walk](https://github.com/lxn/walk) version

![](/images/go-walk-watchdog.PNG)

Build:

```
go get github.com/lxn/walk
go get github.com/akavel/rsrc
rsrc -manifest app.manifest -o rsrc.syso
go build -ldflags="-H windowsgui"
```
