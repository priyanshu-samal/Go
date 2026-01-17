package main

import "fmt"

type FeatureFlag struct {
	enabled bool
}

func (f *FeatureFlag) enable() {
	f.enabled = true
}

func (f *FeatureFlag) disable() {
	f.enabled = false
}

func (f FeatureFlag) isEnabled() bool {
	return f.enabled
}

func main() {
	v := FeatureFlag{false}
	v.enable()
	fmt.Println(v.isEnabled())
	v.disable()
	fmt.Println(v.isEnabled())
}