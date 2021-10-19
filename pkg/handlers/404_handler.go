package handlers

import (
	"net/http"
	"os"
	"path"
	"strings"
)

type FSHandler404 = func(w http.ResponseWriter, r *http.Request) (doDefaultFileServe bool)

func FileSystem404(w http.ResponseWriter, r *http.Request) (doDefaultFileServe bool) {
	r.URL.Path = "/"
	return true
}

func FileServerWith404(root http.FileSystem, handler404 FSHandler404) http.Handler {
	fs := http.FileServer(root)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//make sure the url path starts with /
		upath := r.URL.Path
		if !strings.HasPrefix(upath, "/") {
			upath = "/" + upath
			r.URL.Path = upath
		}
		upath = path.Clean(upath)

		// attempt to open the file via the http.FileSystem
		f, err := root.Open(upath)
		if err != nil {
			if os.IsNotExist(err) {
				// call handler
				if handler404 != nil {
					doDefault := handler404(w, r)
					if !doDefault {
						return
					}
				}
			}
		}

		// close if successfully opened
		if err == nil {
			f.Close()
		}

		fs.ServeHTTP(w, r)
	})
}
