package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	http.HandleFunc("/", clock)
	http.HandleFunc("/entries", entries)
	http.HandleFunc("/add", addPayload)

	http.ListenAndServe(":4567", nil)
}

func entries(w http.ResponseWriter, req *http.Request)  {
	if req.Method == http.MethodGet {
		content, err := ioutil.ReadFile("./data")

		if err != nil {
			fmt.Fprintf(w, "nothing to read")
		} else {
			fmt.Fprintf(w, string(content))
		}
	} else {
		fmt.Fprintf(w, "Méthode non prise en charge")
	}
}

func addPayload(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		if err := req.ParseForm(); err != nil || req.PostForm["author"] == nil || req.PostForm["entry"] == nil{
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("400 - Bad request"))
			return
		}
		fmt.Fprintf(w, "%v:%v", req.PostForm["author"],req.PostForm["entry"])
		saveFile, _ := os.OpenFile("./data", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755)

		defer saveFile.Close()

		saveFile.WriteString("["+ req.PostForm["author"][0]+"]:["+req.PostForm["entry"][0]+"]\n")
	} else {
		fmt.Fprintf(w, "Méthode non prise en charge")
	}
}

func clock(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		currentTime := time.Now()
		fmt.Fprintf(w, "Il est %v", currentTime.Format("15h04"))
	} else {
		fmt.Fprintf(w, "Méthode non prise en charge")
	}
}