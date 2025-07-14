## ABOUT

This is a web application that controls Yeelight smart bulbs in my apartment. It provides two main functionalities:

- toggling bulbs on/off
- adjusting bulb brightness

The application uses HTMX for interactive web elements without full page reloads and is built with:

- Go (Gin framework)
- SQLite for data storage
- HTMX for frontend interactions

## DEVELOPMENT
Prerequisites:
- Go 1.20+
- SQLite3
- Yeelight smart bulbs on the same network  

Project Structure:
```
cmd/main.go             - Main application entry point
internal/               - Core application logic
api/http/               - HTTP responders
bulb/controller/        - Yeelight bulb controller
config/                 - Configuration handling
handler/                - HTTP handlers
repository/sqlite/      - SQLite data access layer
db/sqlite/              - Database schema and seed data
static/                 - Static assets (CSS, JS)
templates/              - HTML templates
config/                 - Configuration files
```

Building the database:
bash
```bash
make build_db # this will create SQLite database file, create necessary tables, insert initial data
```

## RUNNING

application will run with default dev config file (config/dev.json)

```bash
make run # the application will be available at http://localhost:8080
```

## TESTING

The application can be tested by:
- accessing the web interface at http://localhost:8080 and interacting with the toggle and brightness controls
