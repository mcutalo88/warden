Warden
====

<img align="right" src="http://vignette2.wikia.nocookie.net/wowwiki/images/a/a4/Warden_artwork.jpg/revision/latest?cb=20091016194237">
[![build status](http://gitlab.azeroth.io/azeroth/warden/badges/dev/build.svg)](http://10.0.1.13/azeroth/warden/commits/dev)

A Discord Mod.

### Build

`go get`
<br/>
`go build`

### Usage

You must authenticate using either an Authentication Token or both Email and
Password for an account.  Keep in mind official Bot accounts only support
authenticating via Token.

```
./warden --help
Usage of ./warden:
  -e string
        Account Email
  -p string
        Account Password
  -t string
        Account Token
```

The below example shows how to start the bot using an Email and Password for
authentication.

```sh
./warden -e EmailHere -p PasswordHere
```

The below example shows how to start the bot using the bot user's token

```sh
./warden -t "Bot YOUR_BOT_TOKEN"
MjI5MDM5Mjk0NTAzODQ1ODk4.Cu7R_g.0EwXEq6MrGEyXjSAbQma4GqRa1I
```
