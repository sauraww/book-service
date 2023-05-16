**Golang Ports and Adapters Project**


```This project follows the ports and adapters architecture, also known as the hexagonal architecture. It separates the application into layers, allowing for flexibility, testability, and maintainability.```

**#Architecture Overview**

```The project architecture consists of the following layers:```

**Controller Layer**: ```Responsible for handling incoming requests and calling the appropriate app or facade layer. It serves as the entry point for external interactions.```

**App Layer**: ```Contains the application logic and orchestrates the interactions between different parts of the system. It calls the appropriate services and repositories through the defined ports.```

**Middleware Layer**: ```Provides functionalities such as request ID generation, login, or authentication. It sits between the external interactions and the app layer, handling cross-cutting concerns.```

**Logger**: ```Handles logging functionality throughout the application.```

**Ports**: ```Interfaces that define the contract for communication between different layers or external services. There are two types of ports:```

**Incoming Ports**: ```Represent functionalities exposed by the application to the outside world.```

**Outgoing Ports**: ```Represent interactions with external services or systems.
Domain: This layer contains the core business logic and is further divided into the following sub-layers:```

**Model**: ```Defines the data structures and entities used within the application.```

**Ports**: ```Defines interfaces specific to the core business logic, enabling communication between the core and other layers.```

**Application**: ```Implements the use cases and business rules of the application.```

**Infrastructure**: ```Handles external dependencies, such as databases or external services.```

## Running the Project

To run the project locally, follow these steps:

1. Clone the repository:

    ```shell
    git clone (https://github.com/sauraww/book-service.git)
    cd src 
    ```

2. Install the dependencies:

    ```shell
    go mod download
    ```


3. Build and run the project:

    ```shell
    go build
    ./<project_name>
    ```

   Replace `<project_name>` with the actual name of your project.

5. The project should now be running. Access it through the appropriate endpoints or interfaces.
