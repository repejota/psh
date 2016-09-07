package main

const (
	// BashEscapeUser Is the current username returned by an escape secuence
	// http://www.tldp.org/HOWTO/Bash-Prompt-HOWTO/bash-prompt-escape-sequences.html
	BashEscapeUser = "\\u"

	// BashEscapeHostname Is the current hostname returned by an escape sequence
	// http://www.tldp.org/HOWTO/Bash-Prompt-HOWTO/bash-prompt-escape-sequences.html
	BashEscapeHostname = "\\H"

	// BashEscapeUserHostname is the User and hostname concatenated with @
	BashEscapeUserHostname = BashEscapeUser + "@" + BashEscapeHostname

	// BashEscapePath Is the current path returned by an escape sequence
	// http://www.tldp.org/HOWTO/Bash-Prompt-HOWTO/bash-prompt-escape-sequences.html
	BashEscapePath = "\\w"

	// BashEscapePrompt Is the character $ for users and # for root
	BashEscapePrompt = "\\$"

	// BashEscapeJobs Is the number of background jobs of this session
	// returned by an escape sequence
	// http://www.tldp.org/HOWTO/Bash-Prompt-HOWTO/bash-prompt-escape-sequences.html
	BashEscapeJobs = "\\j"
)
