package message

type Descriminator struct {
	RequestType Kind   `json:"requestType"`
	RequestId   string `json:"requestId,omitempty"`
}
