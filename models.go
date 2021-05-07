package yandex_geocode_client

type internalRequest struct {
	endpoint string
	method   string

	withRequest     interface{}
	withResponse    interface{}
	withQueryParams map[string]string

	acceptedStatusCodes []int

	functionName string
	apiName      string
}

type ResponseFormat string

const (
	XML  ResponseFormat = "xml"
	JSON ResponseFormat = "json"
)

func (r ResponseFormat) String() string {
	return string(r)
}

type Lang string

const (
	trTR Lang = "tr_TR" /* trTR Turkish (only for maps of Turkey).*/
	enRU Lang = "en_RU" /* enRU response in English, Russian map features;*/
	enUS Lang = "en_RU" /* enUS response in English, Russian map features;*/
	ruRU Lang = "ru_RU" /* ruRU Russian (default);*/
	ukUA Lang = "uk_UA" /* ukUA Ukrainian;*/
	beBY Lang = "be_BY" /* beBY Belarusian.*/
)

func (l Lang) String() string {
	return string(l)
}

type Kind string

const (
	house    Kind = "house"
	street   Kind = "street"
	metro    Kind = "metro"    /* metro subway station*/
	district Kind = "district" /* district geocode district*/
	locality Kind = "locality" /* locality (geocode, town, village, etc.)*/
)

func (k Kind) String() string {
	return string(k)
}

type Sco string

const (
	longLat Sco = "longlat" /* Longitude, latitude (default) */
	latLong Sco = "latlong" /* Latitude, longitude */
)

func (s Sco) String() string {
	return string(s)
}

type SearchBaseRequestParams struct {
	//Geocode address or geographical coordinates of the object being searched for. The specified data determines the type of geocoding:
	//If an address is specified, it is converted to object coordinates. This process is called forward geocoding.
	//If coordinates are specified, they are converted to the object's address. This process is called reverse geocoding.
	Geocode string

	//Sco only if the geocode parameter sets the coordinates. Order of coordinates. Possible values:
	Sco Sco

	//Kind Only if the geocode parameter sets the coordinates. The type of required toponym.
	Kind Kind

	//The Spn of the search area. The center of the area is set in the ll parameter.
	Spn []string

	//Bbox An alternative method for setting the search area. Record format: bbox=x1,y1~x2,y2
	Bbox           string
	Lang           Lang
	ResponseFormat ResponseFormat
}

type SearchPaginationBaseRequestParams struct {
	SearchBaseRequestParams
	Limit  int
	Offset int
}

type SearchObjectRequest struct {
	SearchPaginationBaseRequestParams
	Types  []ResponseFormat
	SortBy string
	Query  string
	Region string
	Lat    float64
	Lon    float64
	Radius int
}

func (s SearchBaseRequestParams) toQueryParam() map[string]string {
	q := make(map[string]string, 4)
	languages := s.Lang
	if languages != "" {
		q["lang"] = s.Lang.String()
	}

	geocode := s.Geocode
	if geocode != "" {
		q["geocode"] = geocode
	}

	sco := s.Sco
	if sco != "" {
		q["sco"] = s.Sco.String()
	}

	if s.Kind != "" {
		q["kind"] = s.Kind.String()
	}

	if s.Bbox != "" {
		q["bbox"] = s.Bbox
	}

	if s.ResponseFormat != "" {
		q["format"] = s.ResponseFormat.String()
	}

	return q
}

type QueryParam interface {
	toQueryParam() map[string]string
}

type Tit struct {
	Response struct {
		GeoObjectCollection GeoObjectCollection `json:"GeoObjectCollection"`
	} `json:"response"`
}

type GeoObjectCollection struct {
	MetaDataProperty MetaDataProperty `json:"metaDataProperty"`
	FeatureMember    []struct {
		GeoObject GeoObject `json:"GeoObject"`
	} `json:"featureMember"`
}
type GeoObject struct {
	MetaDataProperty struct {
		GeocoderMetaData struct {
			Precision      string         `json:"precision"`
			Text           string         `json:"text"`
			Kind           string         `json:"kind"`
			Address        Address        `json:"Address"`
			AddressDetails AddressDetails `json:"AddressDetails"`
		} `json:"GeocoderMetaData"`
	} `json:"metaDataProperty"`
	Name        string `json:"name"`
	Description string `json:"description"`
	BoundedBy   struct {
		Envelope struct {
			LowerCorner string `json:"lowerCorner"`
			UpperCorner string `json:"upperCorner"`
		} `json:"Envelope"`
	} `json:"boundedBy"`
	Point struct {
		Pos string `json:"pos"`
	} `json:"Point"`
}

type MetaDataProperty struct {
	GeocoderResponseMetaData struct {
		Request string `json:"request"`
		Results string `json:"results"`
		Found   string `json:"found"`
	} `json:"GeocoderResponseMetaData"`
}

type Address struct {
	CountryCode string `json:"country_code"`
	Formatted   string `json:"formatted"`
	Components  []struct {
		Kind string `json:"kind"`
		Name string `json:"name"`
	} `json:"Components"`
}

type AddressDetails struct {
	Country struct {
		AddressLine        string `json:"AddressLine"`
		CountryNameCode    string `json:"CountryNameCode"`
		CountryName        string `json:"CountryName"`
		AdministrativeArea struct {
			AdministrativeAreaName string `json:"AdministrativeAreaName"`
			SubAdministrativeArea  struct {
				SubAdministrativeAreaName string   `json:"SubAdministrativeAreaName"`
				Locality                  Locality `json:"Locality"`
			} `json:"SubAdministrativeArea"`
		} `json:"AdministrativeArea"`
	} `json:"Country"`
}

type Locality struct {
	DependentLocality struct {
		DependentLocalityName string `json:"DependentLocalityName"`
		Thoroughfare          struct {
			ThoroughfareName string `json:"ThoroughfareName"`
		} `json:"Thoroughfare"`
	} `json:"DependentLocality"`
}
