package space

type Planet string

func Age(seconds float64, planet Planet) float64 {

	earthYear := 31557600.0
	planetYear := 1.0

	switch planet {
	case "Mercury":
		planetYear = 0.2408467
	case "Venus":
		planetYear = 0.61519726
	case "Mars":
		planetYear = 1.8808158
	case "Jupiter":
		planetYear = 11.862615
	case "Saturn":
		planetYear = 29.447498
	case "Uranus":
		planetYear = 84.01684
	case "Neptune":
		planetYear = 164.79132
	}

	age := seconds / planetYear / earthYear
	return age
}
