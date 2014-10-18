package main

import(
  "net/http"
  "image"
  "image/jpeg"
  _ "image/png"
  _ "image/gif"
  "github.com/nfnt/resize"
  "strconv"
  "log"
)

// Scaler is a service that scales images and ....
func main() {
  http.HandleFunc("/", scaleImage)

  err := http.ListenAndServe(":12345", nil)
  if err != nil {
    log.Fatal(err)
  }
}

// scaleImage reads an image from r's request body and writes a resized version
// to w's response body.
//
// The desired dimensions are denoted by r's "Image-Width" & "Image-Height"
// headers.
func scaleImage(w http.ResponseWriter, r *http.Request) {
  var width, height uint64

  img, _, err := image.Decode(r.Body)
  if err != nil {
    w.WriteHeader(http.StatusBadRequest)
    w.Write([]byte(err.Error()))
    return
  }

  if r.Header["Image-Width"] == nil && r.Header["Image-Height"] == nil {
    w.WriteHeader(http.StatusBadRequest)
    w.Write([]byte("Both Image-Width & Image-Height headers are missing"))
    return
  }

  if r.Header["Image-Width"] != nil {
    width, err = strconv.ParseUint(r.Header["Image-Width"][0], 10, 16)
    if err != nil {
      w.WriteHeader(http.StatusBadRequest)
      w.Write([]byte("Malformed Image-Width header"))
      return
    }
  }

  if r.Header["Image-Height"] != nil {
    height, err = strconv.ParseUint(r.Header["Image-Height"][0], 10, 16)
    if err != nil {
      w.WriteHeader(http.StatusBadRequest)
      w.Write([]byte("Malformed Image-Height header"))
      return
    }
  }

  img = resize.Resize(uint(width), uint(height), img, resize.NearestNeighbor)

  err = jpeg.Encode(w, img, nil)
  if err != nil {
    panic(err)
  }
}
