package place

import "googlemaps.github.io/maps"

type NearbySearchRequestOption func(*maps.NearbySearchRequest)

type NearbySearchRequestOptions struct {
}

func (n NearbySearchRequestOptions) Raidus(radius uint) NearbySearchRequestOption {
	return func(args *maps.NearbySearchRequest) {
		args.Radius = radius
	}
}

func (n NearbySearchRequestOptions) Location(latLng *maps.LatLng) NearbySearchRequestOption {
	return func(args *maps.NearbySearchRequest) {
		args.Location = latLng
	}
}

func (n NearbySearchRequestOptions) Keyword(keyword string) NearbySearchRequestOption {
	return func(args *maps.NearbySearchRequest) {
		args.Keyword = keyword
	}
}

func (n NearbySearchRequestOptions) Language(language string) NearbySearchRequestOption {
	return func(args *maps.NearbySearchRequest) {
		args.Language = language
	}
}

func (n NearbySearchRequestOptions) MinPrice(minPrice maps.PriceLevel) NearbySearchRequestOption {
	return func(args *maps.NearbySearchRequest) {
		args.MinPrice = minPrice
	}
}

func (n NearbySearchRequestOptions) MaxPrice(maxPrice maps.PriceLevel) NearbySearchRequestOption {
	return func(args *maps.NearbySearchRequest) {
		args.MaxPrice = maxPrice
	}
}

func (n NearbySearchRequestOptions) Name(name string) NearbySearchRequestOption {
	return func(args *maps.NearbySearchRequest) {
		args.Name = name
	}
}

func (n NearbySearchRequestOptions) OpenNow(openNow bool) NearbySearchRequestOption {
	return func(args *maps.NearbySearchRequest) {
		args.OpenNow = openNow
	}
}

func (n NearbySearchRequestOptions) RankBy(rankBy maps.RankBy) NearbySearchRequestOption {
	return func(args *maps.NearbySearchRequest) {
		args.RankBy = rankBy
	}
}

func (n NearbySearchRequestOptions) Type(placeType maps.PlaceType) NearbySearchRequestOption {
	return func(args *maps.NearbySearchRequest) {
		args.Type = placeType
	}
}

func (n NearbySearchRequestOptions) PageToken(pageToken string) NearbySearchRequestOption {
	return func(args *maps.NearbySearchRequest) {
		args.PageToken = pageToken
	}
}
