from fastapi import FastAPI as server

app = server()

@app.get('/')
def start() -> dict:
    return {'service':'web scrape 1'}