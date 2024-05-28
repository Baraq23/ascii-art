# ASCII ART PROJECT

The ASCII Art Program reads data from banner file from which word art content is written. Program then takes an input string and prints its corressponding word art in the console.


## Installation & usage

To run the program,follow these steps:

* clone the program into your computer

```bash
https://barraotieno/gitea.com/ascii-art
```

* Open the cloned folder with your prefered text editor among VScode and Vim.

* In the terminal from your text editor type the following:
```bash
go run . "type string here" | cat -e
```

* Hit Enter. The corresponding word art will be printed in the console.

## Output

The following shows input and results when the program is run:

```bash
:~ go run . "Hello" | cat -e

 _    _          _   _          $
| |  | |        | | | |         $
| |__| |   ___  | | | |   ___   $
|  __  |  / _ \ | | | |  / _ \  $
| |  | | |  __/ | | | | | (_) | $
|_|  |_|  \___| |_| |_|  \___/  $
                                $
                                $
:~ go run . "Cold\n" | cat -e
  _____           _       _  $
 / ____|         | |     | | $
| |        ___   | |   __| | $
| |       / _ \  | |  / _` | $
| |____  | (_) | | | | (_| | $
 \_____|  \___/  |_|  \__,_| $
                             $
                             $
$

:~ go run . "Welcome" | cat -e
__          __         _                                    $
\ \        / /        | |                                   $
 \ \  /\  / /    ___  | |   ___    ___    _ __ ___     ___  $
  \ \/  \/ /    / _ \ | |  / __|  / _ \  | '_ ` _ \   / _ \ $
   \  /\  /    |  __/ | | | (__  | (_) | | | | | | | |  __/ $
    \/  \/      \___| |_|  \___|  \___/  |_| |_| |_|  \___| $
                                                            $
                                                            $

```



## Contributions

If you like to contribute to this project,  please create new branch from where you will push your suggestions.

