# healthy_urls
project to recursively check healthy urls

Create a small service to monitor the health of external urls.

The system should accept a list of http/https urls and a crawl_timeout(seconds) and frequency(in seconds) and failure_threshold(count)

crawl_timeout : System will wait for this much time before giving up on the url
frequency : System will wait for this much time before retrying again.
failure_threshold : count of retries possible for that url

The system shall iterate over all the urls in the system and try to do a HTTP GET on the URL(wait for the crawl_timeout) seconds before giving up on the URL. 

Implement a command line/API that takes a file containing the list of URLs to validate with above mentioned parameters and then run through them.



# Description of the code

# Code in src folder

Server folder contains main.go. 
This file contains an init function that initializes the necessary tables. The table scheme is in the Models folder.

The main function impliments cron to recursively run url health checkups every 1 min. Time can be adjusted by changing the '*' values
in the c.AddFunc to represent days or months or years or hrs as needed.

The main function then runs Router() function that is a running server that accepts postman POST and GET requests to accept.

The health checkup function iterated through all urls in database and initiates pingUrl() function for each of them.

The pingUrl() function sends a Get request to all url depending on the parameters in database (crawl_timeout,frequency,failure_threshold)

All ping requests are stored in the table url_hits .

Routes folder contains routes for POST and GET requests. Any clashes in url leads to automtic update of the table entry respective 
to that url's info.

The program supports creating new table entries and automatic updates of existing entries. 

