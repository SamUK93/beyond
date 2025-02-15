package beyond

type event struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Date string `json:"date"`
}

var events = []event{
	{ID: "1", Name: "Fusion Rocket Luanch", Date: "21st April 2025, 15:23"},
	{ID: "2", Name: "ISS Japan Docking", Date: "18th June 2025, 11:56"},
	{ID: "3", Name: "Mars Drone Imaging Report", Date: "1st March 2025, 14:11"},
}

func GetEvents() []event {
	return events
}
