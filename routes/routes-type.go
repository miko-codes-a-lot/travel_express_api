package routes

import "cloudants/travel-express/helper"

type route struct {
	ID            int     `json:"id"`
	RouteCode     string  `json:"routeCode"`
	Origin        string  `json:"origin"`
	Destination   string  `json:"destination"`
	DurationHours float32 `json:"durationHours"`
	DistanceKM    float32 `json:"distanceKM"`
	IsActive      bool    `json:"isActive"`
	CreatedAt     string  `json:"createdAt"`
	UpdatedAt     string  `json:"updatedAt"`
}

// 'MNL-CEB', 'Manila North Harbor', 'Cebu City Port', 20, 570
var routes = []route{
	{

		ID:            1,
		RouteCode:     "MNL-CEB",
		Origin:        "Manila North Harbor",
		Destination:   "Cebu City Port",
		DurationHours: 20,
		DistanceKM:    570,
		IsActive:      true,
		CreatedAt:     helper.Now(),
		UpdatedAt:     helper.Now(),
	},
}
