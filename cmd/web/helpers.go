package main

import (
	"fmt"
	"net/http"
)



func (app *application) render(w http.ResponseWriter, r *http.Request, name string,
	td *templateData) {
	ts, ok := app.templateCache[name]

	if !ok {
		errorType.ServerError(w, fmt.Errorf("the template %s does not exist", name))
	}


	err := ts.Execute(w, td)
	if err != nil {
		errorType.ServerError(w, err)
	}

}
