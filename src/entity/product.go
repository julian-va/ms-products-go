package entity

type Product struct {
	Id          int     `json:"id"`
	Nombre      string  `json:"nombre"`
	Descripcion string  `json:"descripcion"`
	Precio      float64 `json:"precio"`
	Stock       int     `json:"stock"`
	IdSku       string  `json:"id_sku"`
}
