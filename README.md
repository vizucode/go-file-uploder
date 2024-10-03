# Fiber File Upload API

This is a simple Go-based API built using the **Fiber** framework. The API allows users to upload image files and serve them back for display using an accessible URL. The project includes two key functionalities:
1. Uploading image files to the server.
2. Serving the uploaded image files through a public URL.

## Features

- **Upload Image**: Users can upload image files (JPEG, PNG, etc.) using a `POST` request.
- **Serve Image**: Uploaded image files can be accessed and viewed using a `GET` request via a static URL.

## Installation

### Prerequisites
- Go 1.16+ installed on your system.
- Fiber framework installed (`github.com/gofiber/fiber/v2`).
- cURL or Postman for making API requests (optional but useful for testing).

### Steps

1. Clone the repository or create the project directory:

    ```bash
    git clone https://github.com/vizucode/go-file-uploder
    cd fiber-file-upload-api
    ```

2. Initialize a Go module:

    ```bash
    go mod init fiber-file-upload-api
    ```

3. Install the dependencies:

    ```bash
    go get -u github.com/gofiber/fiber/v2
    ```

4. Create the required folders:
    ```bash
    mkdir uploads
    ```

5. Run the application:

    ```bash
    go run main.go
    ```

    The server will start on `http://localhost:3055`.

## API Endpoints

### 1. **Upload Image**

- **URL**: `/upload`
- **Method**: `POST`
- **Description**: Uploads an image file to the server.
- **Form Data**:
    - `image`: The image file to be uploaded.
- **Response**:
    - On success, it returns the URL of the uploaded file.
  
#### Example Request (using `curl`):
```bash
curl -X POST http://localhost:3055/upload -F "image=@/path/to/your/image.jpg"
```

#### Example Response:
```json
{
  "message": "File uploaded successfully",
  "url": "http://localhost:3055/uploads/1671234567-image.jpg"
}
```

### 2. **Serve Uploaded Image**

- **URL**: `/uploads/{filename}`
- **Method**: `GET`
- **Description**: Fetches the uploaded image and displays it via the browser or API client.
- **Example**: Access the image by opening this URL in your browser:

    ```
    http://localhost:3055/uploads/1671234567-image.jpg
    ```

## Project Structure

```
fiber-file-upload-api/
├── main.go        # Main application file
└── uploads/       # Directory where uploaded images are stored
```

## Example Flow

1. **Upload an Image**:
    - Use `curl` or Postman to upload an image using the `/upload` endpoint.
    - You will receive a response with the image URL, which you can use to access the image.

2. **View the Image**:
    - Use the provided URL (from the `/upload` response) to view the image in your browser or download it using tools like `curl`.

## Running the Project

1. Make sure you have **Go** installed and set up.
2. Clone or download the project.
3. Install the required dependencies.
4. Run the project using `go run main.go`.
5. Use a tool like **Postman** or **cURL** to interact with the API.

## Dependencies

- [Fiber](https://github.com/gofiber/fiber) - Express-inspired web framework built in Go.

## Error Handling

- **400 Bad Request**: If no image is provided or the file format is invalid, the API will return a 400 status with an error message.
- **500 Internal Server Error**: If the server encounters an issue while saving the file, it will return a 500 status.

## License

Golang Fiber Uploader © 2024 by Vizucode is licensed under Creative Commons Attribution 4.0 International. To view a copy of this license, visit https://creativecommons.org/licenses/by/4.0/

---

You can customize the license section and other details as per your project requirements. This **README** should provide a clear guide to installing, running, and using the API.