package prayer

var Location City

type City struct {
	ID       uint64  `json:"id"`
	Title    string  `json:"title"`
	Lng      float64 `json:"lng"` // Longitude
	Lat      float64 `json:"lat"` // Latitude
	Timezone int     `json:"timezone"`
	Region   string  `json:"region"`
	District string  `json:"district"`
}

var Astana = City{
	ID:       3,
	Title:    "Астана қаласы",
	Lng:      71.433333,
	Lat:      51.133333,
	Timezone: 5,
	Region:   "Республикалық маңызы бар қалалар",
	District: "",
}

var Shymkent = City{
	ID:       57,
	Title:    "Шымкент қаласы",
	Lng:      69.612769,
	Lat:      42.368009,
	Timezone: 5,
	Region:   "Республикалық маңызы бар қалалар",
	District: "",
}

var Almaty = City{
	ID:       72,
	Title:    "Алматы қаласы",
	Lng:      76.945465,
	Lat:      43.238293,
	Timezone: 5,
	Region:   "Республикалық маңызы бар қалалар",
	District: "",
}

var Cities = []City{Astana, Almaty, Shymkent}
