package types

import (
	"zz/service/db/models"
)

type StationAndFavorite struct {
	models.Station `boil:",bind"`
	Favorite       bool `boil:"favorite"`
}
