package codegen

// this code is generated by go generate
// DO NOT EDIT BY HAND!

type Collection struct {
	/* indicator about the type of the items in the collection (the default value is 'feature').	*/
	ItemType string `json:"itemType,omitempty"`
	Links []Link `json:"links"`
	/* human readable title of the collection	*/
	Title string `json:"title,omitempty"`
	/* the list of coordinate reference systems supported by the service	*/
	Crs []string `json:"crs,omitempty"`
	/* a description of the features in the collection	*/
	Description string `json:"description,omitempty"`
	/* The extent of the features in the collection. In the Core only spatial and temporal
    extents are specified. Extensions may add additional members to represent other
    extents, for example, thermal or pressure ranges.	*/
	Extent *Extent `json:"extent,omitempty"`
	/* identifier of the collection used, for example, in URIs	*/
	Id string `json:"id"`
}

type Collections struct {
	Collections []Collection `json:"collections"`
	Links []Link `json:"links"`
}

type ConfClasses struct {
	ConformsTo []string `json:"conformsTo"`
}

type Exception struct {
	Code string `json:"code"`
	Description string `json:"description,omitempty"`
}

type Extent struct {
	/* The spatial extent of the features in the collection.	*/
	Spatial *Spatial `json:"spatial,omitempty"`
	/* The temporal extent of the features in the collection.	*/
	Temporal *Temporal `json:"temporal,omitempty"`
}

type LandingPage struct {
	Description string `json:"description,omitempty"`
	Links []Link `json:"links"`
	Title string `json:"title,omitempty"`
}

type Link struct {
	Href string `json:"href"`
	Hreflang string `json:"hreflang,omitempty"`
	Length int64 `json:"length,omitempty"`
	Rel string `json:"rel,omitempty"`
	Title string `json:"title,omitempty"`
	Type string `json:"type,omitempty"`
}

type Spatial struct {
	/* One or more bounding boxes that describe the spatial extent of the dataset.
    In the Core only a single bounding box is supported. Extensions may support
    additional areas. If multiple areas are provided, the union of the bounding
    boxes describes the spatial extent.	*/
	Bbox [][]float64 `json:"bbox,omitempty"`
	/* Coordinate reference system of the coordinates in the spatial extent
    (property `bbox`). The default reference system is WGS 84 longitude/latitude.
    In the Core this is the only supported coordinate reference system.
    Extensions may support additional coordinate reference systems and add
    additional enum values.	*/
	Crs string `json:"crs,omitempty"`
}

type Temporal struct {
	/* Coordinate reference system of the coordinates in the temporal extent
    (property `interval`). The default reference system is the Gregorian calendar.
    In the Core this is the only supported temporal coordinate reference system.
    Extensions may support additional temporal coordinate reference systems and add
    additional enum values.	*/
	Trs string `json:"trs,omitempty"`
	/* One or more time intervals that describe the temporal extent of the dataset.
    The value `null` is supported and indicates an open time interval.
    In the Core only a single time interval is supported. Extensions may support
    multiple intervals. If multiple intervals are provided, the union of the
    intervals describes the temporal extent.	*/
	Interval [][]string `json:"interval,omitempty"`
}
