// the number of datatypes are
// 1.int
// 2.string
// 3.float
// 4.boolean
package main

type speed int
type name string
type freqency float64

func datatypes() (speed, name, freqency) {
	var carspeed speed = 100
	var myname name = "jobinaj"
	var bullet freqency = 104.92
	return carspeed, myname, bullet

}
