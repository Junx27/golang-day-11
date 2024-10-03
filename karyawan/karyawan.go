package karyawan

import "fmt"

type Karyawan struct {
	Nama    string
	Umur    int
	Jabatan string
}

func (k *Karyawan) UbahNama(nama string) {
	k.Nama = nama
	fmt.Printf("Nama karyawan diubah %s\n", k.Nama)
}
func (k *Karyawan) UbahUmur(umur int) {
	k.Umur = umur
	fmt.Printf("umur karyawan diubah %d\n", k.Umur)
}
func (k *Karyawan) UbahJabatan(jabatan string) {
	k.Jabatan = jabatan
	fmt.Printf("jabatan karyawan diubah %s\n", k.Jabatan)
}
