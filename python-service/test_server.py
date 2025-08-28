from fastapi.testclient import TestClient

import server


def test_price_endpoint(monkeypatch):
    def fake_price(ticker: str):
        return {"ticker": ticker, "price": 123.45}

    monkeypatch.setattr(server, "get_price", fake_price)
    client = TestClient(server.app)

    resp = client.get("/price", params={"ticker": "AAPL"})
    assert resp.status_code == 200
    assert resp.json() == {"ticker": "AAPL", "price": 123.45}


def test_analyze_endpoint(monkeypatch):
    def fake_analysis(ticker: str):
        return {"ticker": ticker, "price": 123.45, "pe": 20.0}

    monkeypatch.setattr(server, "get_analysis", fake_analysis)
    client = TestClient(server.app)

    resp = client.get("/analyze", params={"ticker": "AAPL"})
    assert resp.status_code == 200
    assert resp.json()["pe"] == 20.0

