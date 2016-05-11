package main

import (
  "fmt"
  "io"
  "io/ioutil"
  "net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
  resp, err := http.Get("http://169.254.169.254/latest/meta-data/instance-id")
  if err != nil {
    fmt.Println(err)
  }

  defer resp.Body.Close()
  body, err := ioutil.ReadAll(resp.Body)

  io.WriteString(w, "Instance ID: " + string(body))
}

func main() {
  http.HandleFunc("/", index)
  http.ListenAndServe(":8000", nil)
}