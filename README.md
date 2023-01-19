# A Golang Facebook Chat Bot To Gather Reviews

## Setup

#### Setup

##### Requirenments:
- Golang
- Docker
- Create An App For Facebook, Add your Fb page to your app, subscribe to messaging and messaging_feedback, generate token
- Cteate a .env File, and fill it as .env_example
- ngrok to run on post 5000, add the url to webhook post back with your secret token

#### Install

- Run: ngrok http 5000 --region us
- docker compose up -build or
- go run application.go

## Examples

- Triger Word **Buy a **
- Send a message in your app, with this form:  **Buy** a <your_product>.
- The chatbot will ask you for verification
- if you say yes it will send you a Review Template

## Examples With Images

![Sample Conversation Image](https://i.ibb.co/B3vvmk8/Screenshot.png)
![Sample Conversation Image](https://i.ibb.co/kQ7S7rc/Screenshot-1.png)
![Sample Conversation Image](https://i.ibb.co/Rzhhnss/Screenshot-2.png)
![Sample Conversation Image](https://i.ibb.co/pwGGgJt/Screenshot-3.png)

## Explanation Of the Code

### Entities

#### Customer

##### Fields

- ID: each customer has a unique ID
- First_Name: The first name of the customer
- Last_Name: The last name of the customer
- Facebook_id: The facebook ID of the customer
- Language: The language that the customer is using
- CreatedAt: Time when the customer was created
- UpdatedAt: Time when the customer was last updated

##### Functions

Basic functions to *CREATE*, *UPDATE*, *DELETE* a customer. Also we have implemented three get functions:
- Get by ID
- Get All
- Get by language

#### Conversation

#### Fields

- ID: each conversation has a unique ID
- Facebook_id: The facebook ID of the customer
- Stage: The stage of the conversation
- CreatedAt: Time when the customer was created
- UpdatedAt: Time when the customer was last updated

##### Functions

Basic functions to *CREATE*, *UPDATE*, *DELETE* a conversation. Also we have implemented three get functions:
- Get by ID
- Get All
- Get by facebook id

##### Usage of the Conversation Entity

Each customer has a unique conversation. The satge of the conversation guides us to the type of message we will send. There are three  stages:

- *None*: We wait for the trigger word.
- *Buy*: We heard the trigger word and asked for verification.
- *Review*: We get Yes in the verification, we send the review and we wait for the user to send it back to us.

#### Review

#### Fields

- ID: each review has a unique ID
- Customer_id: The facebook ID of the customer
- Text: The text of the review
- Score: The score of the review
- CreatedAt: Time when the customer was created
- UpdatedAt: Time when the customer was last updated

##### Functions

Basic functions to *CREATE*, *UPDATE*, *DELETE* a conversation. Also we have implemented three get functions:
- Get by ID
- Get All
- Get by facebook id

#### Template

This is the review template.

#### Fields

- ID: each template has a unique ID
- Placeholder: The placeholder of the template
- Title: The title of the template
- Language: The language used to create the template
- Subtitle: The subtitle of the template
- Button_Title: The title of the button that the template have.
- CreatedAt: Time when the customer was created
- UpdatedAt: Time when the customer was last updated

##### Functions

Basic functions to *CREATE*, *UPDATE*, *DELETE* a conversation. Also we have implemented three get functions:
- Get by ID
- Get All
- Get by language

##### Example Template

![Example Template](https://i.ibb.co/ysQM8hR/Screenshot-4.png)

#### Messages

This entity is used to read Facebook Messages and send messages.

##### Example of Facebook Get Message

![Facebook Get Message Example](https://i.ibb.co/jgyRjF8/Screenshot-7.png)

