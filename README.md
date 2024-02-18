# TODO_CRUD

## Development

This repository provides a development environment that includes a Docker container defined in `dockerfiles/Dockerfile.dev`.

You can choose one of the following methods to start the development environment:

### Method 1: Manual Build and Run

1. Build the application:

   ```sh
   make build
   ```

2. Run the application:

   ```sh
   ./todo_crud
   ```

### Method 2: Docker Compose

1. Start the application using Docker Compose:

   ```sh
   make docker-compose
   # or
   docker-compose up
   ```

This development environment uses [`github.com/cosmtrek/air`](https://github.com/cosmtrek/air) for live reloading, which automatically refreshes your application upon file save.

## Git Hooks

We enforce code standards using [`pre-commit`](https://pre-commit.com/). Follow these steps to set up and use Git hooks:

1. Run the following command to set up Git hooks:

   ```sh
   make git-hooks
   ```

2. To manually execute all Git hooks, use:

   ```sh
   make lint
   ```

3. If you need to add or remove default hooks, update the `.pre-commit-config.yaml` file.

4. To temporarily skip Git hooks, add the `--no-verify` flag to your `git commit` or `git push` command.

5. Customize the `.gitlint` file to define your team's Git message convention.

## Dependencies

```sh
make tools
```

This will install the required tools for the application.

## Configuration and Running

You can configure the microservice in two ways:

### Method 1: Config File

Create a `todo_crud.yaml` file with your configuration values. For example, `server.Config.Host: localhost` in the YAML file corresponds to `server: host: localhost`.

### Method 2: Environment Variables

By default, all environment variables are in UPPERCASE, and follow the same logic as the [**Method 1: Config File**](#method-1-config-file) values.

## Flow Logic

The flow logic for handling requests in a Go service follows Clean Architecture principles and typically involves the following steps:

1. **HTTP Request Handling (Handler Layer):**
    - Incoming HTTP requests are handled by functions located in the `internal/app/handler/` directory.
    - `Handlers` are responsible for handling incoming requests, parsing parameters, input validation/sanitization, and orchestrating the execution of use cases.

2. **Business Logic (Use Case Layer):**
    - Handlers delegate requests to use cases located in the `internal/app/usecase/` directory.
    - `Use cases` contain core business logic, such as orchestrating the execution for data retrieval, validation, and more.

3. **Data Access (Repository Layer):**
    - Use cases interact with repositories located in the `internal/app/repository/` directory to fetch or update data from storage, often a database.

4. **Response Preparation (Use Case Layer):**
    - Use cases prepare response data structures or DTOs that match the API contract or HTTP response requirements.

5. **HTTP Response Handling (Handler Layer):**
    - Handlers receive use case responses and format them into HTTP responses, including setting headers and encoding the response body.

6. **Middleware (Middleware Layer):**
    - Middleware functions in the `internal/middleware/` directory can be applied at various points in the flow for tasks like authentication, authorization, logging, and error handling.

7. **Error Handling:**
    - Error handling is crucial throughout the process. Errors can occur in handlers, use cases, or repositories. Proper error handling and error propagation to the appropriate layer are essential for robust applications.
    - For a correct logging & error handling, the handlers should return a `&fiber.Error{}`

This separation of concerns makes the codebase modular, easier to test, and maintainable. Additionally, it allows for easy component swapping (e.g., changing the database or switching to another protocol) since dependencies are abstracted through interfaces.
