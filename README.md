# pathtree
Data structure for storing filesystem and API paths for efficient retrieval and visualization

Pathtrees are typically based on ideas from [Trie](ahttps://en.wikipedia.org/wiki/Trie), [Radix Tree](httaps://en.wikipedia.org/wiki/Radix_tree), and [Suffix Tree](https://en.wikipedia.org/wiki/Suffix_tree)

For example, a few of GitHub's APIs can be organized in to the following pathtree:

```
/
|-- /gists
|   |-- /:id
|   |   |-- /forks
|   |   `-- /star
|   |-- /public
|   `-- /starred
|-- /repositories
|-- /user
|   |-- /emails
|   |-- /followers
|   |-- /following
|   |   `-- /:user
|   |-- /issues
|   |-- /keys
|   |   `-- /:id
|   |-- /orgs
|   |-- /repos
|   |-- /starred
|   |   `-- /:owner
|   |       `-- /:repo
|   |-- /subscriptions
|   |   `-- /:owner
|   |       `-- /:repo
|   `-- /teams
`-- /users
    `-- /:user
        |-- /events
        |   |-- /orgs
        |   |   `-- /:org
        |   `-- /public
        |-- /followers
        |-- /following
        |   `-- /:target_user
        |-- /gists
        |-- /keys
        |-- /orgs
        |-- /received_events
        |   `-- public
        |-- /repos
        |-- /starred
        `--/subscriptions
```
