package dht22

import (
	"errors"
	"time"

	"github.com/kidoman/embd"
	_ "github.com/kidoman/embd/host/rpi"
)

// AM2302/DHT22 - digital relative humidity and temperature sensor
// https://cdn-shop.adafruit.com/datasheets/Digital+humidity+and+temperature+sensor+AM2302.pdf

const (
	COLLECTING_PERIOD  = 2 * time.Second
	LOGICAL_1_TRESHOLD = 50 * time.Microsecond
)

var (
	ChecksumError    = errors.New("checksum error")
	HumidityError    = errors.New("humidity range error")
	TemperatureError = errors.New("temperature range error")
)

type DHT22 struct {
	pin         string
	temperature float32
	humidity    float32
	readAt      time.Time
	err         error
}

func New(pin string) *DHT22 {
	return &DHT22{pin: pin}
}

func (d *DHT22) Temperature() (float32, error) {
	if err := d.read(); err != nil {
		d.err = err
		return 0, err
	}

	return d.temperature, nil
}

func (d *DHT22) Humidity() (float32, error) {
	if err := d.read(); err != nil {
		d.err = err
		return 0, err
	}

	return d.humidity, nil
}

func (d *DHT22) read() error {
	if d.readAt.Add(COLLECTING_PERIOD).After(time.Now()) {
		return d.err
	}

	d.err = nil

	d.readAt = time.Now()

	// early allocations before time critical code
	lengths := make([]time.Duration, 40)
	iterator := 0

	embd.InitGPIO()
	defer embd.CloseGPIO()

	pin, err := embd.NewDigitalPin(d.pin)
	if err != nil {
		return err
	}

	defer pin.Close()

	if err := pin.SetDirection(embd.Out); err != nil {
		return err
	}

	if err := pin.Write(embd.High); err != nil {
		return err
	}

	// send init values
	time.Sleep(250 * time.Millisecond)
	if err := pin.Write(embd.Low); err != nil {
		return err
	}

	time.Sleep(5 * time.Millisecond)

	if err := pin.Write(embd.High); err != nil {
		return err
	}

	time.Sleep(20 * time.Microsecond)

	if err := pin.SetDirection(embd.In); err != nil {
		return err
	}

	// read data
	for {
		duration, err := pin.TimePulse(embd.High)
		if err != nil {
			return err
		}
		lengths[iterator] = duration
		iterator++
		if iterator >= 40 {
			break
		}
	}

	// convert to bytes
	bytes := make([]uint8, 5)

	for i := range bytes {
		for j := 0; j < 8; j++ {
			bytes[i] <<= 1
			if lengths[i*8+j] > LOGICAL_1_TRESHOLD {
				bytes[i] |= 0x01
			}
		}
	}

	if err := d.checksum(bytes); err != nil {
		if err != nil {
			return err
		}
	}

	var (
		humidity    uint16
		temperature uint16
	)

	// calculate humidity

	humidity |= uint16(bytes[0])
	humidity <<= 8
	humidity |= uint16(bytes[1])

	if humidity < 0 || humidity > 1000 {
		return HumidityError
	}

	d.humidity = float32(humidity) / 10

	// calculate temperature
	temperature |= uint16(bytes[2])
	temperature <<= 8
	temperature |= uint16(bytes[3])

	// check for negative temperature
	if temperature&0x8000 > 0 {
		d.temperature = float32(temperature&0x7FFF) / -10
	} else {
		d.temperature = float32(temperature) / 10
	}

	// datasheet operating range
	if d.temperature < -40 || d.temperature > 80 {
		return TemperatureError
	}

	return nil
}

func (d *DHT22) checksum(bytes []uint8) error {
	var sum uint8

	for i := 0; i < 4; i++ {
		sum += bytes[i]
	}

	if sum != bytes[4] {
		return ChecksumError
	}

	return nil
}
