package commands

import (
	"fmt"

	"github.com/openshift-splat-team/jira-bot/cmd/issue"
	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
)

var CreateAttributes = Attributes{
	Regex: `\bjira\s+create\b`,
	RequireMention: true,
	Callback: func(evt *slackevents.MessageEvent, args []string) ([]slack.MsgOption, error) {
		url := GetThreadUrl(evt)
		description := args[4]
		if len(url) > 0 {
			description = fmt.Sprintf("%s\n\ncreated from thread: %s", description, url)
		}
		description = fmt.Sprintf("%s\nissue created by splat-bot\n", description)

		issue, err := issue.CreateIssue(args[2], args[3], description, args[5])
		if err != nil {
			return WrapErrorToBlock(err, "error creating issue"), nil
		}
		issueKey := issue.Key
		issueURL := fmt.Sprintf("%s/browse/%s", JIRA_BASE_URL, issueKey)
		return StringToBlock(fmt.Sprintf("issue <%s|%s> created", issueURL, issueKey), false), nil
	},
	RequiredArgs: 6,
	HelpMarkdown: "create a Jira issue: `jira create [project] [summary] [description] [type]`",
}
