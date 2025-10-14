package models

type ArticMetadata struct {
	Data   []ArticData `json:"data"`
	Config ArticConfig `json:"config"`
}

type ArticData struct {
	ID               int      `json:"id"`
	Title            string   `json:"title"`
	ShortDescription string   `json:"short_description"`
	ArtistTitles     []string `json:"artist_titles"`
	ImageID          string   `json:"image_id"`
}

type ArticInfo struct {
	LicenseLinks string `json:"license_links,omitempty"`
}

type ArticConfig struct {
	IiifURL string `json:"iiif_url"`
}
