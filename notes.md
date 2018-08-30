# data structure
``` json
items: [
  {
    name: "",
      // github: =
      // reddit: data.subreddit_name_prefixed + data.title
    description: "",
      // github: =
      // reddit: data.selftext
    url: "",
      // github: html_url
      // reddit: data.url
    created_at: "",
      // github: =
      // reddit: data.created_utc
    updated_at: "",
      // github: =
      // reddit: nil (or last comment post time?)
    source: "github", //etc
    gravity: 139432
       // github: stargazers_count + watchers_count
       // reddit: ups + num_comments
  },
]
```

# items access:
how to access each post/article/item
- github `.items`
- reddit `data.children