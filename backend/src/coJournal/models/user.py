from pydantic import BaseModel, EmailStr, Field
from typing import List
from datetime import datetime

class UserBase(BaseModel):
    username: str
    display_name: str
    email: EmailStr
    journals: List[str] = [] # list of journal uids
    created_at: datetime = Field(default_factory=datetime.utcnow)
    updated_at: datetime = Field(default_factory=datetime.utcnow)

class User(UserBase):
    uid: str
