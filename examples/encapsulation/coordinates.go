package encapsulation

type Coordinates struct {
	latitude  float64
	longitude float64
}

func (c *Coordinates) SetLatitude(latitude float64) {
	c.latitude = latitude
}

func (c *Coordinates) SetLongitude(longitude float64) {
	c.longitude = longitude
}
