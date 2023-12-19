package id

// https://man7.org/linux/man-pages/man5/machine-id.5.html
type MachineId struct {
	Id   Id     `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}
