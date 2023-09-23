from coJournal.daos.DaoInterface import DaoInterface
from coJournal.models.journal import Journal


class JournalDao(DaoInterface):
    collection_name = "journals"
    subcollection_name = "entries"

    def create(self, journal: Journal):
        pass

    def update(self, id: str, updated_journal: Journal):
        pass

    def get(self, id: str):
        pass

    def get_all(self):
        pass

    def delete(self, id: str):
        pass


