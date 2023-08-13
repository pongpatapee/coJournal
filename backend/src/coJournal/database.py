import os
import firebase_admin
from firebase_admin import credentials, auth, firestore
# from dotenv import load_dotenv

SERVICE_KEY_PATH = os.path.join(os.path.dirname(__file__), "FIREBASE_SERVICE_KEY.json")

cred = credentials.Certificate(SERVICE_KEY_PATH)
firebase_admin.initialize_app(cred)
db = firestore.client()

