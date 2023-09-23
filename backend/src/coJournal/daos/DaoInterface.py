from abc import ABC, abstractmethod
from pydantic import BaseModel

class DaoInterface(ABC):
    
    @abstractmethod
    def create(self, model: BaseModel):
        pass

    @abstractmethod
    def update(self, id: str, updated_model: BaseModel):
        pass

    @abstractmethod
    def get(self, id: str):
        pass

    @abstractmethod
    def get_all(self):
        pass

    @abstractmethod
    def delete(self, id: str):
        pass

