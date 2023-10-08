# PlaceDoggo

The Placeholder Image API is a simple service written in GoLang that provides user-defined sized pictures of dogs. It generates placeholder images on-the-fly with the requested dimensions, allowing developers to integrate dynamic dog pictures into their applications for testing or prototyping purposes.

## API Endpoints

The API provides a two endpoints for requesting dog pictures:

`get`

- `{width}`: The desired width of a dog picture in pixels.
- `{height}`: The desired height of a dog picture in pixels.

`any`

- `/any`: A random picture of a dog picture.

## Usage

1. Clone the repository:

`git clone https://github.com/RyanRasi/PlaceDoggo`

2. Build and run the API:

`go build`

`./placedoggo`

The API server will start and listen for incoming requests on `http://localhost:8000` by default.

3. Request a dog picture:

### To retrieve a dog picture with specific dimensions:

GET http://localhost:8000/get/{width}x{height}

Replace `{width}` and `{height}` with the desired dimensions of a random dog picture.

Example:

GET http://localhost:8000/get/800x600

This will return a dog picture with a width of 800 pixels and a height of 600 pixels.

### To retrieve a random dog picture:

GET http://localhost:8000/any

This will return a random dog picture of differing dimensions.

## Dependencies

`github.com/gorilla/mux`

`github.com/nfnt/resize`

## Support and Donations

If you find this project useful and would like to support its development, consider making a donation.

[!["Buy Me A Coffee"](https://www.buymeacoffee.com/assets/img/custom_images/orange_img.png)](https://www.buymeacoffee.com/uiSK0Ex)

## License

This project is licensed under the [MIT License](LICENSE).
