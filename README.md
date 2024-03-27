# sc-bin

My scripts

Clone this repo into `~/sc-bin`.

```shell
git clone git@github.com:briancsparks/sc-bin "$HOME/sc-bin"
```

Run these:

```shell
mkdir -p "$HOME/dev"
cp "$HOME/sc-bin/.editorconfig" "$HOME/dev"
cat "$HOME/sc-bin/dots/_functions" "$HOME/.functions"
cat "$HOME/ac-bin/dots/_bash_aliases" "$HOME/.bash_aliases"
```

Then, put this into your `.profile`:

```shell
# --- BCS Begin ------------------------------
if [ -d "$HOME/sc-bin" ]; then
  PATH="$HOME/sc-bin:$PATH"
fi

if [ -f "$HOME/.functions" ]; then
  . "$HOME/.functions"
fi

if [ -f "$HOME/.bash_aliases" ]; then
  . "$HOME/.bash_aliases"
fi
# --- BCS End ------------------------------
```
