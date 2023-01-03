package main

import "fmt"

func main() {
	var h1, m1, h2, m2 int
	x, y := 0, 0
	fmt.Scanf("%d %d %d %d", &h1, &m1, &h2, &m2)

	if h1 == h2 {
		if m1 == m2 {
			x = 24
			y = 0
		} else {
			if m2 < m1 {
				x = 23
				y = 60 - (m1 - m2)
			} else {
				x = 0
				y = m2 - m1
			}

		}
	} else if h1 < h2 {
		x = (((h2 * 60) + m2) - ((h1 * 60) + m1)) / 60
		y = (((h2 * 60) + m2) - ((h1 * 60) + m1)) % 60

	} else {
		x = (1440 - (((h1 * 60) + m1) - ((h2 * 60) + m2))) / 60
		y = (1440 - (((h1 * 60) + m1) - ((h2 * 60) + m2))) % 60
	}

	fmt.Println("O JOGO DUROU", x, "HORA(S) E", y, "MINUTO(S)")

}
