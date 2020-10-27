package model

// PasporData ...
type PasporData []struct {
	ID             int    `JSON:"id"`
	IDJenis        int    `JSON:"id_jenis"`
	IDUser         int    `JSON:"id_user"`
	IDKantor       int    `JSON:"Iid_kantord"`
	Tanggal        string `JSON:"tanggal"`
	Laki           string `JSON:"laki"`
	Perempuan      string `JSON:"perempuan"`
	Total          int    `JSON:"total"`
	IDWilayahKerja int    `JSON:"Id_wilayah_kerja"`
}

// PasporDataDetail ...
type PasporDataDetail struct {
	ID               int    `JSON:"id"`
	IDJenis          int    `JSON:"id_jenis"`
	IDUser           int    `JSON:"id_user"`
	NamaUser         string `JSON:"name"`
	IDKantor         int    `JSON:"id_kantor"`
	NamaKantor       string `JSON:"nama_kantor"`
	Tanggal          string `JSON:"tanggal"`
	Laki             int    `JSON:"laki"`
	Perempuan        int    `JSON:"perempuan"`
	Total            int    `JSON:"total"`
	IDWilayahKerja   int    `JSON:"id_wilayah_kerja"`
	NamaWilayahKerja string `JSON:"nama_wilayah_kerja"`
}

// StatusResPaspor ...
type StatusResPaspor struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
	Data   []PasporDataDetail
}
