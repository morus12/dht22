## Readme

It's AM2302/DHT22 library. No external c libraries. 

## Installation 
`go get github.com/pravipati/dht22`

## Usage
```
package main 

import (
  "github.com/morus12/dht22"
  "fmt"
  "log"
  "os"
)
func main () {
  sensor := dht22.New("17")
  temperature, err1 := sensor.Temperature()
  humidity, err2 := sensor.Humidity()
  if err1 != nil {
    log.Fatal(err1)
  } else if err2 != nil {
    log.Fatal(err2)
  }
  fmt.Printf("Temperature: %v, Humidity: %v\n", temperature, humidity)
}
```

## Testing environmnet 
- Raspberry Pi 2 Model B
- DHT22 sensor
- go version go1.9.2 linux/arm
