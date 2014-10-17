package main

import(
  "image"
  "net/http"
  "fmt"
  "github.com/nfnt/resize"
  "image/jpeg"
)

func main() {
  http.HandleFunc("/", foo)

  err := http.ListenAndServe(":12345", nil)
  if err != nil {
    panic(err)
  }
}

func foo(w http.ResponseWriter, r *http.Request) {
  img, _, err := image.Decode(r.Body)
  if err != nil {
    fmt.Println(err)
  }

  img = resize.Resize(300, 300, img, resize.NearestNeighbor)

  err = jpeg.Encode(w, img, nil)
  if err != nil {
    panic(err)
  }
}
