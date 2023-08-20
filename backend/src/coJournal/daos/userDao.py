from typing import Union, List
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

    def update(self, uid: str, user: User) -> Union[User, None]:
        user_data = user.model_dump()
        
        doc_ref = db.collection(self.collection_name).document(uid)

        user_doc = doc_ref.get()
        if (not user_doc.exists) or (user_data['uid'] != uid):
            return None
        
        user_doc.update(user_data)

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

        return f"Deleted user {uid}"


