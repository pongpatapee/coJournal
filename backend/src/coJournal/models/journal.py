from pydantic import BaseModel, Field
from typing import List
from datetime import datetime

from coJournal.models.entry import Entry

class JournalBase(BaseModel):
    # google auth uids 
    personA_uid: str
    personB_uid: str
    entries: List[Entry]
    date_activated: datetime
    viewing_window: datetime
    created_at: datetime = Field(default_factory=datetime.utcnow)
    updated_at: datetime = Field(default_factory=datetime.utcnow)

class Journal(JournalBase):
    id: str # firestore uid
