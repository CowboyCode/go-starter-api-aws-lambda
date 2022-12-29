package search

// dataSearchLocationByCords - find location using longitude and latitude
func dataSearchLocationByCords(lat string, long string) (LocationResult, error) {

	var result LocationResult

	// ... search code

	// mock data
	result.Name = "Airlie Beach"
	result.Region = "Whitsundays"
	result.State = "Queensland"
	result.Country = "Australia"
	result.Latitude = lat
	result.Longitude = long

	return result, nil
}
