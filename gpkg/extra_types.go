package gpkg

import (
	"oaf-server/codegen"

	"github.com/go-spatial/geom/encoding/geojson"
)

// this code is generated by go generate
// Not anymore geojson.FeatureType included

type FeatureCollectionGeoJSON struct {
	NumberReturned int64          `json:"numberReturned,omitempty"`
	TimeStamp      string         `json:"timeStamp,omitempty"`
	Type           string         `json:"type"`
	Features       []*Feature     `json:"features"`
	Links          []codegen.Link `json:"links,omitempty"`
	NumberMatched  int64          `json:"numberMatched,omitempty"`
	Crs            string         `json:"crs,omitempty"`
}

type Feature struct {
	Type string      `json:"type"`
	ID   interface{} `json:"id,omitempty"`
	// can be null
	Geometry geojson.Geometry `json:"geometry"`
	// can be null
	Properties map[string]interface{} `json:"properties"`
	// can be null
	Links []codegen.Link `json:"links,omitempty"`
}