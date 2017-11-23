## Readme

A library for getting data from a [AM2302/DHT22](https://cdn-shop.adafruit.com/datasheets/Digital+humidity+and+temperature+sensor+AM2302.pdf) connected to a Raspberry Pi. Uses [embd](github.com/kidoman/embd). No external C libraries. 

## Installation : 

`go get github.com/pravipati/dht22`

## Usage:

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

## Testing Environment ✅:
- Raspberry Pi 2 Model B
- DHT22 sensor
- go version go1.9.2 linux/arm

# Measurement Details 📝:

**Humidity**:
  - Measured as [Relative Humidity](https://en.wikipedia.org/wiki/Relative_humidity)
  - Accurate within +-2%RH

**Temperature**:
  - Measured in Celsius
  - Accurate within +-0.5C

**Additional**:
  - One 100nF capacitor can be added between VDD and GND for wave filtering