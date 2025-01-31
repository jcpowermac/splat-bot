# splat-bot

Basic responder for common questions and problems in the channels that SPLAT helps moderate.

## Building
~~~
./hack/build.sh
~~~

## Running
~~~
export JIRA_PROJECT=SPLAT
export JIRA_BOARD="SPLAT - Scrum Board"
export JIRA_PERSONAL_ACCESS_TOKEN=<your Jira token>
export SLACK_BOT_TOKEN="xoxb-......"
export SLACK_APP_TOKEN="xapp-......"

./slack-bot
~~~

# Adding commands

The bot will receive events for each channel it is in as well DMs with the bot. Commands are invoked by the bot
matching a regex and calling the associated handler. By default, commands are defined in `./pkg/commands`. At a
bare minimum, a command must have a regex and handler:

```go
var HelpAttributes = Attributes{
	Regex: `\bhelp\b`,
	Callback: func(eventsAPIEvent *slackevents.MessageEvent, args []string) ([]slack.MsgOption, error) {
		return []slack.MsgOption{
			slack.MsgOptionText(compileHelp(), true),
		}, nil
	},
}
```
Below are the various attributes which may be applied to a command.
```go
// Attributes define when and how to handle a message
type Attributes struct {
	// Regex when matched, the Callback is invoked.
	Regex          string
	compiledRegex  regexp.Regexp
	// The number of arguments a command must have. var args are not supported.
	RequiredArgs   int
	// Callback function called when the attributes are met
	Callback       Callback
	// Rank: Future - in a situation where multiple regexes match, this allows a priority to be assigned.
	Rank           int64
	// RequireMention when true, @splat-bot must be used to invoke the command.
	RequireMention 		bool
	// HelpMarkdown is markdown that is contributed with the bot shows help.
	HelpMarkdown       string
}
```

# Threads 

If an command is invoked from a thread, any response will be posted back to the thread.

