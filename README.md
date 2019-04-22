# go-discord-bot

A console based discord bot built with Go.

### Table of Contents

+ [Install](https://github.com/Dmitriy-Vas/go-discord-bot#Install)
+ [Setup](https://github.com/Dmitriy-Vas/go-discord-bot#Setup)
+ [Run](https://github.com/Dmitriy-Vas/go-discord-bot#Run)
+ [Commands](https://github.com/Dmitriy-Vas/go-discord-bot#Commands)

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
Loaded 2 commands
Loaded config
Logged in as Dmitriy
```

### Commands

+ [n-n] *n* is any number. Specify number between this range.
+ [word] *word* is any word. Specify any word instead of this.
+ [user] *user* is any user mention. Specify any mention, example: *@Dmitriy#0325*

<details>
<summary>Commands list</summary>

+ Ping
    - Responds with "pong"
+ Del [0-100]  
    - Removes specified amount of last messages
</details>
