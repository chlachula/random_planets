package main

import (
	"math/rand"
	"net/http"
	"os"
)

var version = "0.0.1"
var planet_names = []string{"mercury", "venus", "earth", "mars", "jupiter", "saturn", "uranus", "neptune"}
var planet_diameters = []int{4878, 12104, 12756, 6794, 142984, 120536, 51118, 49532}
var partial_sums []int
var sum int
var planet string

func initPlanetNames() {
	partial_sums = make([]int, len(planet_diameters))
	for i, diameter := range planet_diameters {
		sum += diameter
		partial_sums[i] = sum
	}
}
func generatePlanetName() {
	n := rand.Intn(sum)
	for i, ps := range partial_sums {
		if n < ps {
			planet = planet_names[i]
			break
		}
	}
}
func endpointRoot(w http.ResponseWriter, req *http.Request) {
	_ = req
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte("<html><header><title>Random planets</title></header>\n<body>\n"))
	w.Write([]byte("<h1>Random planets</h1>\n"))
	multiline := `<p>
Generates planets name based tags every second at <a href="metrics">metrics</a>
with probability based on planet diameter.
</p>`
	w.Write([]byte(multiline))
	w.Write([]byte("</body></html>\n"))
}

func endpointMetrics(w http.ResponseWriter, req *http.Request) {
	_ = req
	generatePlanetName()
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("planet_"))
	w.Write([]byte(planet))
	w.Write([]byte("\n"))
}

func main() {
	initPlanetNames()
	http.HandleFunc("/", endpointRoot)
	http.HandleFunc("/metrics", endpointMetrics)

	port := os.Getenv("RANDOM_PLANETS_PORT")
	if port == "" {
		port = "80"
	}
	println("random_planets version " + version)
	println("going to listen and serve at port " + port)
	http.ListenAndServe(":"+port, nil)
}
