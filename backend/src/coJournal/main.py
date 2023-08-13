import uvicorn

from fastapi import FastAPI
from fastapi.middleware.cors import CORSMiddleware
from coJournal.routers import users

app = FastAPI()

app.include_router(users.router)

origins = [
    "http://localhost:3000"
]

app.add_middleware(
    CORSMiddleware,
    allow_origins=["*"],
    allow_methods=["*"],
    allow_headers=["*"],
    allow_credentials=True,
)

@app.get("/")
async def root():
    return "Hello, welcome to coJournal"


if __name__ == "__main__":
    uvicorn.run("main:app", host="0.0.0.0", port=8000, reload=True)
