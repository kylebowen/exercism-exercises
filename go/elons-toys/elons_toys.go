package elon

import "fmt"

// Drive updates the meters driven based on the car's speed and reduces the battery based on the battery drainage.
func (c *Car) Drive() {
	if c.battery-c.batteryDrain > 0 {
		c.battery -= c.batteryDrain
		c.distance += c.speed
	}
}

// DisplayDistance returns the distance traveled for display.
func (c Car) DisplayDistance() string {
	return fmt.Sprintf("Driven %v meters", c.distance)
}

// DisplayBattery returns the battery percentage for display.
func (c Car) DisplayBattery() string {
	return fmt.Sprintf("Battery at %v%%", c.battery)
}

// CanFinish calculates if a car has enough battery left to finish a track.
func (c Car) CanFinish(trackDistance int) bool {
	return (c.battery/c.batteryDrain)*c.speed >= trackDistance
}
