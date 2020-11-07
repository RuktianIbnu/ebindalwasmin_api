package model

import "time"

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
}

// KategoriPNBP ...
type KategoriParent1PNBP struct {
	ID                  int64                 `json:"id"`
	NamaLayanan         string                `json:"nama_layanan"`
	Satuan              string                `json:"satuan"`
	Parent              int64                 `json:"parent,omitempty"`
	TarifPNBP           int64                 `json:"tarif_pnbp,omitempty"`
	CreatedAt           *time.Time            `json:"created_at,omitempty"`
	UpdatedAt           *time.Time            `json:"updated_at,omitempty"`
	Status              int64                 `json:"status"`
	KategoriParent2PNBP []KategoriParent2PNBP `json:"kategoriParent2PNBP"`
}
type KategoriParent2PNBP struct {
	ID                  int64                 `json:"id"`
	NamaLayanan         string                `json:"nama_layanan"`
	Satuan              string                `json:"satuan"`
	Parent              int64                 `json:"parent,omitempty"`
	TarifPNBP           int64                 `json:"tarif_pnbp,omitempty"`
	CreatedAt           *time.Time            `json:"created_at,omitempty"`
	UpdatedAt           *time.Time            `json:"updated_at,omitempty"`
	Status              int64                 `json:"status"`
	KategoriParent3PNBP []KategoriParent3PNBP `json:"kategoriParent3PNBP"`
}
type KategoriParent3PNBP struct {
	ID          int64      `json:"id"`
	NamaLayanan string     `json:"nama_layanan"`
	Satuan      string     `json:"satuan"`
	Parent      int64      `json:"parent,omitempty"`
	TarifPNBP   int64      `json:"tarif_pnbp,omitempty"`
	CreatedAt   *time.Time `json:"created_at,omitempty"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
	Status      int64      `json:"status"`
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
