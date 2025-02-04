package entity

type Product struct {
	Id          int    `json:"id"`
	Nombre      string `json:"nombre"`
	Descripcion string `json:"descripcion"`
	Precio      int64  `json:"precio"`
	Stock       int64  `json:"stock"`
	IdSku       string `json:"id_sku"`
}
