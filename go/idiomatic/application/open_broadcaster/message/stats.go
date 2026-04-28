package message

type GetStatsRequest struct {
}

type GetStatsResponse struct {
	CpuUsage    float64 `json:"cpu_usage"`
	MemoryUsage float64 `json:"memory_usage"`
}
