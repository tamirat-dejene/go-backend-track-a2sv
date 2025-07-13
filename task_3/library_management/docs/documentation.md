# Console-Based Library Management System Documentation

## System Description

The Library Management System (LMS) is a console-based application designed to help manage a collection of books and library users. It provides functionalities such as adding new books, registering users, issuing and returning books, and displaying lists of borrowed and available books. At the startup the system accepts names of librarians as arguemnts.

## Modules Description
### Entry point
- **main.go**: Entry point of the application
### Controllers Module
- **library_controllers.go**: Handles user input and coordinates actions between the models and services. It manages the main application flow and user interactions.

### Models Module
- **book.go**: Defines the `Book` struct and related methods for representing and managing book data.
- **member.go**: Defines the `Member` struct and related methods for representing and managing library user data.

### Services Module
- **library_services.go**: Contains the business logic for the LMS, including operations for adding, issuing, and returning books, as well as managing users records.