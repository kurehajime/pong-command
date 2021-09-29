# pong-command

![pong](https://cloud.githubusercontent.com/assets/4569916/7273449/e6c410be-e92e-11e4-89dd-ba6903089706.gif)

Pong-command is a CLI game.

**PO**ng is **N**ot pin**G**.

## How to use.

### 1. Download Win/Mac OSX/Linux binary

 [Download here. Windows,MacOSX,Linux,BSD binary](https://github.com/kurehajime/pong-command/releases)

### 2. Move file or Add directory to PATH

***Mac OSX / Linux :***

Copy a file into /usr/local/bin 

```
cp ./pong /usr/local/bin/pong
```

***Windows :***

[Add directory to PATH Environment Variable](http://www.nextofwindows.com/how-to-addedit-environment-variables-in-windows-7/)

### 3. pong 

Run command.

`$./pong <IP Address>`

Start game.

```

                                                                     006 - 000
--------------------------------------------------------------------------------








                                                                      1
                                                                        9   ||
                                                                          2 ||
  ||                                                                      . ||
  ||                                                                    1   ||
  ||                                                                  6
  ||                                                                8
                                                                  .
                                                                1
                                                              .
                                                            1


--------------------------------------------------------------------------------
EXIT : ESC KEY


```

## ... or install by 'go install'

```sh

go install github.com/kurehajime/pong-command/pong@latest

```

***Caution:***  
wrong:  go get -u github.com/kurehajime/pong-command  
right:  go get -u github.com/kurehajime/pong-command/...  

Do not work? Please try it.

```sh

 GO111MODULE=off go get -u github.com/kurehajime/pong-command/...

```


## LICENSE

This software is released under the MIT License, see LICENSE.txt.

