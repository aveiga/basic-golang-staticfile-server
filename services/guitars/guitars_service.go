package guitars

import "github.com/aveiga/basic-golang-staticfile-server/domain/guitars"

func CreateGuitar(guitar guitars.Guitar) (*guitars.Guitar, error) {
	return &guitar, nil
}