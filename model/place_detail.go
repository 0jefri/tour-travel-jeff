package model

type PlaceDetail struct {
	ID     int
	Place  Place
	Review Review
	Galery []PhotoGroup
}

type PhotoGroup struct {
	Photos []Photo
}

type Photo struct {
	ID      int
	URL     string
	Caption string
}
