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
type KategoriPNBP struct {
	ID          int64      `json:"id"`
	NamaLayanan string     `json:"nama_layanan"`
	Satuan      string     `json:"satuan"`
	Parent      int64      `json:"parent,omitempty"`
	TarifPNBP   int64      `json:"tarif_pnbp,omitempty"`
	CreatedAt   *time.Time `json:"created_at,omitempty"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
	Status      int64      `json:"status"`
}
