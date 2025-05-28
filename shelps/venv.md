# Virtualenv

## Background

Due to things... Apparently, it is better (particularaly in Ubuntu) to leave the
system Python alone. Do not upgrade it (so focal stays at 3.8 forever.) Do not install
packages into it.

But then, how do you install system packages? Particularly, how do you install the
`virtualenv` package, which creates virtual environments in the first place?

Use `virtualenv.pyz` to create a one-of-a-kind venv in `~/opt/venv`, which leaves
the binary `virtualenv` at `~/opt/venv/bin/virtualenv`, then create a link to it
from one of your user bin dirs, like `~/.local/bin`.

Then, for any system-wide packages, create a virtual environment at i.e. `~/opt/genv`,
and put `~/opt/genv/bin` into your path.

You can create i.e. `~/opt/genv3-12` for things that you need Python 3.12 for, and put
`~/opt/genv3-12/bin` into your path.

## Usage

For projects, usage is transparent. Just use `virtualenv` as you always have.

```shell
virtualenv -p python3.8 venv
. venv/bin/activate
```

To install global python packages (to put Python-written scripts into your path),
takes a little more effort. You have to activate the genv environment first.

```shell
. ~/opt/genv/bin/activate
pip install whatever
deactivate
```


