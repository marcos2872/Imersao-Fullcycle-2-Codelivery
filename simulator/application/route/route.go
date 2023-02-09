package route

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
	"strconv"
	"strings"
)

type Route struct {
		ID string `json:"routeId"`
		ClientID string `json:"clientId"`
		Positions []Position `json:"position"`
}

type Position struct {
		Lat float64 `json:"lat"`
		Long float64 `json:"long"`
}

type PartialRoutePosition struct {
	ID string `json:"routeId"`
	ClientID string `json:"clientId"`
	Position []float64 `json:"position"`
	Finished bool `json:"finished"`
}

func(r *Route) LoadPositions() error {
	if r.ID == "" {
		return errors.New("route id not informed")
	}
	f, err := os.Open("destinations/" + r.ID + ".txt")
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		data := strings.Split(scanner.Text(), ",")
		lat, err := strconv.ParseFloat(data[0], 64)
		if err != nil {
			return nil
		}
		long, err := strconv.ParseFloat(data[1], 64)
		if err != nil {
			return nil
		}
		r.Positions = append(r.Positions, Position{
			Lat: lat,
			Long: long,
		})
	}
	return nil
}

func (r *Route) ExportJsonPositions() ([]string, error) {
	var route PartialRoutePosition
	var result []string
	total := len(r.Positions)
	for i, c := range r.Positions {
		route.ID = r.ID
		route.ClientID = r.ClientID
		route.Position = []float64{c.Lat, c.Long}
		route.Finished = false
		if total -1 == i {
			route.Finished = true
		}
		jsonRute, err := json.Marshal(route)
		if err != nil {
			return nil, err
		}
		result = append(result, string(jsonRute))
	}
	return result, nil
}