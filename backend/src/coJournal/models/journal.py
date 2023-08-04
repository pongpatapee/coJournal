from pydantic import BaseModel
from typing import List
from datetime import datetime

from coJournal.models.entry import Entry

class Journal(BaseModel):
    # google auth uids 
    personA_uid: str
    personB_uid: str
    dateActivated: datetime
    entries: List[Entry]
    viewingWindow: datetime
