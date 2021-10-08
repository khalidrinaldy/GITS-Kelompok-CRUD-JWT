package entities

type transaksi struct {
	id_transaksi        int `gorm:"primary_key, AUTO_INCREMENT"`
	id_transaksi_produk int
	id_supplier         int
	id_admin            int
	total_harga         float64
	tanggal_transaksi   date
}
