package main

import (
	"encoding/csv"
	"os"
)

func main() {
	writer := csv.NewWriter(os.Stdout)
	_ = writer.Write([]string{"Muhammad", "Maulana", "givari"})
	_ = writer.Write([]string{"Test", "Kocak", "Gaming"})
	_ = writer.Write([]string{"Yaha", "BBB", "CC"})

	writer.Flush()
}
