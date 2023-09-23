from typing import Union, List
from coJournal.daos.DaoInterface import DaoInterface
from coJournal.database import db
from coJournal.models.user import UserBase, User
from coJournal.database import auth
from datetime import datetime

class UserDao(DaoInterface):
    collection_name = "users"

    def create(self, user: UserBase) -> Union[User, None]:
        # User object created from google auth
        # create an entry in firestore

        auth_user = auth.create_user(
            email=user.email,
            email_verified=False,
            # phone_number=fake.phone_number(),
            password='secretPassword',
            display_name=user.display_name,
            photo_url='http://www.example.com/12345678/photo.png',
            disabled=False
        )

        print(f"Successfully created auth user {auth_user.uid}")

        user = User(uid=auth_user.uid, **user.model_dump())

        doc_ref = db.collection(self.collection_name).document(user.uid)
        doc_ref.set(user.model_dump())

        return self.get(user.uid)


    def update(self, uid: str, user: UserBase) -> Union[User, None]:
        user_data = user.model_dump()
        
        doc_ref = db.collection(self.collection_name).document(uid)

        user_doc = doc_ref.get()
        if (not user_doc.exists):
            return None

        user_data['updated_at'] = datetime.utcnow()

        
        doc_ref.update(user_data)

        return self.get(uid)


    def get(self, uid: str) -> Union[User, None]:
        doc_ref = db.collection(self.collection_name).document(uid)
        doc = doc_ref.get()
        
        if doc.exists:
            return User(**doc.to_dict())
        
        return 

    def get_all(self) -> List[User]:
        user_collection = db.collection(self.collection_name)
        user_list = []
        
        for user_doc in user_collection.stream():
            user_dict = user_doc.to_dict()
            user = User(**user_dict)
            user_list.append(user)
        
        return user_list

    def delete(self, uid):
        db.collection(self.collection_name).document(str(uid)).delete()
        auth.delete_user(uid)

        return f"Deleted user {uid}"


