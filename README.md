# lynx
Lynx implements a smart bookmark service inspired by [bunny1](http://www.bunny1.org).

It generally works through commands. The most common commands are searches
(Google, Wikipedia) that return a list of results, and redirectors
(Twitter, Zoom) that take you a page related to an ID or name.

Sometimes, Lynx can also auto-detect a relevant command based on your entry.
Examples:

* `@skylar` The twitter page for "skylar"
* `17drmHLZMsCRWz48RchWfrz9Chx1osLe67` Account info for an Bitcoin address.


## How are entries resolved?

1. **Commands** - *twitter drake, mail, zoom 4155551212* This redirects you to a specific URL, or
a custom URL based content that follows the command.
1. **Detection** - *ASH-123, @skylar* Lynx attempts to identify different types of IDs based on patterns.
1. **Search** - If there are no matches lynx will interpret your text as a Google search.

Commands generally also have abbreviated versions:

* `g birds` -> `google birds`
* `z 111` -> `zoom 111`


## How to use it?

* `/` Enter a search in the text field on the main page.
* `/c/{entry}` Redirect directly to the result for `entry`. This is useful if you want to setup Lynx as your default search engine.


## Credits
* Product inspiration by [bunny1](https://github.com/ccheever/bunny1).
* Backend and URLStore adapted from the [iris URL Shortner example](https://github.com/iris-contrib/examples/tree/master/url-shortener).
* Frontend adapted from [tinyurl](https://github.com/tinyurl/tinyurl).
