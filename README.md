# Object-oriented programming in Python and Go

So today I wanted to compare how we can use object-oriented programming in Golang and Python. 
None of those languages is purely object-oriented. We can make working solutions without it, but it is sometimes easier 
to bundle things together in a class.

Updated version and code examples are accessible in repo: https://github.com/iustitia/challenge-blogpost

Python's support for OOP is deeply rooted. But it doesn't mean that we cannot still use some concepts in Golang.

Let's start with some examples in Python.

    class Book:
        def __init__(self, title, author):
            self.title = title
            self.author = author
        
Here we have simple class to store information about a book. 
Context for it can be in a library or a bookstore. It stores info about title and author. 

Important thing to notice here, is that we don't define class attributes in class definition, 
but rather we set them up in constructor. If we define and set field in class definition, 
those will be shared by all objects of the class. That's rarely what we want from out objects.

Let's create some objects:

    harry = Book("Harry Potter and the Philosopher's Stone", "J.K. Rowling")
    prince = Book("The Little Prince", "Antoine de Saint-Exupéry")

We can update any field:

    harry.author = "Joanne Kathleen Rowling"

Or even add totally new field:

    harry.price = 20

Python doesn't stop us from adding additional fields to object. We don't have to stay restricted to what is defined in a 
constructor. If you would like to restrict it for some memory savings, check out `__slots__`
(https://docs.python.org/3/reference/datamodel.html#slots). 
It's not a standard practice, if memory is not your concern.


## Private and public fields
If you know few things about OOP, you probably heard about encapsulation. According to wikipedia: 
"Encapsulation refers to the bundling of data with the methods that operate on that data, or the restricting of direct 
access to some of an object's components" (https://en.wikipedia.org/wiki/Encapsulation_(computer_programming)).
So not only we want to bundle data together (as we already bundled title and author), we also would like to restrict 
access to some data. Python trusts developers that they know what they are doing. There is no possibility to truly hide 
a field. But we can use a convention to indicate that a field should not be accessed directly.

Let's extend out previous class with some private field:

    class Book:
        def __init__(self, title, author):
            self.title = title
            self.author = author
            self.__is_available = True
    
    harry = Book("Harry Potter and the Philosopher's Stone", "J.K. Rowling")
    harry.__is_available = False

So we can edit "private" attribute of the class. Although any advanced IDE will point out that this is not the right thing to do.

To make it more elegant, we can implement public methods that change the private attribute themselves:

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

    book = Book("Harry Potter and the Philosopher's Stone", "J.K. Rowling")
    book.lend()
    book.return_item()

As you can see class methods are defined as part of class definition. This is opposite of what we have in Go.

## Go

Ok, now let's see how things look like in Go. Go is not object-oriented language per se, but we still can make use of 
existing structures to make use of OOP concepts we already know.

    type Book struct {
        Title string
        Author string
        isAvailable bool
    }

Here we have simple struct that will store author and title information. Go doesn't have constructor method, but convention 
is to create a function of name `New<ClassName>`, so let's add that:

    func NewBook(title string, author string) *Book {
        return &Book{Title: title, Author: author, isAvailable: false}
    }

As you can see, we don't return just an object, but a pointer to an object.
Let's now create an object:

    b := book.NewBook("Harry Potter and the Philosopher's Stone",
                      "J.K. Rowling")

We are restricted here to fields defined in the struct. If we try accessing or setting other attributes like `b.Price`, 
we'll get an error.

Ok, so how about private and public attributes? Go supports it, although it might be hard to notice that from the code itself.
To be as concise as possible, Go doesn't use here any keywords, but pays attention to first letter of the field. 
Wait, what?! Yes, in our example above Title acts as a public field, while isAvailable as a private one. 
Go call public field an exported one and private is unexported one.

So if we try to set b.isAvailable directly, we'll get an error:

    b.isAvailable = false

If we want to implement further methods of our class, we can do so. We need to indicate in the definition that it can be 
invoked only on an object:

    func (b *Book) Return() {
        b.isAvailable = true
    }

Here we say that we want method `Return` to be invoked on pointer to Book - `b *Book`. Of course method name needs to start 
with capital letter as well to be accessible outside of module where it was defined.

## Inheritance

Python supports inheritance that is widely used. Go doesn't support inheritance. If we would like to achieve something 
similar, we would need to use composition and add another struct as part of second structure. 

### Python:
Let's extend previous example with derivative class:

    class LibraryBook(Book):
        def __init__(self, title, author, department):
            super(LibraryBook, self).__init__(title, author)
            self.department = department


### Go

We have to use composition and make Book field of new structure:

    type LibraryBook struct {
        Book *Book
        Department string
    }
    
    func NewLibraryBook(title string, author string, department string) *LibraryBook {
        return &LibraryBook{
            Book: &Book{Title: title, Author: author, isAvailable: false},
            Department: department}
    }

And example usage. All fields from Book struct won't be accessible directly, but have to be accessed through Book object:

    lb := book.NewLibraryBook("Harry Potter and the Philosopher's Stone",
                      "J.K. Rowling",
                      "Children's Library, Oak Street")
    fmt.Printf("New book created. Title: %s Author: %s, Department: %s\n",
               lb.Book.Title, lb.Book.Author, lb.Department)


## Sum up

To sum up Python and Go can help us bundle data together in classes. While Python uses conventions to indicate private 
and public class members, Go will restrict access to any unexported ones.
Python methods will be defined as a part of class definition while Go treats them as separate functions with a 
difference that they can be invoked on objects only.
Inheritance is integral part of Python, while in Go we need to use other ways like composition to achieve similar functionality.

