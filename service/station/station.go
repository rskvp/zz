package station

import (
	"context"
	"database/sql"

	pb "zz/protos/gen/station"
	"zz/service/db/models"
	"zz/service/station/mappers"

	"zz/service/auth/jwt"
	"zz/service/db"
	"zz/service/logger"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type stationAPIServer struct {
	pb.UnimplementedStationAPIServer
	db  *db.DB
	jwt *jwt.Service
}

func New(db *db.DB, jwt *jwt.Service) *stationAPIServer {
	if db == nil {
		logger.I.Panic("db is nil")
	}
	if jwt == nil {
		logger.I.Panic("jwt is nil")
	}
	return &stationAPIServer{
		db:  db,
		jwt: jwt,
	}
}

func (s *stationAPIServer) GetManyStations(ctx context.Context, req *pb.GetManyStationsRequest) (*pb.GetManyStationsResponse, error) {
	userID, err := s.jwt.Validate(req.GetToken())
	if err != nil {
		logger.I.Error("authentication error", zap.Error(err), zap.Any("req", req))
		return nil, status.Errorf(codes.Unauthenticated, "failed to authenticate: %s", err)
	}

	total, err := s.db.CountStations(ctx, req.GetQuery())
	if err != nil {
		logger.I.Error("count station failed", zap.Error(err), zap.Any("req", req))
		return nil, err
	}
	pageCount := total/req.GetLimit() + 1

	res, err := s.db.FindManyStationAndFavorite(ctx, userID, req.GetQuery(), int(req.GetLimit()), int(req.GetPage()))
	if err == sql.ErrNoRows {
		return &pb.GetManyStationsResponse{
			Stations: &pb.PaginatedStation{
				Data:      []*pb.Station{},
				Count:     0,
				Total:     total,
				Page:      req.GetPage(),
				PageCount: pageCount,
			},
		}, nil
	}

	data := mappers.StationAndFavoritesFromDB(res)

	return &pb.GetManyStationsResponse{
		Stations: &pb.PaginatedStation{
			Data:      data,
			Count:     int64(len(data)),
			Total:     total,
			Page:      req.GetPage(),
			PageCount: pageCount,
		},
	}, nil
}
func (s *stationAPIServer) GetOneStation(ctx context.Context, req *pb.GetOneStationRequest) (*pb.GetOneStationResponse, error) {
	userID, err := s.jwt.Validate(req.GetToken())
	if err != nil {
		logger.I.Error("authentication error", zap.Error(err), zap.Any("req", req))
		return nil, status.Errorf(codes.Unauthenticated, "failed to authenticate: %s", err)
	}

	res, err := s.db.FindOneStationAndFavorite(ctx, req.GetId(), userID)
	if err == sql.ErrNoRows {
		return nil, status.Error(codes.NotFound, "station not found")
	}
	return &pb.GetOneStationResponse{
		Station: mappers.StationAndFavoriteFromDB(res),
	}, nil
}

type favoriteSetter func(ctx context.Context, m *models.Favorite) error

func (s *stationAPIServer) SetFavoriteOneStation(ctx context.Context, req *pb.SetFavoriteOneStationRequest) (*pb.SetFavoriteOneStationResponse, error) {
	userID, err := s.jwt.Validate(req.GetToken())
	if err != nil {
		logger.I.Error("authentication error", zap.Error(err), zap.Any("req", req))
		return nil, status.Errorf(codes.Unauthenticated, "failed to authenticate: %s", err)
	}

	var fn favoriteSetter
	if req.GetValue() {
		fn = s.db.CreateFavorite
	} else {
		fn = s.db.DeleteFavorite
	}

	if err := fn(ctx, &models.Favorite{
		StationID: req.Id,
		UserID:    userID,
	}); err != nil {
		return nil, err
	}

	return &pb.SetFavoriteOneStationResponse{}, nil
}
