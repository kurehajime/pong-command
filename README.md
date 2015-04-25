# pong-command

![pong](https://cloud.githubusercontent.com/assets/4569916/7273449/e6c410be-e92e-11e4-89dd-ba6903089706.gif)

Pong-command is a CUI game.

***PO***NG is ***N***ot PIN***G***.

## How to use.

### 1. Download

 [Download here. Windows,OSX,Linux binary](https://github.com/kurehajime/pong-command/releases/tag/0.1)

### 2. Move file or Add directory to PATH

***Mac OSX / Linux :***

Copy a file into /usr/local/bin 

```
cp ./pong/<OS>/pong /usr/local/bin/pong
```

***Windows :***

[Add directory to PATH Environment Variable](http://www.nextofwindows.com/how-to-addedit-environment-variables-in-windows-7/)

### 3. pong 

`$./pong <IP Address>`

result:

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

## Build yourself

```sh

git clone https://github.com/kurehajime/pong-command.git
go get -u github.com/nsf/termbox-go
go build pong-command/pong/pong.go

```
