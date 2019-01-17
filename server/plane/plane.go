package plane

import (
	"context"
	"io"
	"math"

	"github.com/kevinsnydercodes/protobuf-example/protobuf/compiled/go/plane"
)

// Server implements the Plane service.
type Server struct{}

var (
	count = map[int64]map[int64]uint64{}
)

func minInt64(a int64, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

func maxInt64(a int64, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func getCount(x int64, y int64) uint64 {
	if count[x] == nil {
		count[x] = map[int64]uint64{}
	}
	return count[x][y]
}

func incrementCount(x int64, y int64) {
	if count[x] == nil {
		count[x] = map[int64]uint64{}
	}
	count[x][y]++
}

// GetPoint gets a single point.
// (No streaming)
func (s *Server) GetPoint(ctx context.Context, req *plane.Coordinates) (*plane.Point, error) {
	resp := &plane.Point{
		Coordinates: &plane.Coordinates{
			X: req.X,
			Y: req.Y,
		},
		Count: getCount(req.X, req.Y),
	}
	return resp, nil
}

// PutPoints puts multiple points.
// (Client streaming)
func (s *Server) PutPoints(stream plane.Plane_PutPointsServer) error {
	min := &plane.Coordinates{
		X: math.MaxInt64,
		Y: math.MaxInt64,
	}
	max := &plane.Coordinates{
		X: math.MinInt64,
		Y: math.MinInt64,
	}

	for {
		coordinates, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		incrementCount(coordinates.X, coordinates.Y)
		min.X = minInt64(min.X, coordinates.X)
		min.Y = minInt64(min.Y, coordinates.Y)
		max.X = maxInt64(max.X, coordinates.X)
		max.Y = maxInt64(max.Y, coordinates.Y)
	}

	bounds := &plane.Bounds{
		Min: min,
		Max: max,
	}
	err := stream.SendAndClose(bounds)
	return err
}

// ListPointsByBounds gets multiple points within the given bounds.
// (Server streaming)
func (s *Server) ListPointsByBounds(req *plane.Bounds, stream plane.Plane_ListPointsByBoundsServer) error {
	for x := range count {
		if req.Min.X < x && x < req.Max.X {
			for y := range count[x] {
				if req.Min.Y < y && y < req.Max.Y {
					point := &plane.Point{
						Coordinates: &plane.Coordinates{
							X: x,
							Y: y,
						},
						Count: getCount(x, y),
					}
					err := stream.Send(point)
					if err != nil {
						return err
					}
				}
			}
		}
	}
	return nil
}

// ListPoints gets multiple points matching the given coordinates.
// (Bidirectional streaming)
func (s *Server) ListPoints(stream plane.Plane_ListPointsServer) error {
	for {
		coordinates, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		resp := &plane.Point{
			Coordinates: &plane.Coordinates{
				X: coordinates.X,
				Y: coordinates.Y,
			},
			Count: getCount(coordinates.X, coordinates.Y),
		}
		stream.Send(resp)
	}
	return nil
}
