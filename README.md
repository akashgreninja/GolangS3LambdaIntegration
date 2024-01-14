# AWS-S3EventApi-s-Lambda

This is a complete guide on how to add an  upload and add an event listener to the s3 bucket and redirect it to the lambda function 

Note-I have not attached events from the code as v2 its impossible to find the docs for v2 you can refer the code in decrepeated.go which has v1 code

Steps
1)Create your IAM role and give appropriate permissions i have given admin access for the ease of this you can refer this docs for setting the credentials https://aws.github.io/aws-sdk-go-v2/docs/getting-started/ 
incase you are getting this error 

<code>2024/01/14 16:52:57 operation error S3: ListObjectsV2, resolve auth scheme: resolve endpoint: endpoint rule error, Invalid region: region was not a valid DNS name.
exit status 1 </code>

2)setup your bucket on s3 for the ease i have given it public access as well {I  have deleted it already }
3)setup the code and run it for testing it->it must upload objects to the s3 successfully


#Create a Lambda Function:

Create a Lambda function in the AWS Lambda Console. Make sure to choose the appropriate runtime (Go in your case).
Write or upload the Lambda function code that you want to be executed when an object is uploaded to your S3 bucket.
Configure S3 Event Notification:

Go to the AWS S3 Console.
Find and click on your S3 bucket.
Navigate to the "Properties" tab.
Under the "Event notifications" section, click on "Create event notification."
Configure Event Notification:

Enter a name for your event notification.
Choose the events that you want to trigger the Lambda function (e.g., "All object create events").
Choose the Lambda function you created as the destination.
Configure any other settings as needed.
Save the configuration





#Configuring Lambda Function:
Create Lambda Function:

Go to the AWS Lambda Console.
Click on "Create function."
Choose "Author from scratch."
Enter a name for your function.
Choose "Go" as the runtime.
For the role, create a new role with basic Lambda permissions.
Click on "Create function."
Configure Trigger:

In the Lambda function designer, click on "Add trigger."
Choose "S3" as the trigger.
Select the S3 bucket and choose the event type (e.g., "All object create events").
Configure the other trigger settings if needed.
Click on "Add."


copy code from 
lambda->lambda.go then use 
GOOS=linux GOARCH=amd64 go build -o main main.go
zip function.zip main
convert it to a zip file 

and upload it on lambda now you must be able to use the event driven s3 api's
