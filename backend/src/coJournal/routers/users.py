from fastapi import APIRouter, Body
from fastapi.responses import JSONResponse
from fastapi.exceptions import HTTPException
from coJournal.database import db, auth
from coJournal.daos.userDao import UserDao
from coJournal.models.user import UserBase, User

from faker import Faker

router = APIRouter(
    prefix="/users",
    tags=["users"]
)

user_dao = UserDao()

@router.post("/")
async def create_user():
# creating user with auth in the backend for now
# will improve with sign in with google from the frontend later
    fake = Faker()

    fake_email = fake.email()

    user_info = UserBase(
        username=fake_email,
        display_name=fake.name(),
        email=fake_email 
    )
    
    user = user_dao.create(user_info)

    return f"User {user.uid} added to the system!"


@router.get("/{uid}")
async def get_user(uid: str):
    user = user_dao.get(uid)

    if not user:
        raise HTTPException(404, detail=f"User {uid} not found") 

    return user

@router.get("/")
async def get_all_users():
    return user_dao.get_all()

@router.put("/{uid}")
async def update_user(uid: str, user: UserBase = Body(...)):
    updated_user = user_dao.update(uid, user)
    if not updated_user:
        raise HTTPException(404, detail=f"User {uid} not found")

    return updated_user

@router.delete("/{uid}")
async def delete_user(uid: str):
    user_dao.delete(uid)

    return f"Deleted user {uid}"

