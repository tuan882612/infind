from fastapi import FastAPI
import uvicorn

app = FastAPI()

@app.get('/')
async def start() -> dict:
    return {
        'information':'indeed web scrape service',
        "version": "0.0.1"
    }

if __name__ == '__main__':
    uvicorn.run(
        'main:app', 
        host='0.0.0.0', 
        port=1001, 
        reload=True
    )