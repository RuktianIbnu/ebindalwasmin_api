package model

import "time"

// Satker ...
type Satker struct {
	IDKantor    int64  `json:"id_kantor"`
	Nama_Kantor string `json:"nama_kantor"`
}

// Report monthyear ...
type ReportMonthYear struct {
	Periode   string     `json:"periode,omitempty"`
	Tanggal   *time.Time `json:"tanggal,omitempty"`
	JenisPnbp string     `json:"jenis_pnbp,omitempty"`
	Total     int64      `json:"total"`
}

// User ...
type User struct {
	ID               int64      `json:"id"`
	Type             string     `json:"type"`
	Name             string     `json:"name"`
	Email            string     `json:"email"`
	EmailVerifiedAt  *time.Time `json:"email_verified_at"`
	Password         string     `json:"password"`
	PasswordChangeAt *time.Time `json:"password_change_at"`
	Active           bool       `json:"active"`
	Timezone         string     `json:"timezone"`
	LastLoginAt      *time.Time `json:"last_login_at"`
	LastLoginIP      string     `json:"last_login_ip"`
	ToBeLoggedOut    bool       `json:"to_be_logged_out"`
	Provider         string     `json:"provider"`
	ProviderID       string     `json:"provider_id"`
	RememberToken    string     `json:"remember_token"`
	CreatedAt        *time.Time `json:"created_at"`
	UpdatedAt        *time.Time `json:"updated_at"`
	DeletedAt        *time.Time `json:"deleted_at"`
	IDKantor         int64      `json:"id_kantor"`
	NamaKantor       string     `json:"nama_kantor"`
}

// KategoriPNBP ...
type KategoriParentPNBP struct {
	ID          int64       `json:"id"`
	NamaLayanan string      `json:"nama_layanan"`
	Satuan      string      `json:"satuan"`
	Parent      int64       `json:"parent,omitempty"`
	TarifPNBP   int64       `json:"tarif_pnbp,omitempty"`
	CreatedAt   *time.Time  `json:"created_at,omitempty"`
	UpdatedAt   *time.Time  `json:"updated_at,omitempty"`
	Status      int64       `json:"status"`
	Child       interface{} `json:"child,omitempty"`
}

// Paspor ...
type Paspor struct {
	ID             int64      `json:"id"`
	IDJenis        int64      `json:"id_jenis"`
	IDUser         int64      `json:"id_user,omitempty"`
	IDKantor       int64      `json:"id_kantor,omitempty"`
	Tanggal        *time.Time `json:"tanggal"`
	Laki           int64      `json:"laki"`
	Perempuan      int64      `json:"perempuan"`
	Total          int64      `json:"total"`
	IDWilayahKerja int64      `json:"id_wilayah_kerja,omitempty"`
}

// Paspor Pivot Perwilayah
type PasporPivotPerwilayah struct {
	Periode          string `json:"periode"`
	KabBangka        int64  `json:"kab_bangka"`
	KabBangkaBarat   int64  `json:"kab_bangka_barat"`
	KabBangkaSelatan int64  `json:"kab_bangka_selatan"`
	KabBangkaTengah  int64  `json:"kab_bangka_tengah"`
	PangkalPinang    int64  `json:"pangkal_pinang"`
	KabBelitungTimur int64  `json:"kab_belitung_timur"`
	KabBelitung      int64  `json:"kab_belitung"`
	KabLainnya1      int64  `json:"kab_lainnya_1"`
	KabLainnya2      int64  `json:"kab_lainnya_2"`
}

type GetPivotPerwilayah struct {
	Tahun   int64  `json:"tahun"`
	Wilayah string `json:"wilayah"`
	Total   int64  `json:"total"`
}

// Paspor Permohonan perkelamin per 10 hari...
type PasporPermohonanperKelaminPer10hari struct {
	Paspor         int64      `json:"paspor"`
	Laki           int64      `json:"laki"`
	Perempuan      int64      `json:"perempuan"`
	Tanggal        *time.Time `json:"tanggal"`
	IDWilayahKerja int64      `json:"id_wilayah_kerja,omitempty"`
	IDKantor       int64      `json:"id_kantor,omitempty"`
}

// Visa ...
type Visa struct {
	ID             int64      `json:"id"`
	IDJenis        int64      `json:"id_jenis"`
	IDUser         int64      `json:"id_user,omitempty"`
	IDKantor       int64      `json:"id_kantor,omitempty"`
	Tanggal        *time.Time `json:"tanggal"`
	Laki           int64      `json:"laki"`
	Perempuan      int64      `json:"perempuan"`
	Total          int64      `json:"total"`
	IDWilayahKerja int64      `json:"id_wilayah_kerja,omitempty"`
}

// Intal ...
type Intal struct {
	ID             int64      `json:"id"`
	IDJenis        int64      `json:"id_jenis"`
	IDUser         int64      `json:"id_user,omitempty"`
	IDKantor       int64      `json:"id_kantor,omitempty"`
	Tanggal        *time.Time `json:"tanggal"`
	Laki           int64      `json:"laki"`
	Perempuan      int64      `json:"perempuan"`
	Total          int64      `json:"total"`
	IDWilayahKerja int64      `json:"id_wilayah_kerja,omitempty"`
}

// Intal Permohonan perkelamin per 10 hari...
type IntalPermohonanperKelaminPer10hari struct {
	IzinTinggal    int64      `json:"izintinggal"`
	Laki           int64      `json:"laki"`
	Perempuan      int64      `json:"perempuan"`
	Tanggal        *time.Time `json:"tanggal"`
	IDWilayahKerja int64      `json:"id_wilayah_kerja,omitempty"`
	IDKantor       int64      `json:"id_kantor,omitempty"`
}

// Pnbp ...
type Pnbp struct {
	ID             int64      `json:"id"`
	IDJenis        int64      `json:"id_jenis"`
	IDUser         int64      `json:"id_user,omitempty"`
	IDKantor       int64      `json:"id_kantor,omitempty"`
	Tanggal        *time.Time `json:"tanggal"`
	Laki           int64      `json:"laki"`
	Perempuan      int64      `json:"perempuan"`
	Total          int64      `json:"total"`
	IDWilayahKerja int64      `json:"id_wilayah_kerja,omitempty"`
}

// PNBP Semua Kategori perbulantahun
type PnbpAllKategoriPerbulanTahun struct {
	Periode     string `json:"periode"`
	Visa        int64  `json:"visa"`
	Paspor      int64  `json:"paspor"`
	IzinTinggal int64  `json:"izin_tinggal"`
	PnbpLainnya int64  `json:"pnbp_lainnya"`
}

// PNBP get total PNBP
type PnbpGetTotalPnbp struct {
	Periode     string `json:"periode"`
	Visa        int64  `json:"visa"`
	Paspor      int64  `json:"paspor"`
	IzinTinggal int64  `json:"izin_tinggal"`
	PnbpLainnya int64  `json:"pnbp_lainnya"`
	Total       int64  `json:"total"`
}

type PNBPPermohonanperKelaminPer10hari struct {
	Pnbp           int64      `json:"pnbp"`
	Laki           int64      `json:"laki"`
	Perempuan      int64      `json:"perempuan"`
	Tanggal        *time.Time `json:"tanggal"`
	IDWilayahKerja int64      `json:"id_wilayah_kerja,omitempty"`
	IDKantor       int64      `json:"id_kantor,omitempty"`
}

type GeneralPnbp struct {
	NamaKantor string `json:"nama_kantor"`
	Periode    string `json:"periode"`
	Total      int64  `json:"total"`
}
