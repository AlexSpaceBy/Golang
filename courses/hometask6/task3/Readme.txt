Web server, responds to “/” (404 on any other request)
Support “POST” and “GET” methods
On GET request returns HTML page, containing FORM with name and address fields + submit button. 
Also there is a placeholder to display token on this page.
On POST request reads name+address from body and creates a token as “name:address”. 
Then saves this token to Cookies
Page reloads, token value displayed on page(read it from cookie)

The server uses unique names for each user in a database.


Note: you can use simple HTML template for this task: https://jsfiddle.net/7vhun1tc/ or create your own, it’s up to you.This task (as course as well) is about golang, so it’s cool to have some practice in HTML/JS, but it’s not a requirement.
