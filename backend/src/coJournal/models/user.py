from pydantic import BaseModel, EmailStr
from typing import List
from coJournal.models.journal import Journal

class UserBase(BaseModel):
    username: str
    displayName: str
    email: EmailStr
    journals: List[Journal] = []

class User(UserBase):
    uid: str
