package distanceUtil

import "math"

const EarthRadius = float64(6371)

var sinPointSquare = func(val float64) float64 { return math.Pow(math.Sin(val), 2) }
var piRadian = func() float64 { return math.Pi / 180.0 }

func haversine(onePhi, oneLambda, twoPhi, twoLambda float64) (c float64) {
	dPhi := (twoPhi - onePhi) * piRadian()
	dLambda := (twoLambda - oneLambda) * piRadian()

	phi := onePhi * piRadian()
	lambda := twoPhi * piRadian()

	a1 := sinPointSquare(dPhi / 2)
	a2 := sinPointSquare(dLambda/2) * (math.Cos(phi) * math.Cos(lambda))

	a := a1 + a2
	c = 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	return
}

func Haversine(lat1, lng1, lat2, lng2 float64) float64 {
	distInMtr := 1000 * EarthRadius * haversine(lat1, lng1, lat2, lng2)
	return distInMtr
}
