# sc-bin

TODO: Do you need `fdd`, `fnn` from gh:setup/delivered/bin?

My scripts

Clone this repo into `~/sc-bin`.

```shell
git clone git@github.com:briancsparks/sc-bin "$HOME/sc-bin"
```

Run the one-time scripts (in this order):

```shell
~/sc-bin/once/setup-basic-system
~/sc-bin/once/setup-dotfiles
~/sc-bin/once/setup-git
~/sc-bin/once/setup-vim
~/sc-bin/once/setup-virtualenv
```

## More

Password-less sudo:

```shell
sudo visudo

# %sudo ALL=(ALL) NOPASSWD:ALL
```

