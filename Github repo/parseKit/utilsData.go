package parsekit

type Room struct {
	Name      string
	X, Y      int
	IsStart   bool
	IsEnd     bool
	Connected bool
}

type Tunnel struct {
	From string
	To   string
}

var (
	AntNum    int
	Err       error
	StartRoom string
	EndRoom   string
	Rooms     = make(map[string]*Room)
	Tunnels   []Tunnel
)
