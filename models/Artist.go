package models

type ArtistType string

const (
	ARTIST_GROUP ArtistType = "Band"
	ARTIST_SOLO             = "Solo"
)

type Artist interface {
	Id() int
	IsPartOf(Artist) bool
	Name() string
	Type() ArtistType
}

type SoloArtist struct {
	artistType ArtistType
	id         int
	name       string
}

func (a SoloArtist) Id() int {
	return a.id
}

func (a SoloArtist) Name() string {
	return a.name
}

func (a SoloArtist) Type() ArtistType {
	return ARTIST_SOLO
}

func (a SoloArtist) IsPartOf(o Artist) bool {
	return (a.Id() == o.Id() && a.Name() == o.Name() && a.Type() == o.Type())
}

var SoloArtistId = 1

func NewSoloArtist(name string) *SoloArtist {
	ret := &SoloArtist{
		name:       name,
		id:         SoloArtistId, // To be replaced by getNextId(ARTIST_SOLO) from the db driver
		artistType: ARTIST_SOLO,
	}
	SoloArtistId++
	return ret
}

type Band struct {
	artistType ArtistType
	id         int
	name       string
	members    []Artist
}

func (a Band) Id() int {
	return a.id
}

func (a Band) Name() string {
	return a.name
}

func (a Band) Type() ArtistType {
	return ARTIST_GROUP
}

func (a Band) IsPartOf(o Artist) bool {
	for _, member := range a.members {
		if member.IsPartOf(o) {
			return true
		}
	}
	return false
}

var GroupId = 1

func NewGroup(name string) *SoloArtist {
	ret := &SoloArtist{
		name:       name,
		id:         GroupId, // To be replaced by getNextId(ARTIST_GROUP) from the db driver
		artistType: ARTIST_GROUP,
	}
	GroupId++
	return ret
}
