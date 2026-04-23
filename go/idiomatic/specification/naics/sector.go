package naics

// https://www.census.gov/naics/reference_files_tools/2022_NAICS_Manual.pdf

type Sector struct {
	Name        string    `json:"name"`
	Code        string    `json:"code"`
	Description string    `json:"description"`
	Parent      *Sector   `json:"parent"`
	SubSectors  []*Sector `json:"sub-sectors"`
}
