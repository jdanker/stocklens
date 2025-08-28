"""Utility functions for retrieving stock data."""

import logging
from typing import Any, Dict, Optional

import yfinance as yf

logger = logging.getLogger(__name__)


def _get_last_price(ticker: str) -> Optional[float]:
    """Fetch the last traded price for ``ticker``.

    Returns ``None`` if the price cannot be determined.
    """

    data = yf.Ticker(ticker)
    try:
        price = data.fast_info["last_price"]
    except Exception:  # pragma: no cover - fallback
        price = None
    if price is None:
        try:
            hist = data.history(period="1d")
            price = float(hist["Close"].iloc[-1])
        except Exception:
            price = None
    return price


def get_price(ticker: str) -> Dict[str, Any]:
    """Return a JSON-serialisable dict with ticker price."""
    ticker = ticker.upper()
    price = _get_last_price(ticker)
    if price is None:
        raise ValueError(f"no price for ticker {ticker}")
    return {"ticker": ticker, "price": float(price)}


def get_analysis(ticker: str) -> Dict[str, Any]:
    """Return basic metrics for ``ticker``.

    Currently includes price and trailing P/E ratio if available.
    """

    ticker = ticker.upper()
    data = yf.Ticker(ticker)
    price = _get_last_price(ticker)
    info = data.info or {}
    pe = info.get("trailingPE")
    if price is None:
        raise ValueError(f"no price for ticker {ticker}")
    result: Dict[str, Any] = {"ticker": ticker, "price": float(price)}
    if pe is not None:
        result["pe"] = float(pe)
    return result

