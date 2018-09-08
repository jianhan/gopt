package place

import (
	"errors"
	"fmt"
	"github.com/asaskevich/govalidator"
	ghttp "github.com/jianhan/gopt/http"
	"googlemaps.github.io/maps"
	"net/http"
	"strings"
)

const MaxRadius = 50000

type NearbySearchRequestOption func(*maps.NearbySearchRequest) error

type NearbySearchRequestOptions struct {
}

func (n NearbySearchRequestOptions) parsePriceLevel(fieldName, priceLevel string) (maps.PriceLevel, error) {
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
		return "", &ghttp.HttpError{
			Message: fmt.Sprintf("invalid %s", fieldName),
			Status:  http.StatusBadRequest,
		}
	}
}

func (n NearbySearchRequestOptions) Raidus(radius uint) NearbySearchRequestOption {
	return func(args *maps.NearbySearchRequest) error {
		if !govalidator.InRange(radius, 1, MaxRadius) {
			return &ghttp.HttpError{
				Message: fmt.Sprintf("invalid radius parameter, must be greater than 1 and less than %d", MaxRadius),
				Status:  http.StatusBadRequest,
			}
		}
		args.Radius = radius

		return nil
	}
}

func (n NearbySearchRequestOptions) Location(location string) NearbySearchRequestOption {
	return func(args *maps.NearbySearchRequest) error {
		latLng, err := maps.ParseLatLng(location)
		if err != nil {
			return err
		}
		args.Location = &latLng

		return nil
	}
}

func (n NearbySearchRequestOptions) Keyword(keyword string) NearbySearchRequestOption {
	return func(args *maps.NearbySearchRequest) error {
		if strings.Trim(keyword, " ") == "" {
			return &ghttp.HttpError{
				Message: "keyword can not be empty",
				Status:  http.StatusBadRequest,
			}
		}
		args.Keyword = keyword

		return nil
	}
}

func (n NearbySearchRequestOptions) Language(language string) NearbySearchRequestOption {
	return func(args *maps.NearbySearchRequest) error {
		if strings.Trim(language, " ") == "" {
			return &ghttp.HttpError{
				Message: "language can not be empty",
				Status:  http.StatusBadRequest,
			}
		}
		args.Language = language

		return nil
	}
}

func (n NearbySearchRequestOptions) MinPrice(minPrice string) NearbySearchRequestOption {
	return func(args *maps.NearbySearchRequest) (err error) {
		if args.MinPrice, err = n.parsePriceLevel("min price", minPrice); err != nil {
			return err
		}

		return nil
	}
}

func (n NearbySearchRequestOptions) MaxPrice(maxPrice string) NearbySearchRequestOption {
	return func(args *maps.NearbySearchRequest) (err error) {
		if args.MaxPrice, err = n.parsePriceLevel("max price", maxPrice); err != nil {
			return err
		}

		return nil
	}
}

func (n NearbySearchRequestOptions) Name(name string) NearbySearchRequestOption {
	return func(args *maps.NearbySearchRequest) error {
		if strings.Trim(name, " ") == "" {
			return &ghttp.HttpError{
				Message: "name can not be empty",
				Status:  http.StatusBadRequest,
			}
		}

		args.Name = name

		return nil
	}
}

func (n NearbySearchRequestOptions) OpenNow(openNow bool) NearbySearchRequestOption {
	return func(args *maps.NearbySearchRequest) error {
		args.OpenNow = openNow
		return nil
	}
}

func (n NearbySearchRequestOptions) RankBy(rankBy string) NearbySearchRequestOption {
	return func(args *maps.NearbySearchRequest) error {
		switch rankBy {
		case "prominence":
			args.RankBy = maps.RankByProminence
		case "distance":
			args.RankBy = maps.RankByDistance
		default:
			return errors.New(fmt.Sprintf("Unknown rank by: \"%v\"", rankBy))
		}

		return nil
	}
}

func (n NearbySearchRequestOptions) Type(placeType string) NearbySearchRequestOption {
	return func(args *maps.NearbySearchRequest) (err error) {
		if args.Type, err = maps.ParsePlaceType(placeType); err != nil {
			return err
		}

		return nil
	}
}

func (n NearbySearchRequestOptions) PageToken(pageToken string) NearbySearchRequestOption {
	return func(args *maps.NearbySearchRequest) error {
		if strings.Trim(pageToken, " ") == "" {
			return &ghttp.HttpError{
				Message: "page token can not be empty",
				Status:  http.StatusBadRequest,
			}
		}
		args.PageToken = pageToken

		return nil
	}
}
