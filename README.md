# days-count-go
This project shows a basic UI where user can enter the date and time in the format '2020-02-18T01:50' and get the number of days for the entered month. It creates an ajax post call to get the response data.
# What it does
The project:
1. Takes the date input from the user (only valid format is `2020-02-18T01:50`)
2. Once the user clicks the submit button, an ajax post call is made with the json request being `{Date: "2020-02-18T01:50"}`
3. The ajax receives the response with the date and the number of days in the entered month in the format `{Date: "2020-02-18T01:50", LastDayOfMonth:29}`
4. The json data received by the ajax call is displayed on the web UI
`
