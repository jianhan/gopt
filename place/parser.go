package place

import (
	"errors"
	"fmt"
	"googlemaps.github.io/maps"
)

func parsePriceLevel(priceLevel string) (maps.PriceLevel, error) {
	switch priceLevel {
	case "0":
		return maps.PriceLevelFree, nil
	case "1":
		return maps.PriceLevelInexpensive, nil
	case "2":
		return maps.PriceLevelModerate, nil
	case "3":
		return maps.PriceLevelExpensive, nil
	case "4":
		return maps.PriceLevelVeryExpensive, nil
	default:
		return "", errors.New("unknown price level")
	}
	return maps.PriceLevelFree, nil
}

func parsePriceLevels(minPrice string, maxPrice string, r *maps.NearbySearchRequest) error {
	if minPrice != "" {
		if r.MinPrice, err = parsePriceLevel(minPrice); err != nil {
			return err
		}
	}

	if maxPrice != "" {
		if r.MaxPrice, err = parsePriceLevel(minPrice); err != nil {
			return err
		}
	}

	return nil
}

func parseRankBy(rankBy string, r *maps.NearbySearchRequest) error {
	switch rankBy {
	case "prominence":
		r.RankBy = maps.RankByProminence
		return nil
	case "distance":
		r.RankBy = maps.RankByDistance
		return nil
	case "":
		return nil
	default:
		return errors.New(fmt.Sprintf("Unknown rank by: \"%v\"", rankBy))
	}
}

func parsePlaceType(placeType string, r *maps.NearbySearchRequest) error {
	if placeType != "" {
		t, err := maps.ParsePlaceType(placeType)
		if err != nil {
			return errors.New(fmt.Sprintf("Unknown place type \"%v\"", placeType))
		}

		r.Type = t
	}
}
