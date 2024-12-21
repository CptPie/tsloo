package models

type SongType string

const (
	SONGTYPE_ORIGINAL SongType = "Original"
	SONGTYPE_COVER             = "Cover"
)

type Song struct {
	Artists  []Artist
	Duration int
	Id       int
	Path     string // part to binary song - flac or mp3
	SongType SongType
	TitleEn  string
	TitleJP  string
}

func (s Song) IsBy(a Artist) bool {
	for _, artist := range s.Artists {
		if artist.IsPartOf(a) {
			return true
		}
	}
	return false
}

var SongId = 1

func NewSong(titleEn string, titleJP string, songType SongType, duration int, path string, artists ...Artist) *Song {
	ret := &Song{
		Artists:  artists,
		Duration: duration,
		Id:       SongId,
		Path:     path,
		SongType: songType,
		TitleEn:  titleEn,
		TitleJP:  titleJP,
	}
	SongId++
	return ret
}
