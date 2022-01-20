package responses

type ProductDesc struct {
	D struct {
		Product            string `json:"Product"`
		Language           string `json:"Language"`
		ProductDescription string `json:"ProductDescription"`
	} `json:"d"`
}
