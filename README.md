# Simple example GO APi to use with AWS Lambda and API Gateway

To get this running, one need to set up AWS lambda function and upload the code executable as zip file.
Further setting up AWS API Gateway with API Key and proxy lamda configuration.

You can edit the /mainapp/envs/.env file with your configuration data.
The url would be the url API Gateway provides to you, otherwise you can set up custom domain.

Test url looks like this

Gateway url /search/location/cords/-20.267500/148.716949

In Postman, set up Authorisation to type API KEY and the key is x-api-key and the value is provided by AWS Api Gateway and AddToHeader selected.

