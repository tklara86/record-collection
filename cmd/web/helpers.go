package main

import (
	"bytes"
	"fmt"
	"net/http"
	"time"
)


func (app *application) common(td *templateData) *templateData {
	if td == nil {
		td = &templateData{}
	}

	// Links
	td.Links = []link{
		{
			LinkTitle: "Home",
			LinkPath:  "/",
		},
		{
			LinkTitle: "Add record",
			LinkPath:  "/record/create",
		},
	}

	// current year
	td.CurrentYear = time.Now().Year()

	return td
}


func (app *application) render(w http.ResponseWriter, r *http.Request, name string,
	td *templateData) {
	ts, ok := app.templateCache[name]

	if !ok {
		errorType.ServerError(w, fmt.Errorf("the template %s does not exist", name))
		return
	}

	buf := new(bytes.Buffer)

	err := ts.Execute(buf, app.common(td))
	if err != nil {
		errorType.ServerError(w, err)
	}

	buf.WriteTo(w)
}
