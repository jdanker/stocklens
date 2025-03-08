from fastapi import FastAPI, Query

app = FastAPI()

@app.get("/price")
def ticker_price(ticker: str = Query(..., min_length=1, max_length=5)):
    return {"ticker": ticker, "price": 12}