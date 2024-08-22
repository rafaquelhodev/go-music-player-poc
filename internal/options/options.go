package options

import "flag"

type Options struct {
	Bpm          *int
	Beats        *int
	Subdivisions *int
}

func ReadOptions() *Options {
	bpm := flag.Int("bpm", 60, "the BPM value")
	beats := flag.Int("beats", 4, "the number of beats")
	subdivions := flag.Int("sub", 1, "the number of subdivisions")

	defer flag.Parse()

	return &Options{
		Bpm:          bpm,
		Beats:        beats,
		Subdivisions: subdivions,
	}

}
