package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/kevinsnydercodes/protobuf-example/client/plane"
	planeproto "github.com/kevinsnydercodes/protobuf-example/protobuf/compiled/go/plane"

	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	app = kingpin.New("grpc-demo", "A small CLI to demonstrate gRPC in Go.")

	getpoint            = app.Command("getpoint", "")
	getpointCoordinates = newCoordinatesArg(getpoint)

	putpoints                 = app.Command("putpoints", "")
	putpointsCoordinatesSlice = newCoordinatesSliceArg(putpoints)

	listpointsbybounds                                       = app.Command("listpointsbybounds", "")
	listpointsbyboundsBoundsMin, listpointsbyboundsBoundsMax = newBoundsArgs(listpointsbybounds)

	listpoints                 = app.Command("listpoints", "")
	listpointsCoordinatesSlice = newCoordinatesSliceArg(listpoints)
)

func newCoordinatesArg(cmd *kingpin.CmdClause) *string {
	return cmd.Arg("coordinates", "").Required().String()
}

func newCoordinatesSliceArg(cmd *kingpin.CmdClause) *[]string {
	return cmd.Arg("coordinatesSlice", "").Required().Strings()
}

func newBoundsArgs(cmd *kingpin.CmdClause) (*string, *string) {
	min := cmd.Arg("min", "").Required().String()
	max := cmd.Arg("max", "").Required().String()
	return min, max
}

func parseCoordinates(coordinates *string) (*planeproto.Coordinates, error) {
	slice := strings.Split(*coordinates, ",")
	if len(slice) != 2 {
		return nil, errors.New("coordinates must use format x,y")
	}

	x, err := strconv.Atoi(slice[0])
	if err != nil {
		return nil, errors.New("unable to parse x")
	}
	y, err := strconv.Atoi(slice[1])
	if err != nil {
		return nil, errors.New("unable to parse y")
	}

	res := &planeproto.Coordinates{
		X: int64(x),
		Y: int64(y),
	}
	return res, nil
}

func parseCoordinatesSlice(coordinatesSlice *[]string) ([]*planeproto.Coordinates, error) {
	res := make([]*planeproto.Coordinates, len(*coordinatesSlice))
	for i := 0; i < len(*coordinatesSlice); i++ {
		coordinates, err := parseCoordinates(&(*coordinatesSlice)[i])
		if err != nil {
			return nil, err
		}
		res[i] = coordinates
	}
	return res, nil
}

func parseBounds(min *string, max *string) (*planeproto.Bounds, error) {
	minCoordinates, err := parseCoordinates(min)
	if err != nil {
		return nil, err
	}
	maxCoordinates, err := parseCoordinates(max)
	if err != nil {
		return nil, err
	}

	if minCoordinates.X >= maxCoordinates.X {
		return nil, errors.New("min.x must be less than max.x")
	}
	if minCoordinates.Y >= maxCoordinates.Y {
		return nil, errors.New("min.y must be less than max.y")
	}

	bounds := &planeproto.Bounds{
		Min: minCoordinates,
		Max: maxCoordinates,
	}
	return bounds, nil
}

func printPoint(point *planeproto.Point) {
	fmt.Printf("Coordinates:  %d,%d\n", point.Coordinates.X, point.Coordinates.Y)
	fmt.Printf("Count:        %d\n", point.Count)
}

func printPoints(points []*planeproto.Point) {
	for _, point := range points {
		printPoint(point)
	}
}

func printBounds(bounds *planeproto.Bounds) {
	fmt.Printf("Min:  %d,%d\n", bounds.Min.X, bounds.Min.Y)
	fmt.Printf("Max:  %d,%d\n", bounds.Max.X, bounds.Max.Y)
}

func run() error {
	getpoint.Action(func(parseContext *kingpin.ParseContext) error {
		coordinates, err := parseCoordinates(getpointCoordinates)
		if err != nil {
			return err
		}

		point, err := plane.GetPoint(coordinates)
		if err != nil {
			return err
		}

		printPoint(point)
		return nil
	})

	putpoints.Action(func(parseContext *kingpin.ParseContext) error {
		coordinatesSlice, err := parseCoordinatesSlice(putpointsCoordinatesSlice)
		if err != nil {
			return err
		}

		bounds, err := plane.PutPoints(coordinatesSlice)
		if err != nil {
			return err
		}

		printBounds(bounds)
		return nil
	})

	listpointsbybounds.Action(func(parseContext *kingpin.ParseContext) error {
		bounds, err := parseBounds(listpointsbyboundsBoundsMin, listpointsbyboundsBoundsMax)
		if err != nil {
			return err
		}

		points, err := plane.ListPointsByBounds(bounds)
		if err != nil {
			return err
		}

		printPoints(points)
		return nil
	})

	listpoints.Action(func(parseContext *kingpin.ParseContext) error {
		coordinatesSlices, err := parseCoordinatesSlice(listpointsCoordinatesSlice)
		if err != nil {
			return err
		}

		points, err := plane.ListPoints(coordinatesSlices)
		if err != nil {
			return err
		}

		printPoints(points)
		return nil
	})

	_, err := app.Parse(os.Args[1:])
	return err
}

func main() {
	if err := run(); err != nil {
		log.Fatalln(err)
	}
}
