package models

type GettyManifest struct {
	Description []string        `json:"description"`
	Label       string          `json:"label"`
	Related     string          `json:"related"`
	Metadata    []GettyMetadata `json:"metadata"`
	Thumbnail   GettyThumbnail  `json:"thumbnail"`
}

type GettyMetadata struct {
	Label string      `json:"label"`
	Value interface{} `json:"value"`
}

type GettyThumbnail struct {
	ID string `json:"@id"`
}

func (gm *GettyManifest) ExtractFields() (artist string, title string, accession string) {
	for _, meta := range gm.Metadata {
		switch meta.Label {
		case "Artist/Maker":
			artist = extractValue(meta.Value)
		case "Title":
			title = extractValue(meta.Value)
		case "Accession Number":
			accession = extractValue(meta.Value)
		}
	}
	return
}

func extractValue(v any) string {
	switch vv := v.(type) {
	case string:
		return vv
	case []interface{}:
		if len(vv) > 0 {
			if s, ok := vv[0].(string); ok {
				return s
			}
		}
	}
	return ""
}
