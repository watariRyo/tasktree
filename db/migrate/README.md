## tool install

Migration tool is ridgepole. But you don't need any gem by using whalebrew.

```
brew install whalebrew
```

```
whalebrew install kei2100/docker-ridgepole
```

## command

cd db/migrate
ridgepole --apply -c config.yml -f ./schema/Schemafile
