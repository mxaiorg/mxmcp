package main

/*
	Tool descriptions

	These mirror the structured email_search variant exposed by the mxmcp2 MCP
	server (internal/mcpsrv/server.go). The calling model fills the search fields
	directly instead of passing a single natural-language sentence, and this client
	forwards them to the backend's advanced-search endpoint.
*/

// kSearchDescription documents the structured email_search tool.
const kSearchDescription = `Search the user's email (messages and attachments). ` +
	`Analyze the user's question and fill in as many of the search fields as ` +
	`apply — the server runs hybrid vector+keyword retrieval, reconstructs threads ` +
	`when needed, and returns results shaped for the request (a count, short email ` +
	`stubs for lists, or full bodies for summaries).`

// Per-field descriptions for the structured search arguments.
const (
	kQueryArgument = "The user's original question or statement, verbatim."

	kSearchTypeArgument = "What kind of result the user wants: 'emails' = a list of " +
		"messages (e.g. 'list', 'show me the emails'); 'summary' = read or " +
		"summarize whole message(s), returns the complete reconstructed email(s); " +
		"'content' = specific points inside messages (problems, objections, " +
		"details), returns only the passages that matched; 'stat' = how many / " +
		"counts; 'seek' = specific messages are referenced but their MsgId is " +
		"unknown."

	kFromArgument = "Sender(s), comma-separated. Use 'USER' for a first-person " +
		"singular sender ('emails I sent'), 'DOMAIN' for first-person plural " +
		"('emails from us'). Only use values that can plausibly substring-match " +
		"an email address or display name."

	kToArgument = "Recipient(s), same rules as From ('USER'/'DOMAIN')."

	kStartDateArgument = "Start of the date range as the user's local wall-clock time, " +
		"'2006-01-02T15:04:05' or just '2006-01-02' (NO timezone offset — the " +
		"server applies the user's timezone). Resolve relative dates ('today', " +
		"'Tuesday', 'last week') against the date reported by the user_time tool. " +
		"If no year is given, use the year from user_time."

	kEndDateArgument = "End of the date range as the user's local wall-clock time " +
		"(same format and rules as StartDate; no offset)."

	kOrderArgument = "'Newest' (latest/most recent) or 'Oldest' (first/earliest)."

	kBiDirectionArgument = "'Yes' for symmetric phrasings — 'between A and B', " +
		"'conversation with X', 'back and forth'. Put one party in From and the " +
		"other in To (never both in one field) and set BiDirection to 'Yes'."

	kWordsArgument = "Search context from the query NOT already captured by other " +
		"fields. If in doubt, leave blank."

	kKArgument = "A specific number of messages requested, as a string."

	kMsgIdArgument = "Comma-separated message ids when specific messages are " +
		"referenced and known."

	kThreadArgument = "'Yes' when the answer needs a complete thread/timeline " +
		"reconstruction or context spanning multiple related emails."

	kMaxTokensArgument = "Optional cap on the response size in tokens. Omit this in " +
		"almost all cases — the server applies a generous default ceiling that " +
		"returns full email bodies and the complete result set. Only set it if you " +
		"genuinely lack context-window room, and even then keep it as large as you " +
		"can afford."
)
