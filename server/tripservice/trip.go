package trip

import (
	"context"
	trippb "coolcar/proto/gen/go"
)

// type TripServiceServer interface {
// 	GetTrip(context.Context, *GetTripRequest) (*GetTripResponse, error)
// 	mustEmbedUnimplementedTripServiceServer()
// }

type Service struct {
	trippb.UnimplementedTripServiceServer
}

func (*Service) GetTrip(c context.Context, req *trippb.GetTripRequest) (*trippb.GetTripResponse, error) {
	return &trippb.GetTripResponse{
		Id: req.Id,
		Trip: &trippb.Trip{
			Start:       "abc",
			End:         "def",
			DurationSec: 3600,
			FeeCent:     10000,
			Status:      trippb.TripStatus_IN_PROGRESS,
			Location: &trippb.Location{
				Lat: 102,
				Lng: 103,
			},
		},
	}, nil
}
