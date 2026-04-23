package movie

type Kind string

type kinds struct {
	Unknown    Kind
	Animated   Kind
	Theatrical Kind
	Mixed      Kind
}

var Kinds = kinds{
	Unknown:    "entertainment.movie.unknown",
	Animated:   "entertainment.movie.animated",
	Theatrical: "entertainment.movie.theatrical",
	Mixed:      "entertainment.movie.mixed",
}
