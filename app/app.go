package app

import (
    "fmt"
    "net/http"
    "net/url"
    "io/ioutil"
    "regexp"
    "path/filepath"
)



func init() {
    http.HandleFunc("/", pageHandler)
    http.HandleFunc("/category/", categoryHandler)
    http.HandleFunc("/feed/", feedHandler)
}

func categoryHandler(w http.ResponseWriter, r *http.Request) {
    rePath := regexp.MustCompile("^/category")
    http.Redirect(w, r, rePath.ReplaceAllString(r.URL.String(), ""), 301)
} // categoryHandler

func feedHandler(w http.ResponseWriter, r *http.Request) {
    http.Redirect(w, r, "/atom.xml", 301)
} // feedHandler

func pageHandler(w http.ResponseWriter, r *http.Request) {
    url, _ := url.Parse(r.URL.String())
    body, err := getFile(url.Path)

    if err != nil {
        http.Error(w, fmt.Sprintf("Could not find file: %s", url.Path), 404)
    }

    fmt.Fprintf(w, "%s", string(body))
}

func getFile(urlPath string) ([]byte, error) {
    rePath := regexp.MustCompile("/(|index|index.html)$")
    reTrailingHtml := regexp.MustCompile(".html$")
    trimmedPath := reTrailingHtml.ReplaceAllString(rePath.ReplaceAllString(urlPath, ""), "")
    deployPath, _ := filepath.Abs("deploy/")
    
    body, err := ioutil.ReadFile(fmt.Sprintf("%s%s.html", deployPath, trimmedPath))
    
    if err != nil {
        body, err = ioutil.ReadFile(fmt.Sprintf("%s%s/index.html", deployPath, trimmedPath))
    }
    
    return body, err;
} // getFile