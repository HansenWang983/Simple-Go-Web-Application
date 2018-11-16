package service

import (
    "net/http"
    "github.com/unrolled/render"
    "log"
    "html/template"
)

func homeHandler(formatter *render.Render) http.HandlerFunc {

    return func(w http.ResponseWriter, req *http.Request) {
        formatter.HTML(w, http.StatusOK, "index", struct {
            ID      string `json:"id"`
            Content string `json:"content"`
        }{ID: "8675309", Content: "Hello from Go!"})
    }
}

// when request URL path = "/unknown", return status code 5xx
func NotImplementedHandler(formatter *render.Render) http.HandlerFunc {

    return func(w http.ResponseWriter, req *http.Request) {
      if req.URL.Path == "/unknown" {
        t, err := template.ParseFiles("src/github.com/hansenbeast/cloudgo-io/templates/5xx.html")
        if err != nil {
          log.Println(err)
        }
        t.Execute(w, nil)  
      }     
    }
  }
  

  func loginHandler(formatter *render.Render) http.HandlerFunc {

    return func(w http.ResponseWriter, req *http.Request) {
      if req.URL.Path == "/login" {
        t, err := template.ParseFiles("src/github.com/hansenbeast/cloudgo-io/templates/login.html")
  
        req.ParseForm()
        ua := req.FormValue("username")
        pa := req.FormValue("password")
        formatter.HTML(w, http.StatusOK, "table", struct {
          Username      string `json:"username"`
          Password        string `json:"password"`
        } {Username: ua, Password: pa})
        if err != nil {
          log.Println(err)
        }
        t.Execute(w, nil)  
      }     
    }
}