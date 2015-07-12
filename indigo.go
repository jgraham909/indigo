package main

import (
  "html/template"
  "log"
  "net/http"
  "os"
  "path"
)

const webDir = "/root/static/"
const templateDir = webDir + "templates/"
const contentDir = webDir + "content/"
const staticDir = webDir + "static/"

func main() {
  serveSingle("/favicon.ico", staticDir + "/favicon.ico")
  fs := http.FileServer(http.Dir(staticDir))
  http.Handle("/static/", http.StripPrefix("/static", fs))
  http.HandleFunc("/", serveTemplate)

  log.Println("Listening...")
  http.ListenAndServe(":80", nil)
}

func serveTemplate(w http.ResponseWriter, r *http.Request) {
  var fp string;
  pp := path.Join(templateDir, "page.html")
  hp := path.Join(templateDir, "header.html")
  if (r.URL.Path == "/") {
    fp = path.Join(contentDir, "/index.html")
  } else {
    fp = path.Join(contentDir, r.URL.Path)
  }
  // Return a 404 if the template doesn't exist
  info, err := os.Stat(fp)
  if err != nil {
    if os.IsNotExist(err) {
      log.Println(err.Error())
      http.NotFound(w, r)
      return
    }
  }


  // Return a 404 if the request is for a directory
  if info.IsDir() {
    log.Println("Directory listing requested")
    http.NotFound(w, r)
    return
  }

  tmpl, err := template.ParseFiles(pp, hp, fp)
  if err != nil {
    // Log the detailed error
    log.Println(err.Error())
    // Return a generic "Internal Server Error" message
    http.Error(w, http.StatusText(500), 500)
    return
  }

  if err := tmpl.ExecuteTemplate(w, "page", nil); err != nil {
    log.Println(err.Error())
    http.Error(w, http.StatusText(500), 500)
  }
}

func serveSingle(pattern string, filename string) {
    http.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, filename)
    })
}
