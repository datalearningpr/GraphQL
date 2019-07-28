import graphene
from models import Book, Author


class AuthorType(graphene.ObjectType):
    id = graphene.ID()
    name = graphene.String()
    age = graphene.Int()
    books = graphene.List(lambda: BookType)

    def resolve_books(self, info):
        param = str(self.id)
        return list(Book.objects(authorId=param))


class BookType(graphene.ObjectType):
    id = graphene.ID()
    name = graphene.String()
    genre = graphene.String()
    author = graphene.Field(AuthorType)

    def resolve_author(self, info):
        param = str(self.authorId)
        result = Author.objects(pk=param)
        return result[0]


class Query(graphene.ObjectType):
    book = graphene.Field(BookType, id=graphene.ID(required=True))
    books = graphene.List(BookType)

    author = graphene.Field(AuthorType, id=graphene.ID(required=True))
    authors = graphene.List(AuthorType)

    def resolve_book(self, info, id):
        result = Book.objects(pk=id)
        return result[0]

    def resolve_books(self, info):
        return list(Book.objects.all())

    def resolve_author(self, info, id):
        result = Author.objects(pk=id)
        return result[0]

    def resolve_authors(self, info):
        return list(Author.objects.all())


class AddAuthor(graphene.Mutation):
    class Arguments:
        name = graphene.String(required=True)
        age = graphene.Int(required=True)

    id = graphene.ID()
    name = graphene.String()
    age = graphene.Int()
    books = graphene.List(lambda: BookType)

    def resolve_books(self, info):
        param = str(self.id)
        return list(Book.objects(authorId=param))

    def mutate(self, info, name, age):
        result = Author(name=name, age=age)
        result.save()
        return result


class AddBook(graphene.Mutation):
    class Arguments:
        name = graphene.String(required=True)
        genre = graphene.String(required=True)
        authorId = graphene.ID(required=True)

    id = graphene.ID()
    name = graphene.String()
    genre = graphene.String()
    author = graphene.Field(AuthorType)

    def resolve_author(self, info):
        param = str(self.authorId)
        result = Author.objects(pk=param)
        return result[0]

    def mutate(self, info, name, genre, authorId):
        result = Book(name=name, genre=genre, authorId=authorId)
        result.save()
        return result


class Mutation(graphene.ObjectType):
    add_author = AddAuthor.Field()
    add_book = AddBook.Field()
