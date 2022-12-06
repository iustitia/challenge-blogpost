
class Book:
    def __init__(self, title, author):
        self.title = title
        self.author = author
        self.__is_available = True

    def lend(self):
        if self.__is_available:
            self.__is_available = False
            return True
        return False

    def return_item(self):
        self.__is_available = True


class LibraryBook(Book):
    def __init__(self, title, author, department):
        super(LibraryBook, self).__init__(title, author)
        self.department = department


if __name__ == "__main__":
    book = Book("Harry Potter and the Philosopher's Stone", "J.K. Rowling")
    book.lend()
    book.return_item()
    book.price = 20
    print(f"New book created. Title: {book.title} Author: {book.author}. "
          f"Price: {book.price}\n")

    book = Book("The Little Prince", "Antoine de Saint-Exupéry")
    print(f"New book created. Title: {book.title} Author: {book.author}. "
          # f"Price: {book.price}"
          f"\n")
