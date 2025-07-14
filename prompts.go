package main

/*
	Tool descriptions
*/

const kQueryArgument = `An email search query.`

const kSearchDescription = `Search request for emails from multiple email accounts.

Args:
	query (str): 	
		Set the 'query' parameter to a search sentence reflecting the user's request.
		The sentence should be a specific and concise full sentence. Be sure to include
		aspects that relate to the email subject, sender, recipients, dates, etc.
		Use words like, From, To, etc. Include any date ranges if applicable.
		If you need to search for specific emails AND you have the msgIds,
		refer to the notes below.

Returns:

	This function will return details about the search results plus 
	a list of emails that match the user's request - if any.

	Each email entry returned is a JSON object with the following fields:
	- content: the email content
	- subject: the email subject
	- from: the email sender
	- to: the email recipient
	- date: the email date
	- msgId: the email id
	- link: a URL to view the email
	- related: a list of related emails by msgId

Notes:
	- When searching for specific emails and you have the msgIds, indicate and use the msgId in your request or simply provide the links, if you have them.
	- Only assume the user is asking about their own email if they EXPLICITLY IMPLY it in the query.
	- Always share the link to the email with the user.
	- If the user is requesting a "summary" (or similar) of a specific email, include the word "summary" in the query argument sentence.
	- If the user's question is statistical in nature, like, how many emails, or how many emails were sent in a specific month, include the word "statistical" in the query argument sentence.'
	- To search for missing related emails, conduct a search using the msgIds in the "related" field if it is present.`
