package data

type (
	distance   float64
	distanceKm float64
)

func (miles distance) ToKm() distanceKm {
	return distanceKm(miles * 1.60934)
}

func (km distanceKm) ToMiles() distance {
	return distance(km / 1.60934)
}

func test() {
	d := distance(4.5)
	d.ToKm()
	print(d)
}
