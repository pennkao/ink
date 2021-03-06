package app

func DefaultConfig() Config {
	return Config{
		Snapshot: SnapshotConfig{
			Width:  2400,
			Height: 2400,
			Dir:    "snapshots",
		},
		Window: WindowConfig{
			Title:  "ink",
			Width:  800,
			Height: 800,
		},
	}
}

type Config struct {
	Snapshot SnapshotConfig
	Window   WindowConfig
	Trace    bool
}

type WindowConfig struct {
	Title         string
	Width, Height int
	X, Y          int
	Hidden        bool
}

type SnapshotConfig struct {
	Width, Height int
	Dir           string
}
