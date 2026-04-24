package hardware

type Contact struct {
	Id           ider.Id         `json:"id,omitempty"`
	SerialNumber id.SerialNumber `json:"serial-number,omitempty"`
}
