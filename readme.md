Greetings do it yourselfers!

It's AM2302/DHT22 library. No external c libraries. 

## installation 
`go get github.com/morus12/dht22`

## usage
```
import "github.com/morus12/dht22"

sensor := dht22.New("GPIO_17")
temperature, err := sensor.Temperature()
humidity, err := sensor.Humidity()
```

## testing environmnet 
- Raspberry Pi 3
- two DHT22 sensors
- go1.7.5 linux/arm
- Raspbian GNU/Linux 8

Pi3 has 64 bit CPU, but system was setup on 32bit. Currently (Feb 2016) there is no official linux/arm64 kernel for Pi 3.