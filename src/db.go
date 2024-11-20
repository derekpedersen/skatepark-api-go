package skatepark_api

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	"github.com/derekpedersen/imgur-go/album"
)

var (
	SKATEPARKS     = Skateparks{}
	IMGUR_SVC      album.AlbumService
	last_loaded_at = time.Time{}
)

func ParseSkateparks(
	dir string,
	loadImgurAlbums bool,
) (
	skateparks Skateparks,
	err error,
) {

	if err := filepath.Walk(
		dir,
		func(
			path string,
			info os.FileInfo,
			err error,
		) error {

			if err != nil {
				return err
			}

			if info.IsDir() {
				return nil
			}

			raw, err := ioutil.ReadFile(path)
			if err != nil {
				return err
			}

			m := []Skatepark{}
			if err = json.Unmarshal(raw, &m); err != nil {
				return err
			}

			if loadImgurAlbums {

				for k := range m {

					if m[k].Album == nil || len(m[k].Album.Images) <= 0 {

						m[k].Album, err = IMGUR_SVC.GetAlbum(m[k].AlbumID)

						if err != nil {
							return err
						}

						raw, err := json.Marshal(m)
						if err != nil {
							return err
						}

						if err := ioutil.WriteFile(path, raw, 0644); err != nil {
							return err
						}
					}
				}

				if err := ioutil.WriteFile(path, raw, 0644); err != nil {
					return err
				}
			}

			skateparks = append(skateparks, m...)

			return nil
		},
	); err != nil {
		return nil, err
	}

	SKATEPARKS = skateparks

	return skateparks, nil
}

func GetSkateparks(
	refresh bool,
	loadImgurAlbums bool,
) (
	skateparks Skateparks,
	err error,
) {

	if SKATEPARKS == nil ||
		len(SKATEPARKS) <= 0 ||
		refresh ||
		time.Now().UTC().Sub(last_loaded_at).Hours() > 24 {

		if skateparks, err = ParseSkateparks(os.Getenv("SKATEPARKS_FILE"), loadImgurAlbums); err != nil {
			return nil, err
		}
	}

	return skateparks, nil
}
