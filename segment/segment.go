package segment

import "agent/config"

func Split(data []byte) [][]byte {
	var segments [][]byte

	for i := 0; i < len(data); i += config.SegmentSize {
		end := i + config.SegmentSize
		if end > len(data) {
			end = len(data)
		}
		segments = append(segments, data[i:end])
	}

	return segments
}
