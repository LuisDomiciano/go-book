package main

import (
	"net/http"
)

func (app *application) bookHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]string {
		"status": "avaliable",
		"environment": app.config.env,
		"version": version,
	}

	err := app.writeJSON(w, http.StatusOK, data, nil)

	if err != nil {
		app.logger.Println(err)
		http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
	}
/*
	js, err := json.Marshal(data)
	if err != nil {
		app.logger.Println(err)
		http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
		return
	}
	js = append(js, '\n')
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
*/
}
