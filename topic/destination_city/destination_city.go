package destination_city

// Method 1
func DestCity(paths [][]string) string {
	cities := make(map[string]bool)

	for _, path := range paths {
		if _, exists := cities[path[0]]; exists {
			delete(cities, path[0])
		} else {
			cities[path[0]] = false
		}
		if _, exists := cities[path[1]]; exists {
			delete(cities, path[1])
		} else {
			cities[path[1]] = true
		}
	}

	for city, position := range cities {
		if position {
			return city
		}
	}

	return ""
}

// Method 2
func DestCity2(paths [][]string) string {
	cities := make(map[string]bool)

	for _, path := range paths {
		cities[path[0]] = true
	}

	for _, path := range paths {
		if !cities[path[1]] {
			return path[1]
		}
	}

	return ""
}
