# iBlind Serverless Emailer

Google cloud function setup to be an emailer service that can be used by the iBlind front end.

## Usage

1. Replace "from", "to", "pass", and "subject" variables with appropriate values.
2. Deploy contents of function.go into a google cloud function.
3. Requests with the following json from can be posted to the google cloud function's http link.

```
{
	"name": "NAME",
	"email": "EMAIL",
	"subject": "EMAIL SUBJECT",
	"message": "EMAIL MESSAGE"
}
```

## Testing

-   When deployed into a cloud function Google's built in testing utility can be used. Alternatively a post request can be made to the http link.
-   When testing locally a user can simply run "go run main.go" from the root of this repo.
