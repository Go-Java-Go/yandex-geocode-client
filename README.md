# yandex-geocode-client

First, ensure the library is installed and up to date by running 
`go get -u github.com/Go-Java-Go/yandex-geocode-client`.

Second, you mast be get api key [JavaScript API и HTTP Геокодер](https://developer.tech.yandex.ru/services/)

Third, you mast get acquainted [documentation](https://yandex.ru/dev/maps/geocoder/doc/desc/concepts/input_params.html)  

## Example

```go
package main

import (
	"github.com/rs/zerolog/log"
	yg_client "github.com/Go-Java-Go/yandex-geocode-client"
)

func main() {

	client, err := yg_client.NewClient(yg_client.Config{
		APIKey:     "887d760f-df46-420a-949d-74da00e47fe5",
		Host:       "https://geocode-maps.yandex.ru/",
		ApiVersion: "1"})	
	if err != nil {
		log.Error().Err(err).Msg("")
	}

	resp, err := client.Geocode().Search(yg_client.SearchBaseRequestParams{Geocode: "Москва", Lang: ruRU, ResponseFormat: JSON, Sco: latLong})
	if err != nil {
		log.Error().Err(err).Msg("")
		return
	}

}
```

