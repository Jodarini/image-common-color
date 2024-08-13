# Image Color Extraction API

A Go-based serverless API for extracting the most prevalent color from an image. This service processes images retrieved from a provided URL and returns the average RGB values of the most common color.

## Features

- **Serverless Architecture**: Designed to run as a serverless function, making it easy to deploy on platforms like Vercel.
- **Image Processing**: Decodes images in JPEG and PNG formats and computes the average RGB values of the predominant color.
- **HTTP API**: Exposes a simple HTTP endpoint to request color analysis.

## Endpoint

- **GET** `/api/color?url=<image_url>`

  **Query Parameters:**
  - `url` (string): The URL of the image to analyze.

  **Response:**
  ```json
  {
    "red": 123,
    "green": 234,
    "blue": 45
  }
