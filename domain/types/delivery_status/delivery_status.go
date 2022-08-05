package delivery_status

type Status int

const (
	Created        Status = 1
	LoadedIntoSack Status = 2
	Loaded         Status = 3
	Unloaded       Status = 4
)
