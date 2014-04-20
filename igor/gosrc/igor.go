package main

import ("./moo"
  "flag"
 )

var gui moo.GUI

func main() {
  useQml := flag.Bool("gui", true, "use the graphical interface")
  if(*useQml == true) {
    gui = &moo.QmlGUI {}
  } else {
    gui = &moo.TermboxGUI {}
  }
  gui.Main()
}
