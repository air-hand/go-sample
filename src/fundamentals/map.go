package fundamentals

type MapValue struct {
	Age int
}

func (mapValue *MapValue) GetAge() int {
	return mapValue.Age
}

func MapSample() map[string]MapValue {
	map_value := make(map[string]MapValue)

	map_value["Bob"] = MapValue{Age: 100}
	map_value["Alice"] = MapValue{Age: 200}
	map_value["John Doe"] = MapValue{Age: 100000}

	return map_value
}
