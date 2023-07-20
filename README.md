How to Run the Golang REST API with Fiber
To run the Golang REST API with Fiber, follow the steps below:

Prerequisites
Before you begin, ensure that you have the following installed on your system:

Golang (Go) - You can download and install it from the official Golang website: https://golang.org/dl/
Clone the Repository
First, clone the repository containing the Golang REST API code:

bash
Copy code
git clone https://github.com/dalmasjunior/passLocker
cd passLocker/
Install Dependencies
Next, navigate to the project directory and install the required dependencies:

bash
Copy code
go mod download
Run the API
Now, you can run the Golang REST API using the following command:

bash
Copy code
go run main.go
The API will start and will be accessible at http://localhost:8000.

API Endpoints
The following API endpoints are available:

GET /password-cards: Get all cards.
POST /password-cards: Create a new card.
GET /password-cards/:cardId: Get a specific card by ID.
PUT /password-cards/:cardId: Update a specific card by ID.
DELETE /password-cards/:cardId: Delete a specific card by ID.
Please note that these routes are prefixed with /password-cards and have corresponding functionalities implemented based on the cardsHandler in the SetupRoutes function.

You can interact with the API using tools like curl, Postman, or any other API client.

That's it! You've successfully set up and run the Golang REST API using Fiber.

If you encounter any issues or errors, make sure to check the Golang and Fiber documentation or consult the project's repository for troubleshooting information. Happy coding!