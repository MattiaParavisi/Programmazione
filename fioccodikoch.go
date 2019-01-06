// Requires go get github.com/holizz/terrapin
package main

import (
	"fmt"
	"image"
	"image/png"
	"math"
	"net/http"

	"github.com/holizz/terrapin"
)

func kochCurve(t *terrapin.Terrapin, lung float64, liv int) {
	if liv == 0 {
		t.Forward(lung)
	} else {
		kochCurve(t, lung, liv-1)
		t.Left(math.Pi / 3.0)
		kochCurve(t, lung, liv-1)
		t.Right(2.0 * math.Pi / 3.0)
		kochCurve(t, lung, liv-1)
		t.Left(math.Pi / 3.0)
		kochCurve(t, lung, liv-1)
	}
}

func kochSnowflake(t *terrapin.Terrapin, lung float64, liv int) {
	kochCurve(t, lung, liv)
	t.Right(2.0 * math.Pi / 3.0)
	kochCurve(t, lung, liv)
	t.Right(2.0 * math.Pi / 3.0)
	kochCurve(t, lung, liv)
	t.Right(2.0 * math.Pi / 3.0)
}

func pageMain(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`
			<!doctype html>
			<title>Koch</title>
			<h1>Koch</h1>
			<img src="/koch.png">
			`))
}

func pageImage(w http.ResponseWriter, r *http.Request) {
	i := image.NewRGBA(image.Rect(0, 0, 500, 500))
	t := terrapin.NewTerrapin(i, terrapin.Position{250.0, 450.0})

	kochSnowflake(t, 10, 3)

	png.Encode(w, i)
}

func main() {
	http.HandleFunc("/", pageMain)
	http.HandleFunc("/koch.png", pageImage)
	fmt.Println("Listening on http://localhost:3000/")
	http.ListenAndServe(":3000", nil)
}
