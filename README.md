# LiaMovie

LiaMovie is a Golang-based web application for streaming movies. Users can easily browse and watch their favorite films using this application.

## Installation

1. **Clone the Repository**: Clone the LiaMovie repository to your local machine:
   ```bash
   git clone https://github.com/metelgameryt/LiaMovie.git
   ```

2. **Set up Environment Variables**: Open the `.env` file and configure the necessary environment variables such as database connection details.

3. **Install Dependencies**: Install the required dependencies using Go modules:
   ```bash
   go mod tidy
   ```

4. **Run the Application**: Start the LiaMovie server by running:
   ```bash
   go run main.go
   ```

5. **Access the Website**: Open your web browser and navigate to `http://localhost:8080` to access the LiaMovie website.

## Usage

- **Browsing Movies**: Visit the homepage to browse through the available movies.
- **Watching Movies**: Click on a movie thumbnail to start watching the selected film.

## Adding Movies

To add a new movie to LiaMovie, follow these steps:

1. **Add Movie File**: Place the movie file (in MP4 format) in the designated directory defined in `main.go` as `MoviePath`.
   
2. **Add Thumbnail**: Place the corresponding movie thumbnail (in JPG format) in the directory defined in `main.go` as `ThumbnailPath`. Ensure that the thumbnail file has the same name (excluding the extension) as the movie file.

3. **File Naming Convention**: It's important to name the movie file appropriately. Any underscores (_) or dashes (-) in the filename will be replaced with spaces to generate the movie's title.

4. **Automatic Refresh**: Once the movie file and thumbnail are added to the specified directories, the movie will automatically appear on the LiaMovie website without the need for restarting the application.

## Todos

- Implement user authentication and authorization.
- Add support for TV series.
- Enhance the user interface with additional features like search functionality and movie categories.

[Old Version of LiaMovie | MetelPlus](https://github.com/MetelGamerYT/MetelPlus)
