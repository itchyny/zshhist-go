package zshhist

const meta = 0x83

// Unmetafy a metafied line
func Unmetafy(str string) string {
	var j int
	bs := []byte(str)
	for i := 0; i < len(bs); i++ {
		if bs[i] == meta {
			i++
			bs[j] = bs[i] ^ 32
		} else {
			bs[j] = bs[i]
		}
		j++
	}
	return string(bs[:j])
}
