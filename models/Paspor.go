package models

type PasporData struct {
	Id 			int `JSON:"Id"`
	Id_jenis 	int `JSON:"Id"`
	Id_user 	int `JSON:"Id"`
	Id_kantor 	int `JSON:"Id"`
	Tanggal 	string `JSON:"Id"`
	Laki 		string `JSON:"Id"`
	Perempuan 	string `JSON:"Id"`
	Total 		int `JSON:"Id"`
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
}

type StatusResPaspor struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
	Data []PasporDataDetail
}