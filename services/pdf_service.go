package services
import (
    "context"
    "fmt"
    "html/template"
    "io/ioutil"
    "os"
    "path/filepath"
    "strings"
    "time"
    "strconv"

    "Itenary_Backend_API/models"
    "github.com/chromedp/chromedp"
    "github.com/chromedp/cdproto/page"
)

//calculates the duration in days and nights
func calculateDuration(startDateStr, endDateStr string) (string, error) {
    layout := "2006-01-02" //format (YYYY-MM-DD)
    
    startDate, err := time.Parse(layout, startDateStr)
    if err != nil {
        return "", fmt.Errorf("invalid start date format: %v", err)
    }
    
    endDate, err := time.Parse(layout, endDateStr)
    if err != nil {
        return "", fmt.Errorf("invalid end date format: %v", err)
    }
    
    //difference in days
    duration := endDate.Sub(startDate)
    days := int(duration.Hours() / 24)
    
    //calculate nights
    nights := days - 1
    if nights < 0 {
        nights = 0
    }
    var durationStr string
    if days == 1 {
        durationStr = "1 Day"
        if nights == 1 {
            durationStr += " 1 Night"
        }
    } else if days > 1 {
        durationStr = fmt.Sprintf("%d Days", days)
        if nights > 0 {
            durationStr += fmt.Sprintf(" %d Nights", nights)
        }
    } else {
        durationStr = "1 Day" 
    }
    
    return durationStr, nil
}

func groupActivitiesByTimeOfDay(activities []models.Activity) []models.TimeGroup {
    timeGroups := []models.TimeGroup{
        {TimeOfDay: "Morning", Activities: []models.Activity{}},
        {TimeOfDay: "Afternoon", Activities: []models.Activity{}},
        {TimeOfDay: "Evening", Activities: []models.Activity{}},
    }

    for _, activity := range activities {
        timeStr := strings.TrimSpace(activity.Time)
        timeParts := strings.Fields(timeStr)
        if len(timeParts) < 2 {
            timeGroups[0].Activities = append(timeGroups[0].Activities, activity)
            continue
        }

        timeValue := timeParts[0]
        period := strings.ToUpper(timeParts[1])
        
        parts := strings.Split(timeValue, ":")
        if len(parts) < 2 {
            timeGroups[0].Activities = append(timeGroups[0].Activities, activity)
            continue
        }

        hour, err := strconv.Atoi(parts[0])
        if err != nil {
            timeGroups[0].Activities = append(timeGroups[0].Activities, activity)
            continue
        }

        if period == "PM" && hour < 12 {
            hour += 12
        } else if period == "AM" && hour == 12 {
            hour = 0
        }
        if hour<12 {
            timeGroups[0].Activities = append(timeGroups[0].Activities, activity) //Morning
        } else if hour < 17 {
            timeGroups[1].Activities = append(timeGroups[1].Activities, activity) //Afternoon
        } else {
            timeGroups[2].Activities = append(timeGroups[2].Activities, activity) //Evening
        }
    }

    result := make([]models.TimeGroup, 0, 3)
    for _, group := range timeGroups {
        if len(group.Activities) > 0 {
            result = append(result, group)
        }
    }

    return result
}

func GeneratePDF(req models.ItineraryRequest) (string, error) {
    duration, err := calculateDuration(req.StartDate, req.EndDate)
    if err != nil {
        return "", fmt.Errorf("error calculating duration: %v", err)
    }

    templateContent, err := os.ReadFile("templates/Itenary.html")
    if err != nil {
        return "", fmt.Errorf("error reading template file: %v", err)
    }

    tmpl, err := template.New("itinerary").Parse(string(templateContent))
    if err != nil {
        return "", fmt.Errorf("error parsing template: %v", err)
    }

    safeName := req.UserName
    if safeName == "" {
        safeName = "itinerary"
    }
    outputHTML := fmt.Sprintf("output/itinerary_%s.html", safeName)

    os.MkdirAll("output", os.ModePerm)

    file, err := os.Create(outputHTML)
    if err != nil {
        return "", fmt.Errorf("error creating output file: %v", err)
    }
    defer file.Close()

    data := map[string]interface{}{
        "TripName":    req.TripName,
        "UserName":    req.UserName,
        "StartDate":   req.StartDate,
        "EndDate":     req.EndDate,
        "Duration":    duration,
        "Days":        req.Days,
        "CurrentDate": time.Now().Format("January 2, 2006"),
    }

    if err := tmpl.Execute(file, data); err != nil {
        return "", fmt.Errorf("error executing template: %v", err)
    }

    ctx, cancel := chromedp.NewContext(context.Background())
    defer cancel()

    var pdfBuf []byte
    absPath, _ := filepath.Abs(outputHTML)
    fullPath := "file:///" + strings.ReplaceAll(absPath, "\\", "/")

    printToPDF := page.PrintToPDF()
    printToPDF = printToPDF.
        WithPrintBackground(true).
        WithMarginTop(0.4).
        WithMarginBottom(0.4).
        WithMarginLeft(0.4).
        WithMarginRight(0.4)

    err = chromedp.Run(ctx,
        chromedp.Navigate(fullPath),
        chromedp.Sleep(2*time.Second),
        chromedp.ActionFunc(func(ctx context.Context) error {
            var err error
            pdfBuf, _, err = printToPDF.Do(ctx)
            return err
        }),
    )
    if err != nil {
        return "", fmt.Errorf("error generating PDF: %v", err)
    }

    pdfPath := fmt.Sprintf("output/itinerary_%s.pdf", safeName)
    if err := ioutil.WriteFile(pdfPath, pdfBuf, 0644); err != nil {
        return "", fmt.Errorf("error writing PDF file: %v", err)
    }

    return pdfPath, nil
}
