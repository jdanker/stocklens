# StockLens

StockLens is a small multi-language toolkit for retrieving and analysing
stock market data.

It consists of:

* A **FastAPI** microservice that fetches live data from Yahoo Finance.
* A **Go** command line interface built with Cobra that consumes the service.

## Setup

### Python service

```
cd python-service
pip install -r requirements.txt
uvicorn server:app --reload
```

The service listens on `http://localhost:8000` by default.

### Go CLI

```
go run ./cli price AAPL
```

The CLI contacts the Python service using the environment variable
`STOCKLENS_SERVICE_URL` (default `http://localhost:8000`).

## Testing

```
pytest python-service
go test ./...
```

## Future work

* Enhance analysis metrics and incorporate additional data sources.
* Improve error handling and add more extensive tests.

