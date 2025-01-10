package database

type Database struct{}

type Entry struct {
	Youtube string
	Mp3     string
	Id      int
}

func GetDB() (*Database, error) {
	return &Database{}, nil
}

func (*Database) GetEntry(id int) (*Entry, error) {
	return &Entry{
		Id:      1,
		Youtube: "SZECHjr0iOU",
		Mp3:     "./media/foo.mp3",
	}, nil
}
