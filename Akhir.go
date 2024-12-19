/* NOTE
Untuk Login ke Menu Admin Silahkan ketik "admin"
*/

package main

import (
	"fmt"
	"time"
)

// Nilai Konstanta NMAX
const NMAX int = 512

// Type Struct untuk admin
type CalonLegislatif struct {
	nama   string
	partai string
	suara  int
}

// type struct untuk pemilih
type PemilihanUmum struct {
	nama     string
	kandidat string
}

// Type Alias dengan elemen NMAX
type DaftarCalonLegislatif [NMAX]CalonLegislatif
type DaftarPemilu [NMAX]PemilihanUmum

func main() {
	var legislatif DaftarCalonLegislatif
	var pemilu DaftarPemilu
	var InputMenu, SubInput int
	var nLegislatif, nPemilu int
	var Waktu bool = false
	var i2 bool = true
	var i bool = true

	fmt.Print("\033[H\033[2J")
	for i {
		//cetak daftar menu
		DaftarMenu()

		//input daftar menu
		InputMenu = InputDaftarMenu()
		fmt.Print("\033[H\033[2J")
		if InputMenu == 1 {
			fmt.Print("\033[H\033[2J")
			MemilihCalonLegislatif(&legislatif, nLegislatif, &pemilu, &nPemilu, Waktu)
		} else if InputMenu == 2 {
			TampilanDataLegislatif()
			SubInput = ProsesMenampilkanDataLegislatif()
			if SubInput == 1 {
				fmt.Print("\033[H\033[2J")
				DataBerdasarkanNama(&legislatif, nLegislatif)
			} else if SubInput == 2 {
				fmt.Print("\033[H\033[2J")
				DataBerdasarkanPartai(&legislatif, nLegislatif)
			} else if SubInput == 3 {
				fmt.Print("\033[H\033[2J")
				DataBerdasarkanSuara(&legislatif, nLegislatif)
			}
		} else if InputMenu == 3 {
			fmt.Print("\033[H\033[2J")
			CariDataLegislatif()
			SubInput = ProsesCariDataLegislatif()
			if SubInput == 1 {
				fmt.Print("\033[H\033[2J")
				CariBerdasarkanNama(&legislatif, nLegislatif)
			} else if SubInput == 2 {
				fmt.Print("\033[H\033[2J")
				CariBerdasarkanPartai(&legislatif, nLegislatif)
			} else if SubInput == 3 {
				fmt.Print("\033[H\033[2J")
				CariBerdasarkanPemilih(&legislatif, pemilu, nLegislatif, nPemilu)
			}
		} else if InputMenu == 4 {
			fmt.Print("\033[H\033[2J")
			ThresholdKandidat(&legislatif, nLegislatif, nPemilu)
		} else if InputMenu == 5 {
			fmt.Print("\033[H\033[2J")
			i = false
		} else if InputMenu == 6 {
			fmt.Print("\033[H\033[2J")
			i2 = true
			for i2 {
				//cetak Menu Admin
				MenuAdmin()
				//Masukan Input Admin
				SubInput = InputAdmin()
				fmt.Print("\033[H\033[2J")
				if SubInput == 1 {
					fmt.Print("\033[H\033[2J")
					TambahCalonLegislatif(&legislatif, &nLegislatif)
					fmt.Print("\033[H\033[2J")
				} else if SubInput == 2 {
					fmt.Print("\033[H\033[2J")
					HapusCalonLegislatif(&legislatif, &nLegislatif)
				} else if SubInput == 3 {
					fmt.Print("\033[H\033[2J")
					EditCalonlegislatif(&legislatif, nLegislatif)
				} else if SubInput == 4 {
					fmt.Print("\033[H\033[2J")
					HapusSuaraLegislatif(&legislatif, nLegislatif, &pemilu, &nPemilu)
				} else if SubInput == 6 {
					i2 = false
					fmt.Print("\033[H\033[2J")
				} else if SubInput == 5 {
					fmt.Print("\033[H\033[2J")
					Waktu = TenggatWaktu()
				}
			}
		}
	}
}

func DaftarMenu() {
	//Cetak Menu
	fmt.Println("Selamat datang di aplikasi pemilihan umum")
	cetak_waktu()
	fmt.Println("Silahkan masukkan opsi yang ingin anda pilih")
	fmt.Println("1. Memilih Calon Legislatif")
	fmt.Println("2. List Calon Legislatif")
	fmt.Println("3. Pencarian Data Calon Legislatif")
	fmt.Println("4. Menampilkan Pemenang")
	fmt.Println("5. Exit")
}

func CariDataLegislatif() {
	//Cetak Menu Pencarian
	fmt.Println("Silahkan masukkan pencarian yang ingin anda gunakan")
	fmt.Println("1. Pencarian berdasarkan nama")
	fmt.Println("2. Pencarian berdasarkan partai")
	fmt.Println("3. Pencarian berdasarkan pemilih")
}

func TampilanDataLegislatif() {
	//Cetak Menu List
	fmt.Println("Silahkan masukkan pencarian yang ingin anda gunakan")
	fmt.Println("1. menampilkan berdasarkan nama")
	fmt.Println("2. menampilkan berdasarkan partai")
	fmt.Println("3. menampilkan berdasarkan suara")
}

func MenuAdmin() {
	//Cetak Menu Admin
	fmt.Println("Selamat datang di menu Admin")
	cetak_waktu()
	fmt.Println("Silahkan masukkan opsi yang ingin anda pilih")
	fmt.Println("1. Menginput anggota parlemen baru")
	fmt.Println("2. Menghapus anggota parlemen")
	fmt.Println("3. Mengedit anggota parlemen")
	fmt.Println("4. Menghapus data suara parlemen")
	fmt.Println("5. Mengubah rentang waktu pemilihan")
	fmt.Println("6. Kembali ke menu awal")
}

func InputDaftarMenu() int {
	//Membaca input daftar menu
	var x string
	var i bool = true
	for i {
		fmt.Scan(&x)
		if x == "admin" || x == "Admin" {
			return 6
		} else if x == "1" {
			return 1
		} else if x == "2" {
			return 2
		} else if x == "3" {
			return 3
		} else if x == "4" {
			return 4
		} else if x == "5" {
			return 5
		} else {
			DaftarMenu()
			fmt.Println("Input tidak benar, silahkan coba lagi")
		}
	}
	return -1
}

func InputAdmin() int {
	//membaca Input Admin
	var n int
	var i bool = true
	for i {
		fmt.Scan(&n)
		if n <= 6 && n >= 1 {
			return n
		} else {
			MenuAdmin()
			fmt.Println("Input tidak benar, silahkan masukkan kembali input yang benar")
		}
	}
	return -1
}

func MemilihCalonLegislatif(DataLegislatif *DaftarCalonLegislatif, n int, DataPemilu *DaftarPemilu, nPemilu *int, waktu bool) {
	/* Fungsi ini digunakan untuk memilih anggota calon legislatif */

	var NamaPemilih, NamaKandidat string
	var i, ada bool
	var j int
	ada = false
	i = waktu
	if !i {
		fmt.Println("Anda tidak sedang berada di masa pemilihan, Silahkan kembali saat masa pemilihan sedang berlangsung.")
	}
	for i {
		ada = false
		for !ada {
			fmt.Println("Silahkan masukkan pilihan anda dengan format 'Nama_Pemilih Nama_Kandidat_Yang_Dipilih'. jika sudah selesai memilih, silahkan ketik Selesai")
			fmt.Scan(&NamaPemilih, &NamaKandidat)
			if NamaPemilih == "Selesai" || NamaPemilih == "selesai" || NamaKandidat == "selesai" || NamaKandidat == "Selesai" {
				ada = true
				i = false
			}
			for j = 0; j < n; j++ {
				if DataLegislatif[j].nama == NamaKandidat {
					ada = true
					DataLegislatif[j].suara++
					TambahDataPemilu(DataPemilu, NamaPemilih, NamaKandidat, nPemilu)
				}
			}

			if !ada {
				fmt.Println("Kandidat yang anda pilih tidak ada, silahkan coba lagi")
			}
		}
	}
}

func TambahDataPemilu(DataPemilu *DaftarPemilu, pemilih, kandidat string, n *int) {
	/* Fugsi ini digunakan untuk menambahkan suara kepada calon legislatif.
	DataPemilu adalah daftar Pemilihan umum dengan nilai awal n.
	Pemilih dan Kandidat yang berisi nama pemilih dan kandidat yang dipilih.

	n bertambah. Data yang baru di input akan masuk ke DataPemilu dengan format pemilih dan kandidat.
	*/
	var size = len(*DataPemilu)
	for i := 0; i < size; i++ {
		if DataPemilu[i].nama == "" {
			DataPemilu[i].nama = pemilih
			DataPemilu[i].kandidat = kandidat
			*n++
			i = size + 10
		}
	}
}

func ProsesMenampilkanDataLegislatif() int {
	//Membaca Input untuk menampilkan data legislatif
	var n int
	var i bool = true
	for i {
		fmt.Scan(&n)
		if n >= 1 && n <= 3 {
			return n
		} else {
			TampilanDataLegislatif()
			fmt.Println("Input tidak benar, silahkan masukkan kembali input yang benar")
		}
	}
	return -1
}

func DataBerdasarkanNama(DataLegislatif *DaftarCalonLegislatif, n int) {
	//Menampilkan Data legislatif dalam urutan naik berdasarkan nama
	var i int
	sort_nama(DataLegislatif, n)
	fmt.Printf("%-10s %-18s %-10s\n", "NAMA", "PARTAI", "SUARA")
	for i = 0; i < n; i++ {
		fmt.Printf("%-10s %-18s %-10d\n", DataLegislatif[i].nama, DataLegislatif[i].partai, DataLegislatif[i].suara)
	}
}

func DataBerdasarkanPartai(DataLegislatif *DaftarCalonLegislatif, n int) {
	//Menampilkan Data legislatif dalam urutan naik berdasarkan Partai
	var i int
	sort_partai(DataLegislatif, n)
	fmt.Printf("%-10s %-18s %-10s\n", "NAMA", "PARTAI", "SUARA")
	for i = 0; i < n; i++ {
		fmt.Printf("%-10s %-18s %-10d\n", DataLegislatif[i].nama, DataLegislatif[i].partai, DataLegislatif[i].suara)
	}
}

func DataBerdasarkanSuara(DataLegislatif *DaftarCalonLegislatif, n int) {
	//Menampilkan Data legislatif dalam urutan naik berdasarkan Suara
	var i int
	sort_suara(DataLegislatif, n)
	fmt.Printf("%-10s %-18s %-10s\n", "NAMA", "PARTAI", "SUARA")

	for i = 0; i < n; i++ {
		fmt.Printf("%-10s %-18s %-10d\n", DataLegislatif[i].nama, DataLegislatif[i].partai, DataLegislatif[i].suara)
	}
}

func sort_nama(DataLegislatif *DaftarCalonLegislatif, n int) {
	// Fungsi Selection sort mengurutkan berdasarkan nama
	var idx int
	var i, j int
	for i = 0; i < n-1; i++ {
		idx = i
		for j = i + 1; j < n; j++ {
			if DataLegislatif[j].nama < DataLegislatif[idx].nama {
				idx = j
			}
		}
		if idx != i {
			tukar_string(&DataLegislatif[i].nama, &DataLegislatif[idx].nama)
			tukar_string(&DataLegislatif[i].partai, &DataLegislatif[idx].partai)
			tukar_int(&DataLegislatif[i].suara, &DataLegislatif[idx].suara)
		}
	}
}

func sort_partai(DataLegislatif *DaftarCalonLegislatif, n int) {
	// Fungsi Selection sort mengurutkan berdasarkan partai
	var idx int
	var i, j int
	for i = 0; i < n-1; i++ {
		idx = i
		for j = i + 1; j < n; j++ {
			if DataLegislatif[j].partai < DataLegislatif[idx].partai {
				idx = j
			}
		}
		if idx != i {
			tukar_string(&DataLegislatif[i].nama, &DataLegislatif[idx].nama)
			tukar_string(&DataLegislatif[i].partai, &DataLegislatif[idx].partai)
			tukar_int(&DataLegislatif[i].suara, &DataLegislatif[idx].suara)
		}
	}
}

func sort_suara(DataLegislatif *DaftarCalonLegislatif, n int) {
	// Fungsi Selection sort mengurutkan berdasarkan suara
	var idx int
	var i, j int
	for i = 0; i < n-1; i++ {
		idx = i
		for j = i + 1; j < n; j++ {
			if DataLegislatif[j].suara > DataLegislatif[idx].suara {
				idx = j
			}
		}
		if idx != i {
			tukar_string(&DataLegislatif[i].nama, &DataLegislatif[idx].nama)
			tukar_string(&DataLegislatif[i].partai, &DataLegislatif[idx].partai)
			tukar_int(&DataLegislatif[i].suara, &DataLegislatif[idx].suara)
		}
	}
}

func tukar_string(s1, s2 *string) {
	//menukar string
	temp1 := *s1
	*s1 = *s2
	*s2 = temp1
}

func tukar_int(s1, s2 *int) {
	//menukar integer
	temp2 := *s1
	*s1 = *s2
	*s2 = temp2
}

func ProsesCariDataLegislatif() int {
	//Proses input Pencarian Data legislatif
	var n int
	var i bool = true
	for i {
		fmt.Scan(&n)
		if n >= 1 && n <= 3 {
			return n
		} else {
			fmt.Print("\033[H\033[2J")
			CariDataLegislatif()
			fmt.Println("Input tidak benar, silahkan masukkan kembali input yang benar")
		}
	}
	return -1
}

func CariBerdasarkanNama(DataLegislatif *DaftarCalonLegislatif, n int) {
	//Fungsi ini digunakan untuk mencari berdasarkan nama daftar calon legislatif, jika ada maka akan di tampilkan.
	var NamaPencarian string
	var i int
	var ada bool = false
	var idx int = -1
	fmt.Println("Silahkan masukkan nama yang ingin anda cari")
	fmt.Scan(&NamaPencarian)
	for i = 0; i < n; i++ {
		if DataLegislatif[i].nama == NamaPencarian {
			ada = true
			idx = i
		}
	}
	if ada {
		fmt.Println("Berikut data yang anda cari:")
		fmt.Println(DataLegislatif[idx].nama, DataLegislatif[idx].partai, DataLegislatif[idx].suara)
	} else {
		fmt.Println("Mohon maaf, data yang anda cari tidak dapat ditemukan")
	}
}

func CariBerdasarkanPartai(DataLegislatif *DaftarCalonLegislatif, n int) {
	//Fungsi ini digunakan untuk mencari berdasarkan partai daftar calon legislatif, jika ada maka akan di tampilkan.

	var NamaPencarian string
	var ada bool = false
	var i int
	sort_partai(DataLegislatif, n)
	fmt.Println("Silahkan masukkan partai yang ingin anda cari")
	fmt.Scan(&NamaPencarian)
	for i = 0; i < n; i++ {
		if NamaPencarian == DataLegislatif[i].partai {
			ada = true
			fmt.Println(DataLegislatif[i].nama, DataLegislatif[i].partai, DataLegislatif[i].suara)
		}
	}
	if !ada {
		fmt.Println("Maaf, partai yang anda cari tidak ada")
	}

}

func CariBerdasarkanPemilih(DataLegislatif *DaftarCalonLegislatif, DataPemilu DaftarPemilu, n, nPemilu int) {
	//Fungsi ini digunakan untuk mencari berdasarkan Pemilih daftar calon legislatif, jika ada maka akan di tampilkan.
	var NamaPencarian string
	var ada bool = false
	var i, j int
	fmt.Scan("Silahkan masukkan nama pemilih yang ingin anda cari")
	fmt.Scan(&NamaPencarian)
	for i = 0; i < nPemilu; i++ {
		if NamaPencarian == DataPemilu[i].nama {
			ada = true
			for j = 0; j < n; j++ {
				if DataPemilu[i].kandidat == DataLegislatif[j].nama {
					fmt.Println("Nama anggota yang dipilih:", DataPemilu[i].kandidat)
					fmt.Println(DataLegislatif[j].nama, DataLegislatif[j].partai, DataLegislatif[j].suara)
				}
			}
		}
	}
	if !ada {
		fmt.Println("Nama yang anda pilih tidak dapat ditemukan")
	}
}

func ThresholdKandidat(DataLegislatif *DaftarCalonLegislatif, n, nPemilu int) {
	//Fungsi ini digunakan untuk menghitung suara terbanyak dan menampilkan pemenang
	var rerata float64
	var i int
	sort_nama(DataLegislatif, n)
	rerata = float64(nPemilu) / float64(n)
	fmt.Println("Pemilih yang lolos untuk sebagai anggota legislatif adalah:")
	for i = 0; i < n; i++ {
		if rerata < float64(DataLegislatif[i].suara) {
			fmt.Printf("%s dari partai %s dengan perolehan suara %d\n", DataLegislatif[i].nama, DataLegislatif[i].partai, DataLegislatif[i].suara)
		}
	}
}

func TambahCalonLegislatif(DataLegislatif *DaftarCalonLegislatif, x *int) {
	/*fungsi ini digunakan untuk menambahkan Calon Legislatif.
	Data legislatif adalah Daftar Calon Legislatif dan x adalah nilai awalnya.
	x akan bertambah sesuai dengan yang di inputkan :)
	*/
	var n, i int
	fmt.Println("Silahkan masukkan jumlah anggota parlemen yang ingin ditambahkan terlebih dahulu")
	fmt.Scan(&n)

	fmt.Println("Silahkan masukkan data anggota parlemen dengan format 'nama nama_partai'")

	for i = *x; i < *x+n; i++ {
		fmt.Scan(&DataLegislatif[i].nama, &DataLegislatif[i].partai)
	}
	*x += n
}

func EditCalonlegislatif(DataLegislatif *DaftarCalonLegislatif, x int) {
	/* Fungsi ini digunakan untuk mengedit Salah satu kandidat calon legislatif
	yang sudah tersimpan di DataLegislatif */
	var ada bool = false
	var idx int = -1
	var nama string
	var i int
	fmt.Println("Silahkan ketikkan nama anggota parlemen yang ingin anda edit")
	fmt.Scan(&nama)
	for i = 0; i < x; i++ {
		if DataLegislatif[i].nama == nama {
			ada = true
			idx = i
		}
	}
	if ada {
		fmt.Println("Silahkan masukkan perubahan yang ingin anda ubah denan format 'nama nama_partai")
		fmt.Scan(&DataLegislatif[idx].nama, &DataLegislatif[idx].partai)
	} else {
		fmt.Println("Nama anggota yang anda cari tidak ditemukan!")
	}
}

func HapusCalonLegislatif(DataLegislatif *DaftarCalonLegislatif, n *int) {
	/*Fungsi ini digunakan untuk menghapus salah satu kandidat calon Legislatif
	yang sudah tesimpan di DataLegislatif dan nilai n akan berkurang jika berhasil menghapus */
	var nama string
	var ada bool = false
	var idx int = -1
	var i int
	fmt.Println("Silahkan masukkan nama anggota parlemen yang ingin anda hapus")
	fmt.Scan(&nama)
	for i = 0; i < *n; i++ {
		if DataLegislatif[i].nama == nama {
			ada = true
			idx = i
		}
	}

	if ada {
		for i := idx; i < *n-1; i++ {
			DataLegislatif[i] = DataLegislatif[i+1]
		}
		*n--
	}
}

func HapusSuaraLegislatif(DataLegislatif *DaftarCalonLegislatif, n int, DataPemilu *DaftarPemilu, nPemilu *int) {
	/*Fungsi ini digunakan untuk menghapus suara salah satu pemilih yang sudah memilih kandidat calon Legislatif
	yang sudah tesimpan di DataLegislatif. Jika Penghapusan berhasil maka suara berkurang 1. */
	var nama string
	var ada bool = false
	var idx int = -1
	var Kandidat string
	fmt.Println("Silahkan masukkan nama pemilih yang ingin anda hapus")
	fmt.Scan(&nama)

	for i := 0; i < *nPemilu; i++ {
		if DataPemilu[i].nama == nama {
			ada = true
			idx = i
			Kandidat = DataPemilu[i].kandidat
		}
	}
	if ada {
		for i := idx; i < *nPemilu-1; i++ {
			DataPemilu[i] = DataPemilu[i+1]
		}
		*nPemilu--
		for i := 0; i < n; i++ {
			if DataLegislatif[i].nama == Kandidat {
				DataLegislatif[i].suara--
			}
		}

	}
}

func cetak_waktu() {
	// Mendapatkan tanggal dan waktu saat ini
	now := time.Now()

	// Mendapatkan tanggal
	tahun, bulan, hari := now.Date()
	fmt.Printf("Tanggal hari ini: %02d/%02d/%d\n", hari, bulan, tahun)
}

func Batas_Waktu() string {
	var waktu_string string
	//menginput waktu dalam bentuk string
	fmt.Scan(&waktu_string)
	return waktu_string
}

func TenggatWaktu() bool {
	var rentang1, rentang2 string

	//Menginput Rentang Waktu Awal Pemilu
	fmt.Println("masukkan rentang waktu awal dengan format dd/mm/yyyy")
	rentang1 = Batas_Waktu()

	/* menggunakan module time untuk mengubah format penanggalan menjadi bertipe time.time (time.Parse)
	NOTE :
	Tipe data time.Time digunakan untuk merepresentasikan waktu dalam Go. Ini adalah tipe data yang memiliki metode dan properti
	untuk memanipulasi dan memformat waktu.
	*/
	cek_start, err := time.Parse("02/01/2006", rentang1)
	/*mengecek error dalam input
	nil digunakan untuk menunjukan bahwa tidak ada kesalahan yang terjadi dalam parsing tanggal*/
	if err != nil {
		fmt.Println("Format tanggal awal tidak valid:", err)
	}

	fmt.Println("masukkan rentang waktu akhir dd/mm/yyyy")
	rentang2 = Batas_Waktu()
	cek_akhir, err := time.Parse("02/01/2006", rentang2)
	if err != nil {
		fmt.Println("Format tanggal akhir tidak valid:", err)
	}

	//mengecek tanggal hari ini
	hari_ini := time.Now()
	//me return nilai true atau false sesuai dengan rentang waktu
	return hari_ini.After(cek_start) && hari_ini.Before(cek_akhir)
}
