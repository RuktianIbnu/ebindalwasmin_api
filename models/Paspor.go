package models

type PasporData []struct {
	Id 			int `JSON:"id"`
	Id_jenis 	int `JSON:"id_jenis"`
	Id_user 	int `JSON:"id_user"`
	Id_kantor 	int `JSON:"Iid_kantord"`
	Tanggal 	string `JSON:"tanggal"`
	Laki 		string `JSON:"laki"`
	Perempuan 	string `JSON:"perempuan"`
	Total 		int `JSON:"total"`
	Id_wilayah_kerja 		int `JSON:"Id_wilayah_kerja"`
}

type PasporDataDetail struct {
	Id 			int `JSON:"id"`
	Id_jenis 	int `JSON:"id_jenis"`
	Id_user 	int `JSON:"id_user"`
	Nama_user 	string `JSON:"name"`
	Id_kantor 	int `JSON:"id_kantor"`
	Nama_kantor string `JSON:"nama_kantor"`
	Tanggal 	string `JSON:"tanggal"`
	Laki 		int `JSON:"laki"`
	Perempuan 	int `JSON:"perempuan"`
	Total 		int `JSON:"total"`
	Id_wilayah_kerja 		int `JSON:"id_wilayah_kerja"`
	Nama_wilayah_kerja	string `JSON:"nama_wilayah_kerja"`
}

type StatusResPaspor struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
	Data []PasporDataDetail
}