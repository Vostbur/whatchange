from tkinter import *
from tkinter import filedialog
from tkinter import messagebox as mb
from watchdog.observers import Observer
from watchdog.events import PatternMatchingEventHandler
import time

class Watchdog(PatternMatchingEventHandler, Observer):
    def __init__(self, path='.', patterns='*', logfunc=print):
        PatternMatchingEventHandler.__init__(self, patterns)
        Observer.__init__(self)
        self.schedule(self, path=path, recursive=False)
        self.log = logfunc

    def on_created(self, event):
        self.log(f"{event.src_path} добавлен")

    def on_deleted(self, event):
        self.log(f"Удалён {event.src_path}!")

    def on_modified(self, event):
        self.log(f"{event.src_path} изменён")

    def on_moved(self, event):
        self.log(f"{event.src_path} перемещён в {event.dest_path}")

class GUI:
    def __init__(self):
        self.watchdog = None
        self.watch_path = '.'
        self.root = Tk()
        self.root.attributes('-topmost', True)
        self.root.title('Контроль и оповещение об изменениях в каталоге')
        self.messagebox = Text(width=80, height=10)
        self.scrollbar_ver = Scrollbar(self.root, orient=VERTICAL, command=self.messagebox.yview)
        self.messagebox.configure(yscrollcommand=self.scrollbar_ver.set)
        self.scrollbar_ver.pack(side=RIGHT, fill=Y)
        self.messagebox.pack(fill=BOTH, expand=1)
        frm = Frame(self.root)
        Button(frm, padx="20", pady="4", width="10", bg="#DCDCDC", activebackground="#ADD8E6", text='Выбрать каталог', command=self.select_path).pack(side=LEFT)
        Button(frm, padx="20", pady="4", width="10", bg="#DCDCDC", activebackground="#ADD8E6", text='Остановить', command=self.stop_watchdog).pack(side=RIGHT)
        Button(frm, padx="20", pady="4", width="10", bg="#DCDCDC", activebackground="#ADD8E6", text='Запустить', command=self.start_watchdog).pack(side=RIGHT)
        frm.pack(fill=X, expand=0)
        self.root.mainloop()

    def start_watchdog(self):
        if self.watchdog is None:
            self.watchdog = Watchdog(path=self.watch_path, logfunc=self.log)
            self.watchdog.start()
            self.log('Контроль запущен')
        else:
            self.log('Контроль уже запущен')

    def stop_watchdog(self):
        if self.watchdog:
            self.watchdog.stop()
            self.watchdog = None
            self.log('Контроль остановлен')
        else:
            self.log('Контроль не запущен')

    def select_path(self):
        path = filedialog.askdirectory()
        if path:
            self.watch_path = path
            self.log(f'Выбран каталог: {path}')

    def log(self, message):
        mb.showinfo(title="Новое сообщение", message=f'{message}\n')
        now = time.strftime("%Y-%m-%d %H:%M:%S", time.localtime())
        self.messagebox.insert(END, f'{now} --> {message}\n')
        self.messagebox.see(END)

if __name__ == '__main__':
    GUI()


