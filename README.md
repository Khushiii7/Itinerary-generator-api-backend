# Itinerary Generator API (Go)

A backend service that generates PDF itineraries from JSON input. This API processes travel data and creates well-formatted PDF documents with travel details, schedules and booking information.


## Prerequisites
- Docker (with Docker Compose) - [Download Docker](https://www.docker.com/products/docker-desktop)
- 4GB+ of free RAM (required for Chromium PDF generation)

## Getting Started

### Method 1: Using Docker (Recommended)

1. **Clone the repository**
   ```bash
   git clone <repository-url>
   cd Itenary_Backend_API
   ```

2. **Build and run the container**
   ```bash
   docker-compose up -d
   ```

### Method 2: Manual Setup

1. **Install Go 1.24.5 or later**
   - Download from [golang.org](https://golang.org/dl/)

2. **Install dependencies**
   ```bash
   go mod download
   ```

3. **Run the application**
   ```bash
   go run main.go
   ```

## API Usage

### Generate Itinerary

**Endpoint**: `POST /generate-itinerary`

**Headers**:
- `Content-Type: application/json`

**Example Request**:
```bash
curl -X POST http://localhost:8080/generate-itinerary \
  -H "Content-Type: application/json" \
  -d '@input/khushi.json'
```

**Example Response**:
```json
{
  "status": "success",
  "message": "Itinerary generated successfully",
  "filename": "itinerary_Khushi_20250725.pdf"
}
```

### Sample Input

Create a JSON file (e.g., `input/my_trip.json`) with your itinerary details:

```json
{
  "trip_name": "Mumbai Getaway",
  "user_name": "Khushi",
  "start_date": "2025-07-25",
  "end_date": "2025-07-28",
  "Payment": {
    "Method": "Credit Card",
    "Installment": "Full Payment",
    "Date": "2025-07-22"
  },
  "days": [
    {
      "day_number": 1,
      "date": "2025-07-25",
      "time_groups": [
        {
          "time_of_day": "Morning",
          "activities": [
            {
              "name": "Flight to Mumbai",
              "time": "07:00 AM",
              "location": "DEL to BOM"
            }
          ]
        },
        {
          "time_of_day": "Afternoon",
          "activities": [
            {
              "name": "Lunch at Restaurant",
              "time": "01:30 PM",
              "location": "Mumbai Downtown"
            }
          ]
        },
        {
          "time_of_day": "Evening",
          "activities": [
            {
              "name": "Evening Walk",
              "time": "07:00 PM",
              "location": "Marine Drive"
            }
          ]
        }
      ]
    }
  ]
}
```

## Output
Generated PDFs are saved in the `output` directory with filenames in the format: `itinerary_[username]_[date].pdf`

## Development

```
### Building the Docker Image

```bash
docker build -t itinerary-api .
```


## License
This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
