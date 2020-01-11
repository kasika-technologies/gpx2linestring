package cmd

import (
	"encoding/json"
	"errors"
	"github.com/kasika-technologies/gpx2linestring/entities"
	"github.com/spf13/cobra"
	"github.com/twpayne/go-gpx"
	"os"
	"path/filepath"
	"strings"
)

var rootCmd = &cobra.Command{
	Use:                "gpx2linestring",
	Short:              "GPX to GeoJSON LineString",
	RunE:               convert,
	FParseErrWhitelist: cobra.FParseErrWhitelist{},
}

func Execute() error {
	return rootCmd.Execute()
}

func convert(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		err := errors.New("Please input GPX filename.")
		return err
	}

	inputFilepath := args[0]

	inputFile, err := os.Open(inputFilepath)
	if err != nil {
		return err
	}
	defer inputFile.Close()

	t, err := gpx.Read(inputFile)
	if err != nil {
		return err
	}

	var coordinates [][]float64

	for _, wpt := range t.Trk[0].TrkSeg[0].TrkPt {
		coordinate := []float64{wpt.Lon, wpt.Lat}
		coordinates = append(coordinates, coordinate)
	}

	geometry := &entities.Geometry{
		Type:        "LineString",
		Coordinates: coordinates,
	}

	j, err := json.Marshal(geometry)
	if err != nil {
		return err
	}

	outputFilepath := strings.TrimSuffix(inputFilepath, filepath.Ext(inputFilepath))
	outputFilepath += ".json"

	outputFile, err := os.Create(outputFilepath)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	outputFile.Write(j)

	return nil
}
