package main

const resolution8K = 7_680 * 4_320 * 3

func main() {
	selfie := [resolution8K]byte{
		/* ..color data.. */
	}
	selfie = invertColors(selfie)
	// Process the selfie further...
}

func invertColors(colors [resolution8K]byte) [resolution8K]byte {
	for i := range colors {
		colors[i] = 255 - colors[i]
	}
	return colors
}
