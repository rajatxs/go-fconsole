# Fivemin Console

This project provides interface to manage content on the [Fivemin](https://www.fivemin.in) platform.

> This project is exclusively targeting the Windows platform.

![Fivemin Console](https://res.cloudinary.com/dcwxfpep4/image/upload/v1697043718/screens/rgpebbvpsqm2oop16jv6.webp)

## Requirements

- [Go 1.18](https://go.dev/dl)
- [Node.js 18](https://nodejs.org/en/download)
- [Wails 2.6.0](https://wails.io/docs/gettingstarted/installation)
- [WebView2](https://developer.microsoft.com/en-us/microsoft-edge/webview2)

## Installation

1. Clone the repository:

```shell
git clone https://github.com/rajatxs/go-fconsole.git fconsole
```

2. Change directory to the cloned repository:

```shell
cd fconsole
```

3. Set the following environment variables to your system:

| Variable | Description | Required | Default |
|----------|-------------|----------|---------|
| ```FMC_ENV``` | Platform environment | No | `development` |
| ```FMC_ADMIN_ID``` | Admin account Id | Yes | - |
| ```FMC_CLIENT_URL``` | Client Application URL | Yes | - |
| ```FMC_MONGODB_CONN_URL``` | [MongoDB Connection URL](https://www.mongodb.com) | Yes | - |
| ```FMC_MONGODB_NAME``` | [MongoDB Database Name](https://www.mongodb.com) | Yes | - |
| ```FMC_CLOUDINARY_ID``` | [Cloudinary Public ID](https://cloudinary.com) | Yes | - |
| ```CLOUDINARY_URL``` | [Cloudinary URL](https://cloudinary.com) | Yes | - |
| ```FMC_ALGOLIA_APP_ID``` | [Algolia App ID](https://www.algolia.com) | Yes | - |
| ```FMC_ALGOLIA_API_KEY``` | [Algolia API Key](https://www.algolia.com) | Yes | - |

## Usage

Run the application in development mode:

```shell
wails dev
```

Build the application:

```shell
wails build
```

For more information or inquiries, please contact the project owner: Rajat (rxx256+github@outlook.com)
