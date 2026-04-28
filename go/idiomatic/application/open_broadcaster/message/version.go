package message

import "github.com/boundedinfinity/canonical_model/go/idiomatic/ider"

func RequestVersion() GetVersionRequest {
	return GetVersionRequest{
		RequestType: Kinds.GetVersion,
		RequestId:   ider.New().String(),
	}
}

type GetVersionRequest struct {
	RequestType Kind   `json:"requestType"`
	RequestId   string `json:"requestId,omitempty"`
}

type GetVersionResponse struct {
	RequestType   Kind   `json:"requestType,omitempty"`
	RequestId     string `json:"requestId,omitempty"`
	RequestStatus struct {
		Result  bool   `json:"result"`
		Code    int    `json:"code"`
		Comment string `json:"comment,omitempty"`
	} `json:"requestStatus"`

	ObsVersion            string   `json:"obsVersion"`
	ObsWebSocketVersion   string   `json:"obsWebSocketVersion"`
	RPCVersion            int      `json:"rpcVersion"`
	AvailableRequests     []string `json:"availableRequests"`
	SupportedImageFormats []string `json:"supportedImageFormats"`
	Platform              string   `json:"platform"`
	PlatformDescription   string   `json:"platformDescription"`
}
