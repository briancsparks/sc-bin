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
cat "$HOME/sc-bin/dots/_functions" >> "$HOME/.functions"
cat "$HOME/sc-bin/dots/_bash_aliases" >> "$HOME/.bash_aliases"
```

Then, put this into your `.profile`:

```shell
# --- BCS Begin ------------------------------
if [ -d "$HOME/sc-bin" ]; then
  PATH="$HOME/sc-bin:$PATH"
  PATH="$HOME/sc-bin/wip:$PATH"
fi

if [ -f "$HOME/opt/venv/bin/activate" ]; then
  . "$HOME/opt/venv/bin/activate"
  echo 'Run `deactivate` to get out of the venv'
fi

if [ -f "$HOME/.functions" ]; then
  . "$HOME/.functions"
fi

if [ -f "$HOME/.bash_aliases" ]; then
  . "$HOME/.bash_aliases"
fi
# --- BCS End ------------------------------
```

Run the one-time scripts:

```shell
~/sc-bin/once/setup-git
~/sc-bin/once/setup-vim
```

## More

Password-less sudo:

```shell
sudo visudo

# %sudo ALL=(ALL) NOPASSWD:ALL
```

