package plane

import (
	"context"
	"io"

	"github.com/kevinsnydercodes/protobuf-example/protobuf/compiled/go/plane"
	"google.golang.org/grpc"
)

const (
	host = "localhost:8080"
)

func getClient() (plane.PlaneClient, error) {
	connection, err := grpc.Dial(host, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	return plane.NewPlaneClient(connection), nil
}

// GetPoint handles calling the gRPC GetPoint method.
// (No streaming)
func GetPoint(coordinates *plane.Coordinates) (*plane.Point, error) {
	client, err := getClient()
	if err != nil {
		return nil, err
	}

	return client.GetPoint(context.Background(), coordinates)
}

// PutPoints handles calling the gRPC PutPoints method.
// (Client streaming)
func PutPoints(coordinatesSlice []*plane.Coordinates) (*plane.Bounds, error) {
	client, err := getClient()
	if err != nil {
		return nil, err
	}

	stream, err := client.PutPoints(context.Background())
	if err != nil {
		return nil, err
	}

	for _, coordinates := range coordinatesSlice {
		err := stream.Send(coordinates)
		if err != nil {
			return nil, err
		}
	}

	bounds, err := stream.CloseAndRecv()
	if err != nil {
		return nil, err
	}

	return bounds, nil
}

// ListPointsByBounds handles calling the gRPC ListPointsByBounds method.
// (Server streaming)
func ListPointsByBounds(bounds *plane.Bounds) ([]*plane.Point, error) {
	client, err := getClient()
	if err != nil {
		return nil, err
	}

	stream, err := client.ListPointsByBounds(context.Background(), bounds)
	if err != nil {
		return nil, err
	}

	points := []*plane.Point{}
	for {
		point, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		points = append(points, point)
	}

	return points, nil
}

// ListPoints handles calling the gRPC ListPoints method.
// (Bidirectional streaming)
func ListPoints(coordinatesSlice []*plane.Coordinates) ([]*plane.Point, error) {
	client, err := getClient()
	if err != nil {
		return nil, err
	}

	stream, err := client.ListPoints(context.Background())
	if err != nil {
		return nil, err
	}

	points := []*plane.Point{}
	done := make(chan struct{})
	var recvErr error
	go func() {
		for {
			point, err := stream.Recv()
			if err == io.EOF {
				close(done)
				break
			}
			if err != nil {
				recvErr = err
				close(done)
				break
			}

			points = append(points, point)
		}
	}()

	for _, coordinates := range coordinatesSlice {
		err := stream.Send(coordinates)
		if err != nil {
			return nil, err
		}
	}
	err = stream.CloseSend()
	if err != nil {
		return nil, err
	}

	<-done
	if recvErr != nil {
		return nil, recvErr
	}

	return points, nil
}
