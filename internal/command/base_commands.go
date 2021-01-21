package command

import (
	"net/url"
	"regexp"
	"strings"
)

// BITCOIN
const BlockchainBtcBaseString = "https://www.blockchain.com/btc/address/"

func NewBtcAddressResolver() *Command {
	return NewIdResolverCommand(
		"btc",
		BlockchainBtcBaseString,
		"Get account information and transaction history for a BTC address.",
		[]string{},
		"^(bc1|[13])[a-zA-HJ-NP-Z0-9]{25,39}$",
	)
}

// BLUEJEANS
const BlueJeansBaseString = "https://bluejeans.com/"

func NewBluejeansResolver() *Command {
	return NewIdResolverCommand(
		"bluejeans",
		BlueJeansBaseString,
		"Open a Bluejeans conference for a given meeting ID.",
		[]string{"bj", "bjn"},
		"^[0-9]{10}$",
	)
}

// GOOGLE
const GoogleSearchString = "https://www.google.com/search?q="

func NewGoogleCommand() *Command {
	return NewSearchCommand(
		"google",
		GoogleSearchString,
		"Google Search.",
		[]string{"g"},
	)
}

// JIRA
const JiraBaseString = "https://demo.atlassian.net/browse/"

func NewJiraResolver() *Command {
	return NewIdResolverCommand(
		"jira",
		JiraBaseString,
		"Open a JIRA ticket, given the item ID.",
		[]string{"j"},
		"^[a-zA-Z]{3}[a-zA-Z]*-[0-9]+$",
	)
}

// LIST COMMAND
func NewListCommand() *Command {
	return NewBookmarkCommand(
		"list",
		"/list",
		"Shows the list of available commands.",
	)
}

// TWITTER
const TwitterBaseString = "https://www.twitter.com/"
const TwitterPrefix = "@"

func NewTwitterCommand() *Command {
	cmd := NewIdResolverCommand(
		"twitter",
		TwitterBaseString,
		"Go to a Twitter feed based on handle.",
		[]string{"t"},
		"^@+",
	)
	cmd.handler = func(param string) *url.URL {
		param = strings.TrimPrefix(param, TwitterPrefix)
		url, _ := url.Parse(TwitterBaseString + param)
		return url
	}
	return cmd
}

// WIKIPEDIA
const WikipediaSearchString = "https://en.wikipedia.org/wiki/Special:Search/"

func NewWikipediaCommand() *Command {
	return NewSearchCommand(
		"wikipedia",
		WikipediaSearchString,
		"Search Wikipedia. For exact matches, goes directly to the article.",
		[]string{"w", "wiki"},
	)
}

// XRP
const LivenetBaseString = "https://livenet.xrpl.org/accounts/"

func NewXrpAddressResolver() *Command {
	return NewIdResolverCommand(
		"xrp",
		LivenetBaseString,
		"Get account information and transaction history for an XRP address.",
		[]string{},
		"^r[0-9a-zA-Z]{24,34}$",
	)
}

// ZOOM
const ZoomBaseString = "https://zoom.us/j/"
const ZoomUsernameBaseString = "https://zoom.us/my/"

func NewZoomCommand() *Command {
	nameDetector, _ := regexp.Compile("^[a-zA-Z]([a-zA-Z0-9.]){4,39}$")
	cmd := NewIdResolverCommand(
		"zoom",
		ZoomBaseString,
		"Open a Zoom conference for a given meeting ID.",
		[]string{"z"},
		"^[0-9]{10}$",
	)
	cmd.handler = func(param string) *url.URL {
		base := ZoomBaseString
		if nameDetector.MatchString(param) {
			base = ZoomUsernameBaseString
		}
		url, _ := url.Parse(base + param)
		return url
	}
	return cmd
}
