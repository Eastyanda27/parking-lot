package parking

import (
	"fmt"
	"sort"
)

// Struct Car menyimpan info mobil dan slot yang ditempati
type Car struct {
	Number string // Nomor plat mobil
	Slot   int    // Slot tempat parkir
}

// ParkingLot menyimpan informasi keadaan tempat parkir
type ParkingLot struct {
	Capacity  int            // Kapasitas maksimum
	Slots     map[int]*Car   // Slot yang terisi (slot -> mobil)
	Occupied  map[string]int // Map nomor mobil -> slot
	FreeSlots []int          // Daftar slot kosong (siap dipakai)
}

// Constructor untuk buat ParkingLot baru
func NewParkingLot() *ParkingLot {
	return &ParkingLot{
		Slots:     make(map[int]*Car),
		Occupied:  make(map[string]int),
		FreeSlots: []int{},
	}
}

// Fungsi helper untuk parsing angka dari string
func Atoi(s string) int {
	var n int
	fmt.Sscanf(s, "%d", &n)
	return n
}

// Buat parking lot dengan jumlah slot tertentu
func (p *ParkingLot) Create(capacity int) {
	p.Capacity = capacity
	p.FreeSlots = make([]int, capacity)

	// Isi FreeSlots dengan angka 1 sampai n
	for i := 0; i < capacity; i++ {
		p.FreeSlots[i] = i + 1
	}
	fmt.Printf("Created parking lot with %d slots\n", capacity)
}

// Parkirkan mobil di slot terdekat (slot terkecil)
func (p *ParkingLot) Park(number string) {
	if len(p.FreeSlots) == 0 {
		fmt.Println("Sorry, parking lot is full")
		return
	}

	// Ambil slot terkecil
	sort.Ints(p.FreeSlots)
	slot := p.FreeSlots[0]
	p.FreeSlots = p.FreeSlots[1:] // Hapus dari list kosong

	// Simpan mobil ke slot
	car := &Car{Number: number, Slot: slot}
	p.Slots[slot] = car
	p.Occupied[number] = slot

	fmt.Printf("Allocated slot number: %d\n", slot)
}

// Keluarkan mobil dan hitung biaya parkir
func (p *ParkingLot) Leave(number string, hours int) {
	slot, found := p.Occupied[number]
	if !found {
		fmt.Printf("Registration number %s not found\n", number)
		return
	}

	// Hapus mobil dari data
	delete(p.Slots, slot)
	delete(p.Occupied, number)
	p.FreeSlots = append(p.FreeSlots, slot)

	// Hitung biaya: $10 untuk 2 jam pertama, +$10/jam berikutnya
	charge := 10
	if hours > 2 {
		charge += (hours - 2) * 10
	}

	fmt.Printf("Registration number %s with Slot Number %d is free with Charge $%d\n", number, slot, charge)
}

// Tampilkan status semua slot terisi
func (p *ParkingLot) Status() {
	fmt.Println("Slot No. Registration No.")

	// Urutkan slot untuk tampilan rapi
	keys := make([]int, 0, len(p.Slots))
	for slot := range p.Slots {
		keys = append(keys, slot)
	}
	sort.Ints(keys)

	for _, slot := range keys {
		fmt.Printf("%d %s\n", slot, p.Slots[slot].Number)
	}
}
