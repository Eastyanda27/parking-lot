# 🚗 Parking Lot CLI App - Golang

A command-line application written in Go that simulates an automated parking lot ticketing system.  
Cars are parked in the nearest available slot, and charges are calculated based on parking duration.

---

## 📦 Features

- Create a parking lot of fixed capacity  
- Park cars in nearest available slots  
- Unpark cars and calculate parking charges  
- View current parking status  
- Read commands from an input file  

---

## ⚙️ How to Run
```bash
1. Clone or Download the Project

git clone https://github.com/Eastyanda27/parking-lot
cd parking-lot


2. Prepare Your Input File
Create a file named input.txt with parking commands. Example:

create_parking_lot 6
park KA-01-HH-1234
park KA-01-HH-9999
leave KA-01-HH-1234 4
status
You can modify or add more commands as needed.


3. Run the App

go run main.go input.txt
You’ll see output on the terminal based on the commands in the file.


🧪 Example Output

Created parking lot with 6 slots
Allocated slot number: 1
Allocated slot number: 2
Registration number KA-01-HH-1234 with Slot Number 1 is free with Charge $30
Slot No. Registration No.
2 KA-01-HH-9999


📁 Project Structure

parking-lot/
├── main.go             # Entry point
├── input.txt           # Command input file
├── go.mod              # Go module file
├── parking/
│   └── parking_lot.go  # Core parking logic
