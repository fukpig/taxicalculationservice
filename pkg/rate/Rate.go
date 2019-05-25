package rate

type Rate struct {
	Id        int    `sql:"id"`
	Name      string `sql:"name"`
	PerMinute int32  `sql:"per_minute"`
	PerKm     int32  `sql:"per_km"`
}
