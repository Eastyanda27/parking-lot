package main

import (
	"bufio"
	"fmt"
	"os"
	"parkinglot/parking" // import package parking (folder parking)
	"strings"
)

func main() {
	// Cek apakah user memberikan nama file input sebagai argumen
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go input.txt")
		return
	}

	// Baca nama file dari argumen
	fileName := os.Args[1]
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Inisialisasi parking lot
	lot := parking.NewParkingLot()

	// Scanner untuk membaca file baris per baris
	scanner := bufio.NewScanner(file)

	// Proses setiap baris perintah
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line) // Pisahkan per kata

		if len(parts) == 0 {
			continue // Lewati baris kosong
		}

		// Jalankan perintah sesuai input
		switch parts[0] {
		case "create_parking_lot":
			capacity := parking.Atoi(parts[1])
			lot.Create(capacity)

		case "park":
			lot.Park(parts[1])

		case "leave":
			lot.Leave(parts[1], parking.Atoi(parts[2]))

		case "status":
			lot.Status()
		}
	}
}
