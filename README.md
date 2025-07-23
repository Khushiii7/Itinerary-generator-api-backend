# Itinerary Generator API

A robust backend service that generates beautiful, printable PDF itineraries from JSON input. This API processes travel data and creates well-formatted PDF documents with travel details, schedules, and booking information.

## Features

- Generate professional PDF itineraries from JSON input
- Support for multi-day itineraries with activities grouped by time of day
- Automatic calculation of trip duration
- Clean, responsive PDF output
- Easy-to-use REST API
- Containerized with Docker for easy deployment

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
   This will:
   - Build the Docker image
   - Start the API on port 8080
   - Create necessary directories

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
  "trip_name": "My Awesome Trip",
  "user_name": "John Doe",
  "start_date": "2025-08-15",
  "end_date": "2025-08-20",
  "days": [
    {
      "day_number": 1,
      "date": "2025-08-15",
      "time_groups": [
        {
          "time_of_day": "Morning",
          "activities": [
            {
              "name": "Flight to Destination",
              "time": "07:00 AM",
              "location": "JFK to LAX"
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

### Project Structure

```
.
├── Dockerfile           # Docker configuration
├── docker-compose.yml   # Docker Compose configuration
├── go.mod              # Go module definition
├── main.go             # Application entry point
├── input/              # Sample input JSON files
├── output/             # Generated PDFs (created at runtime)
└── templates/          # HTML templates for PDF generation
```

### Building the Docker Image

```bash
docker build -t itinerary-api .
```

### Running Tests

```bash
go test ./...
```

## Troubleshooting

### Common Issues

1. **Port 8080 is already in use**
   - Change the port mapping in `docker-compose.yml` or use `-p` flag:
     ```bash
     docker run -p 8081:8080 -d itinerary-api
     ```

2. **Permission denied when writing to output directory**
   - Ensure the output directory has write permissions:
     ```bash
     chmod -R 777 output/
     ```

3. **Container fails to start**
   - Check logs:
     ```bash
     docker logs <container_id>
     ```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request
