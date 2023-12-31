package mappers

import (
	"encoding/json"

	pb "zz/protos/gen/station"
	"zz/service/db/types"

	"zz/service/logger"

	"go.uber.org/zap"
)

func stringToCoordinates(in string) []float64 {
	var out []float64
	if err := json.Unmarshal([]byte(in), &out); err != nil {
		logger.I.Panic("failed to convert", zap.Error(err))
	}
	return out
}

func stringToGeo(in string) *pb.Geometry {
	var out pb.Geometry
	if err := json.Unmarshal([]byte(in), &out); err != nil {
		logger.I.Panic("failed to convert", zap.Error(err))
	}
	return &out
}

func StationAndFavoritesFromDB(ss []*types.StationAndFavorite) []*pb.Station {
	rr := make([]*pb.Station, 0, len(ss))
	for _, s := range ss {
		rr = append(rr, StationAndFavoriteFromDB(s))
	}
	return rr
}

func StationAndFavoriteFromDB(s *types.StationAndFavorite) *pb.Station {
	return &pb.Station{
		Id:          s.ID,
		Commune:     s.Commune,
		YWgs84:      s.YWGS84,
		XWgs84:      s.XWGS84,
		Libelle:     s.Libelle,
		Idgaia:      s.Idgaia,
		Voyageurs:   s.Voyageurs,
		GeoPoint_2D: stringToCoordinates(s.GeoPoint2D),
		CodeLigne:   s.CodeLigne,
		XL93:        s.XL93,
		CGeo:        stringToCoordinates(s.CGeo),
		RgTroncon:   s.RGTroncon,
		GeoShape:    stringToGeo(s.GeoShape),
		Pk:          s.PK,
		Idreseau:    s.Idreseau,
		Departemen:  s.Departemen,
		YL93:        s.YL93,
		Fret:        s.Fret,
		IsFavorite:  s.Favorite,
	}
}
