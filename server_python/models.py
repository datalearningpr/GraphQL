from mongoengine import Document
from mongoengine import fields

class Book(Document):
    meta = {'collection': 'books'}

    name = fields.StringField()
    genre = fields.StringField()
    authorId = fields.StringField()
    v = fields.IntField(db_field='__v')


class Author(Document):
    meta = {'collection': 'authors'}

    name = fields.StringField()
    age = fields.IntField()
    v = fields.IntField(db_field='__v')