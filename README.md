# A Golang Facebook Chat Bot To Gather Reviews

## Setup

#### Docker Setup

##### Requirenments:

- Docker
- Create An App For Facebook, Add your Fb page to your app, subscribe to messaging and messaging_feedback, generate token
- Cteate a .env File, and fill it as .env_example
- ngrok to run on post 5000, add the url to webhook post back with your secret token

#### Install

- Run: ngrok http 5000 --region us
- docker compose up -build

## Examples

- Triger Word **Buy**
- Send a message in your app with the word  **Buy** in it.
- The chatbot will ask you for verification
- if you say yes it will send you a Review Templat