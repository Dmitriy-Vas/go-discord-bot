# go-discord-bot

A console based discord bot built with Go.

### Table of Contents

+ [FAQ](https://github.com/Dmitriy-Vas/go-discord-bot#FAQ)
+ [Install](https://github.com/Dmitriy-Vas/go-discord-bot#Install)
+ [Setup](https://github.com/Dmitriy-Vas/go-discord-bot#Setup)
+ [Run](https://github.com/Dmitriy-Vas/go-discord-bot#Run)
+ [Commands](https://github.com/Dmitriy-Vas/go-discord-bot#Commands)

### FAQ

<details>
<summary>How can I get discord token?</summary>

__User token:__
+ Run your internet browser.
+ Go to the `https://discordapp.com/` site.
+ Open developer tools in your browser (Ctrl^Shift^I).
+ Sign in to your account.
+ Open local storage in dev-tools.
+ Find the "token" scope and copy value.

__Bot token:__
+ Run your internet browser.
+ Go to the `https://discordapp.com/developers/applications/` site.
---
+ __If you already have a bot__:
    + Select your application.
    + Go to the "Bot" page.
    + Below "Username" in the "Token" field, click to the "Copy".
+ __If you don't have a bot__:
    + At the upper right corner click to the "New Application" button.
    + Specify any name to your bot.
    + ^See how to get token if you already have a bot^
---
</details>

<details>
<summary>How to invite my bot to the server?</summary>

+ Run your internet browser.
+ Go to the `https://discordapp.com/developers/applications/` site.
+ Select your application.
+ Below "Name" in the "Client ID" field, click to the "Copy".
+ Put your Client ID instead of "CLIENTID" in this link:
`https://discordapp.com/oauth2/authorize?client_id=CLIENTID&scope=bot`
+ If you want to add a bot without any permissions, then just use link from above and invite bot to your server.
+ If you need permissions, then go to the "Bot" page and scroll to the bottom, then check scopes.
+ Now copy your "Permissions Integer", put your Client ID instead of "CLIENTID" and Permissions Integer instead of "PERMISSIONS" in this link:
`https://discordapp.com/oauth2/authorize?client_id=CLIENTID&scope=bot&permissions=PERMISSIONS`
+ Use this link and invite bot to your server.
</details>

### Install

__Compile bot for yourself with Go:__

+ Download Go from official [site](https://golang.org/)
+ Unpack Go somewhere.
+ Add Go bin folder to your PATH.
+ Clone the repo to your computer:

```
git clone https://github.com/Dmitriy-Vas/go-discord-bot.git
```

After these manipulations, you can start configuring bot.

### Setup

Now that bot is installed, you will need to setup your config.json file. This can be done in few steps:

1. Open the project folder in file explorer.
2. Rename the file config-sample.json to config.json. (Note: Depending on your computer's settings you might not see the .json part of the file name)
3. Change the bot settings with your own settings.

```
{
    // Your discord token to connect bot
    "token": "KuZrjndpA2sjTCuwqGecWUrkXd2ehysFRx6AM8rqYxr56H",
    // true if you use user account, false if bot
    "user": false,
    // Your prefix to use with commands
    "prefix": ">"
}
```

### Run

After setting up the config.json file, bot is ready to go. To run program, simply use the command go run main.go in the console.
If you have setup your config.json properly (and used the correct credentials) you should see an output similar to this

```
Loaded 5 commands
Loaded config
Logged in as Dmitriy
```

### Commands

+ [n-n] *n* is any number. Specify number between this range.
+ [word] *word* is any word. Specify any word instead of this.
+ [user] *user* is any user **mention**. Specify any mention, example: *@Dmitriy#0325*
+ [role] *role* is any role **name**. Specify any role name instead of this.
+ If command has "+" suffix, then you can specify multiple values.
+ Don't forget to add command prefix.

<details>
<summary>Commands list</summary>

+ Ping
    - Responds with "pong"
+ Del [0-100]  
    - Removes specified amount of last messages in the current channel
+ Notice
    - Shows bot's notice
+ Role [user]+ [role]
    - Adds or removes roles from specified users
+ Ban [user] [0-n] [word]+
    - Banhammer's hit. Specified user will lost his soul for specified time.
    - Provide the correct reason instead of [word], nobody likes to lose souls without reason.
</details>
