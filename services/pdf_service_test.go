package services

import (
	"Itenary_Backend_API/models"
	"os"
	"path/filepath"
	"testing"
	"time"
)
//tests the GeneratePDF function
func TestGeneratePDF(t *testing.T) {
	testCases := []struct {
		name        string
		setup       func() models.ItineraryRequest
		wantErr     bool
		cleanup     func()
	}{
		{
			name: "successful PDF generation",
			setup: func() models.ItineraryRequest {
				return models.ItineraryRequest{
					TripName:  "Test Trip",
					UserName:  "Test User",
					StartDate: "2025-07-22",
					EndDate:   "2025-07-25",
					Days: []models.Day{
						{
							DayNumber: 1,
							Date:      "2025-07-22",
							Activities: []models.Activity{
								{
									Name:     "Breakfast",
									Time:     "08:00 AM",
									Location: "Hotel Restaurant",
								},
							},
						},
					},
				}
			},
			wantErr: false,
			cleanup: func() {
				// Clean
				files, _ := filepath.Glob("output/itinerary_Test User.*")
				for _, f := range files {
					os.Remove(f)
				}
			},
		},
		{
			name: "empty user name",
			setup: func() models.ItineraryRequest {
				req := models.ItineraryRequest{
					TripName: "Test Trip",
					StartDate: "2025-07-22",
					EndDate:   "2025-07-25",
				}
				return req
			},
			wantErr: false, 
			cleanup: func() {
				files, _ := filepath.Glob("output/itinerary_itinerary.*")
				for _, f := range files {
					os.Remove(f)
				}
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			req := tc.setup()
			os.MkdirAll("output", 0755)

			pdfPath, err := GeneratePDF(req)

			if (err != nil) != tc.wantErr {
				t.Fatalf("GeneratePDF() error = %v, wantErr %v", err, tc.wantErr)
			}

			if !tc.wantErr {
				if _, err := os.Stat(pdfPath); os.IsNotExist(err) {
					t.Fatalf("Expected PDF file was not created: %s", pdfPath)
				}

				fileInfo, err := os.Stat(pdfPath)
				if err != nil {
					t.Fatalf("Error checking PDF file: %v", err)
				}
				if fileInfo.Size() == 0 {
					t.Error("Generated PDF file is empty")
				}
			}
			if tc.cleanup != nil {
				tc.cleanup()
			}
		})
	}
}

func TestGeneratePDF_ErrorCases(t *testing.T) {
	tests := []struct {
		name    string
		setup   func()
		cleanup func()
		req     models.ItineraryRequest
		wantErr bool
	}{
		{
			name: "invalid template path",
			setup: func() {
				os.Chdir("..")
			},
			cleanup: func() {
				os.Chdir("services")
			},
			req: models.ItineraryRequest{
				TripName:  "Test",
				UserName:  "User",
				StartDate: "2025-01-01",
				EndDate:   "2025-01-02",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setup != nil {
				tt.setup()
			}
			_, err := GeneratePDF(tt.req)

			if (err != nil) != tt.wantErr {
				t.Errorf("GeneratePDF() error = %v, wantErr %v", err, tt.wantErr)
			}

			if tt.cleanup != nil {
				tt.cleanup()
			}
		})
	}
}

//verifies that the current date is correctly formatted
func TestHelper_CurrentDate(t *testing.T) {
	now := time.Now()
	req := models.ItineraryRequest{
		TripName:  "Test",
		UserName:  "TestUser",
		StartDate: "2025-01-01",
		EndDate:   "2025-01-02",
	}

	pdfPath, err := GeneratePDF(req)
	if err != nil {
		t.Fatalf("GeneratePDF() error = %v", err)
	}
	defer os.Remove(pdfPath)

	//verify the HTML file was created and contains the current date
	htmlPath := "output/itinerary_TestUser.html"
	content, err := os.ReadFile(htmlPath)
	if err != nil {
		t.Fatalf("Error reading generated HTML: %v", err)
	}

	//format January 2, 2006
	expectedDate := now.Format("January 2, 2006")
	if !strings.Contains(string(content), expectedDate) {
		t.Errorf("Expected to find date '%s' in the generated HTML, but didn't", expectedDate)
	}
}
