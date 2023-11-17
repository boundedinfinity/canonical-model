package hardware

import "github.com/boundedinfinity/schema/idiomatic/id"

type CentralProcessingUnit struct {
	Id                    id.Id                        `json:"id,omitempty"`
	CoreCount             int                          `json:"core-count,omitempty"`
	SocketType            string                       `json:"socket-type,omitempty"`
	VirtualRealityReady   bool                         `json:"virtual-reality-ready,omitempty"`
	IntegratedGraphics    bool                         `json:"integrated-graphics,omitempty"`
	HyperThreadingSupport bool                         `json:"hyper-threading-support,omitempty"`
	Virtualization        bool                         `json:"virtualization,omitempty"`
	Cache                 []CentralProcessingUnitCache `json:"cache,omitempty"`
	Memory                CentralProcessingUnitMemory  `json:"memory,omitempty"`
	ThermalPower          string                       `json:"thermal-power,omitempty"`
}

type CentralProcessingUnitCache struct {
	Level int `json:"level,omitempty"`
	Size  int `json:"size,omitempty"`
}

type CentralProcessingUnitMemory struct {
	Type int `json:"type,omitempty"`
	Size int `json:"size,omitempty"`
}
