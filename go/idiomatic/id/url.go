package id

type Url struct {
	Id   Id     `json:"id,omitempty"`
	Host string `json:"host,omitempty"`
	Port int    `json:"port,omitempty"`
	Path int    `json:"path,omitempty"`
}
