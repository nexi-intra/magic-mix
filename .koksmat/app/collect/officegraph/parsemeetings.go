package officegraph

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"
)

type Booking struct {
	ID    string `json:"id"`
	Start struct {
		DateTime string `json:"dateTime"`
		TimeZone string `json:"timeZone"`
	} `json:"start"`
	End struct {
		DateTime string `json:"dateTime"`
		TimeZone string `json:"timeZone"`
	} `json:"end"`
}

type Room struct {
	ParentID struct {
		EmailAddress string `json:"emailAddress"`
	} `json:"parentId"`
	Details []Booking `json:"details"`
}

type RoomData struct {
	Data Room   `json:"data"`
	File string `json:"file"`
}

type Timeslot struct {
	Hour   int `json:"hour"`
	Minute int `json:"minute"`
}

type TimeslotEntry struct {
	Date  string     `json:"date"`
	Slots []Timeslot `json:"slots"`
}

type BookingOutput struct {
	ResourceId string          `json:"resource_id"`
	BookingID  string          `json:"booking_id"`
	StartDate  string          `json:"start_date"`
	EndDate    string          `json:"end_date"`
	Timeslots  []TimeslotEntry `json:"timeslots"`
}

func ParseMeetings(inputfile string) {
	// Read JSON file
	jsonFile, err := os.Open(inputfile)
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer jsonFile.Close()

	// Parse JSON file
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var roomData []RoomData
	if err := json.Unmarshal(byteValue, &roomData); err != nil {
		log.Fatalf("Failed to parse JSON: %v", err)
	}

	var results []BookingOutput

	// Process each booking
	for _, room := range roomData {
		for _, booking := range room.Data.Details {
			start, err := time.Parse(time.RFC3339, booking.Start.DateTime[:19]+"Z")
			if err != nil {
				log.Fatalf("Failed to parse start time: %v", err)
			}
			end, err := time.Parse(time.RFC3339, booking.End.DateTime[:19]+"Z")
			if err != nil {
				log.Fatalf("Failed to parse end time: %v", err)
			}

			bookingOutput := BookingOutput{
				ResourceId: room.Data.ParentID.EmailAddress,
				BookingID:  booking.ID,
				StartDate:  start.Format("2006-01-02"),
				EndDate:    end.Format("2006-01-02"),
				Timeslots:  generateTimeslots(start, end),
			}
			results = append(results, bookingOutput)
		}
	}

	// Write the result to the output JSON file
	outputFile := "booking_timeslots_output.json"
	outputData, err := json.MarshalIndent(results, "", "  ")
	if err != nil {
		log.Fatalf("Failed to marshal results: %v", err)
	}
	if err := ioutil.WriteFile(outputFile, outputData, 0644); err != nil {
		log.Fatalf("Failed to write output file: %v", err)
	}

	fmt.Printf("Booking timeslots data has been written to %s\n", outputFile)
}

func generateTimeslots(start, end time.Time) []TimeslotEntry {
	interval := 15 * time.Minute
	var timeslots []TimeslotEntry

	for current := start; current.Before(end); {
		date := current.Format("2006-01-02")
		slotEntry := TimeslotEntry{
			Date:  date,
			Slots: []Timeslot{},
		}

		for current.Before(end) && current.Format("2006-01-02") == date {
			slotEntry.Slots = append(slotEntry.Slots, Timeslot{
				Hour:   current.Hour(),
				Minute: current.Minute(),
			})
			current = current.Add(interval)
		}

		timeslots = append(timeslots, slotEntry)
	}

	return timeslots
}
