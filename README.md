## ASCII Art Program in Go!
![Go-Logo_Blue-removebg-preview(2)](https://github.com/makebelief/makebelief/assets/166484145/ad53f422-f338-4dd7-9ef1-ab772aa1fbb5)
### Overview

This project is an ASCII art program implemented in Go. It takes a string as input and outputs a graphic representation of the string using ASCII characters. The output resembles the input string with ASCII characters.
### Content

- files/
    - test/
        - standard.txt
        - start_test.go
    - formatspecialcharacters.go
    - printart.go
    - readbanner.go
    - start.go
    - validate.go
- go.mod
- main.go
- readme.md
- shadow.txt
- standard.txt
- thinkertoy.txt
### Features

- This project is written in go lang.
- Handles input with ascii values between ascii values 32 to 126, include special characters suchas newline and tab.
- Uses graphical templates such as standard banner file, shadow banner file and also thinkertoy banner file..
- Includes test files for testing.

### Banner Formats

The following banner formats are included:
- `shadow`
- `standard`
- `thinkertoy`

### Usage

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


### Packages

Only standard Go packages are used in this project.

### Contributor

<table>
<tr>
    <td align="center" style="word-wrap: break-word; width: 150.0; height: 150.0">
        <a href=https://www.linkedin.com/in/barrack-kope-064a43244?lipi=urn%3Ali%3Apage%3Ad_flagship3_profile_view_base_contact_details%3BtM4PTXbQRFml0XDk5FKlrA%3D%3Dlipi=urn%3Ali%3Apage%3Ad_flagship3_profile_view_base_contact_details%3Buf0Ls4oWR2O2WLUMO5sIBg%3D%3D>
            <img src=https://learn.zone01kisumu.ke/git/avatars/69ae4e7472c685f60beaf04edb53b624?size=870 width="100;"  style="border-radius:50%;align-items:center;justify-content:center;overflow:hidden;padding-top:10px" alt=Aaron/>
            <br />
            <sub style="font-size:14px"><b>Barrack Otieno </b></sub>
        </a>
    </td>   
</tr>

</table>

# Goals

- Understanding the Go file system (fs) API.
- Manipulating data in Go.

## HOW TO CONTRIBUTE ? 👷 

**1.** Clone [this](https://github.com/gitea.com/barraotieno/ascii-art) repository.

**2.** Clone the forked repository.

```terminal
git clone https://github.com/011369/starter
```

**3.** Navigate to the project directory.

```terminal
cd ascii-art
```

**4.**  MAKE A NEW FOLDER WITH YOUR PROJECT NAME INSIDE 
<br>

**5.**  Also Add a README file in your project folder which consists of Description/screenshots about your project !
          
 
<br>

**6.** Create a new branch.

```terminal
git checkout -b <your_branch_name>
```

**7.** Add & Commit your changes.

```terminal
  git add .
  git commit -m "<your_commit_message>"
```

**7.** Push your local branch to the remote repository.

```terminal
git push -u origin <your_branch_name>
```

**8.** Create a Pull Request!

**Review** We will review the PR you've made to the codebase of our project.

### Issues and Contributions

If you encounter any issues or have suggestions for improvements, feel free to submit an issue or propose a change.

## Enjoy ASCII Art in Go!

Feel free to explore the codebase, run the program with different inputs, and contribute to enhancing the project.