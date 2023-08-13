from fastapi import APIRouter
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

userDao = UserDao()

@router.post("/")
async def create_user():
# creating user with auth in the backend for now
# will improve with sign in with google from the frontend later
    fake = Faker()
    auth_user = auth.create_user(
        email=fake.email(),
        email_verified=False,
        # phone_number=fake.phone_number(),
        password='secretPassword',
        display_name=fake.name(),
        photo_url='http://www.example.com/12345678/photo.png',
        disabled=False
    )

    print(f"Successfully created auth user {auth_user.uid}")

    user = userDao.create(User(
        uid=auth_user.uid,
        username=auth_user.email,
        displayName=auth_user.display_name,
        email=auth_user.email,
    ))

    print(f"Successfully added user entry in firestore")
    
    return user


@router.get("/{uid}")
async def get_user(uid: str):
    user = userDao.get(uid)

    if not user:
        raise HTTPException(404, detail=f"User {uid} not found") 

    return user


