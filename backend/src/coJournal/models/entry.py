from pydantic import BaseModel, Field
from datetime import datetime

class EntryBase(BaseModel):
    created_at: datetime = Field(default_factory=datetime.utcnow)
    updated_at: datetime = Field(default_factory=datetime.utcnow)
    title: str
    note: str
    been_viewed: bool = False

class Entry(EntryBase):
    id: str #firestore uid

