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
heroku config:add username="maro"
heroku config:add channel="#randam"
heroku config:add iconemoji=":maro:"
```

# Use case

![maro](http://i.gyazo.com/5a10199a515261142662e559d610bd5b.png)
