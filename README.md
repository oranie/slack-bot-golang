# to heroku app deploy

```
git push heroku master
```

# heroku app log tail

```
 heroku logs --app [app_name] --tail
```

# config setting

```
example:
heroku config:add webhook_url="https://hooks.slack.com/services/~~~~~"
heroku config:add username="test_post_user"
heroku config:add channel="#randam"
heroku config:add iconemoji=":test_post_user:"
```
