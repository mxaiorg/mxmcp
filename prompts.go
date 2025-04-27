package main

/*
	Tool descriptions
*/

const kSearchDescription = `Search query for emails from multiple email accounts.

Args:
- query: The user's query or statement related to email. Create a full sentence that best reflects what the user wants regarding emails. If the user is requesting a specific email, refer to tne notes below.

Example:
- query: "Get me the email from Bob about the new product"

Return:
It will return details about the search results plus 
a list of emails that match the query - if any.

Each email entry returned is a JSON object with the following fields:
- content: the email content
- subject: the email subject
- from: the email sender
- to: the email recipient
- date: the email date
- msgId: the email id
- link: a URL to view the email

Notes:
- When the user requests a specific email and you have the msgId, indicate and use the msgId in your request or simply provide the link, if you have it.
- Only assume the user is asking about their own email if they EXPLICITLY IMPLY it in the query.
- Always share the link to the email with the user.`
