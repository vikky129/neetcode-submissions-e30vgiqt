

type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	if len(strs) == 0 {
		return "#😊#"
	}

	var separator = "😊"
	return strings.Join(strs, separator)
}

func (s *Solution) Decode(encoded string) []string {
	if encoded == "#😊#" {
		return []string{}
	}

	var separator = "😊"
	return strings.Split(encoded, separator)
}