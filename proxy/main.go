package main

import (
	"encoding/json"
	"fmt"
)

type Video struct {
	Name     string
	Duration string
}

// Interface to service
type YoutubeServiceInterface interface {
	ListVideos() string
}

//Service struct or Class
type YoutubeService struct {
}

// YoutubeService struct implements the YoutubeServiceInterface interface with the ListVideos() method
func (y *YoutubeService) ListVideos() []Video {
	v1 := Video{Name: "ironman", Duration: "90 mn"}
	v2 := Video{Name: "avenger", Duration: "120 mn"}
	v3 := Video{Name: "Simba", Duration: "130 mn"}

	allvideosMap := []Video{
		1: v1,
		2: v2,
		3: v3,
	}

	data, _ := json.Marshal(allvideosMap)

	fmt.Println(string(data))
	return allvideosMap
}

// Proxy Object for YoutubeService struct
type YoutubeServiceProxy struct {
	iscached bool
	cache    map[string]Video
	service  YoutubeService
}

func newYoutubeServiceProxy() *YoutubeServiceProxy {
	return &YoutubeServiceProxy{
		iscached: false,
		service:  YoutubeService{},
		cache:    make(map[string]Video), // Key: Name of the video, Value: Video
	}
}

var cacheVideos map[string]Video

func (yp *YoutubeServiceProxy) ListVideos() {
	if len(yp.cache) == 0 {
		fmt.Println("Cache is empty, No ListVideos cache: Listing from original service instance")

		// Listing from original service instance
		all := yp.service.ListVideos()

		//Updating Proxy's cache
		for _, each := range all {

			cacheVideos = map[string]Video{
				each.Name: each,
			}
		}
		fmt.Println("Cache updated")

	} else {
		fmt.Println("All videos from cache:")
		fmt.Println("  Name |  Duration")
		for n, d := range yp.cache {

			fmt.Printf("- %s :  %s\n", n, d.Duration)
		}
	}
}

func main() {

	YoutubeClient := newYoutubeServiceProxy()
	YoutubeClient.cache = map[string]Video{
		"ironman": Video{Name: "ironman", Duration: "90 mn"},
		"Fast":    Video{Name: "Fast", Duration: "120 mn"},
	}
	YoutubeClient.ListVideos()

	fmt.Println("\nCache size: ", len(YoutubeClient.cache))

}
