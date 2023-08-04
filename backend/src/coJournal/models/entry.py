from pydantic import BaseModel, Field
from datetime import datetime
from uuid import uuid4, UUID

class Entry(BaseModel):
    id: UUID = Field(default_factory=uuid4)
    created_at: datetime = Field(default_factory=datetime.utcnow)
    updated_at: datetime = Field(default_factory=datetime.utcnow)
    title: str
    note: str
    beenViewd: bool = False

