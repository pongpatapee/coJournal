from typing import Union
from coJournal.database import db
from coJournal.models.user import UserBase, User

class UserDao:
    collection_name = "users"

    def create(self, user: User) -> Union[User, None]:
        # User object created from google auth
        # create an entry in firestore
        doc_ref = db.collection(self.collection_name).document(user.uid)
        doc_ref.set(user.model_dump())

        return self.get(user.uid)

    def update(self):
        pass

    def get(self, uid) -> Union[User, None]:
        doc_ref = db.collection(self.collection_name).document(uid)
        doc = doc_ref.get()
        
        if doc.exists:
            return User(**doc.to_dict())
        
        return 

    def get_all(self):
        pass

    def delete(self):
        pass

