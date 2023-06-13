package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Video struct {
	Name     string
	Duration string
}

func listAllVideos(rw http.ResponseWriter, rq *http.Request) {
	v1 := Video{Name: "ironman", Duration: "90 mn"}
	v2 := Video{Name: "avenger", Duration: "120 mn"}
	v3 := Video{Name: "Simba", Duration: "130 mn"}

	// allvideosMap := "ironman"
	allvideosMap := []Video{
		1: v1,
		2: v2,
		3: v3,
	}

	/*Encode allvideosMap to byte using encoding/gob then writing responseWriter
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf) // Will write to buf.

	err := enc.Encode(allvideosMap)
	if err != nil {
		log.Fatal("Encode error:", err)
	}

	rw.WriteHeader(http.StatusOK)
	rw.Write(buf.Bytes())
	*/

	/*Create JSON from the instance data*/
	data, _ := json.Marshal(allvideosMap)

	rw.WriteHeader(http.StatusOK)
	rw.Write(data)

}

//Start Server
func Start() {

	http.HandleFunc("/", listAllVideos)
	log.Fatal(http.ListenAndServe(":3333", nil))
	fmt.Println("Server started, listening tp port:", 3333)

}
func main() {

	Start()

}
