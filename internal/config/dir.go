package config

type Dir struct {
	Views string
}

func getDir() *Dir {
	return &Dir{
		Views: "/",
	}
}
