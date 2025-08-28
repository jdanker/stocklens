import logging

from fastapi import FastAPI, Query, HTTPException

from fetch_data import get_price, get_analysis

logging.basicConfig(level=logging.INFO)
logger = logging.getLogger(__name__)

app = FastAPI()


@app.get("/price")
def ticker_price(ticker: str = Query(..., min_length=1, max_length=5)):
    """Return the latest price for a ticker."""
    try:
        return get_price(ticker)
    except Exception as exc:  # pragma: no cover - logging only
        logger.error("Failed to fetch price for %s: %s", ticker, exc)
        raise HTTPException(status_code=404, detail="Ticker not found")


@app.get("/analyze")
def analyze(ticker: str = Query(..., min_length=1, max_length=5)):
    """Return a basic analysis for a ticker."""
    try:
        return get_analysis(ticker)
    except Exception as exc:  # pragma: no cover - logging only
        logger.error("Failed to analyze %s: %s", ticker, exc)
        raise HTTPException(status_code=404, detail="Ticker not found")

