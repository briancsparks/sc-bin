# sc-bin

My scripts

Clone this repo into `~/sc-bin`.

```shell
git clone git@github.com:briancsparks/sc-bin "$HOME/sc-bin"
```

Then, put this into your `.profile`:

```shell
if [ -d "$HOME/sc-bin" ]; then
  PATH="$HOME/sc-bin:$PATH"
fi
```
