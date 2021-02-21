package input

import (
	"fmt"
	geojson "github.com/paulmach/go.geojson"
)

func ReadGeoJson(input []byte) {
	content, err := geojson.UnmarshalFeatureCollection(input)
	if err != nil {
		fmt.Println("Could not parse GeoJSON")
		return
	}
	if content.Type != "FeatureCollection" {
		fmt.Printf("should have type of FeatureCollection, got %v", content.Type)
	}
	if len(content.Features) <= 0 {
		fmt.Println("Need at least 1 Feature")
	}
	fmt.Println(content.Features[0].Geometry.IsLineString())

}
