package netcontrol

import (
	"fmt"
	"net/http"
)

//Register  urlHelp
func reqHelpRegister() {

	//ONE_POIN
	http.HandleFunc("/help", func(
		w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, helpContent)
	}) //HandleFunc

}

var helpContent string = "" +

	"This is help content\n" +
	"for user guide"
