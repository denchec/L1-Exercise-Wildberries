package main

import "fmt"

func main() {
	sequence := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	newMap := make(map[int][]float64)

	for i := 0; i < len(sequence); i++ {
		if sequence[i] == 0 {
			newMap[0] = append(newMap[0], sequence[i])
		} else if sequence[i] > 0 {
			if sequence[i] > 0 && sequence[i] <= 10 {
				newMap[0] = append(newMap[0], sequence[i])
			} else if sequence[i] > 10 && sequence[i] <= 20 {
				newMap[10] = append(newMap[10], sequence[i])
			} else if sequence[i] > 20 && sequence[i] <= 30 {
				newMap[20] = append(newMap[20], sequence[i])
			} else if sequence[i] > 30 && sequence[i] <= 40 {
				newMap[30] = append(newMap[30], sequence[i])
			}
		} else {
			if sequence[i] < 0 && sequence[i] >= -10 {
				newMap[-0] = append(newMap[-0], sequence[i])
			} else if sequence[i] < -10 && sequence[i] >= -20 {
				newMap[-10] = append(newMap[-10], sequence[i])
			} else if sequence[i] < -20 && sequence[i] >= -30 {
				newMap[-20] = append(newMap[-20], sequence[i])
			} else if sequence[i] < -30 && sequence[i] >= -40 {
				newMap[-30] = append(newMap[-30], sequence[i])
			}
		}
	}
	fmt.Println(newMap)
}
