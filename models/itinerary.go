package models

type Activity struct {
    Name     string `json:"name"`
    Time     string `json:"time"`
    Location string `json:"location"`
}

type TimeGroup struct {
    TimeOfDay string     `json:"time_of_day"`
    Activities []Activity `json:"activities"`
}

type Day struct {
    DayNumber  int         `json:"day_number"`
    Date       string      `json:"date"`
    TimeGroups []TimeGroup `json:"time_groups"`
}

type ItineraryRequest struct {
    TripName     string `json:"trip_name"`
    UserName     string `json:"user_name"`
    StartDate    string `json:"start_date"`
    EndDate      string `json:"end_date"`
    Days         []Day  `json:"days"`
}