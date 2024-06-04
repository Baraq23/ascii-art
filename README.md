## ASCII Art Program in Go!
![Go-Logo_Blue-removebg-preview(2)](https://github.com/makebelief/makebelief/assets/166484145/ad53f422-f338-4dd7-9ef1-ab772aa1fbb5)
### Overview

This project is an ASCII art program implemented in Go. It takes a string as input and outputs a graphic representation of the string using ASCII characters. The output resembles the input string with ASCII characters.
### Content

- functions/
    - test_files/
        - standard.txt
        - start_test.go
    - formatspecialcharacters.go
    - printart.go
    - readbanner.go
    - start.go
    - validatebannerfile.go
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

### How to use the program
- Strings to be printed must be put in quotation marks.
```bash
:~ go run main.go "Hello\n"
```
- The following banner formats are included:
  - `shadow`    -sh
  - `standard`  -st
  - `thinkertoy`-th
 - To use any type the banners above:
    specify its corresponding flag as follows:
    -sh, -st and -th respectively.

 _Examples_
```bash
:~ go run main.go "Welcome" -th

o       o     o                    
|       |     |                    
o   o   o o-o |  o-o o-o o-O-o o-o 
 \ / \ /  |-' | |    | | | | | |-' 
  o   o   o-o o  o-o o-o o o o o-o 
                                   
:~ go run main.go "Golang" -sh 
                                                
  _|_|_|          _|                            
_|         _|_|   _|   _|_|_| _|_|_|     _|_|_| 
_|  _|_| _|    _| _| _|    _| _|    _| _|    _| 
_|    _| _|    _| _| _|    _| _|    _| _|    _| 
  _|_|_|   _|_|   _|   _|_|_| _|    _|   _|_|_| 
                                             _| 
                                         _|_|                           
```

- Standard banner template will be automatically selected if user has not entered any flag.

### Error handling
- The program notifies if the template bannerfile is corrupted or empty. In this case the user should download a new banner template from the following links;
    -  <a href="https://learn.zone01kisumu.ke/git/root/public/src/branch/master/subjects/ascii-art/shadow.txt">shadow</a>
    - <a href="https://learn.zone01kisumu.ke/git/root/public/src/branch/master/subjects/ascii-art/standard.txt">standard</a>
    - <a href="https://learn.zone01kisumu.ke/git/root/public/src/branch/master/subjects/ascii-art/thinkertoy.txt">thinkertoy</a> 

  
### Collaborators


<table>
<tr>
  <td align="center" style="word-wrap: break-word; width: 150.0; height: 150.0">
        <a href=https://www.linkedin.com/in/barrack-kope-064a43244/ >
            <img src=https://learn.zone01kisumu.ke/git/avatars/69ae4e7472c685f60beaf04edb53b624?size=870 width="100;"  style="border-radius:50%;align-items:center;justify-content:center;overflow:hidden;padding-top:10px" alt=Aaron/>
            <br />
            <sub style="font-size:14px"><b><i>Barrack Otieno</i></b></sub>
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

  Push your local branch to the remote repository.

```terminal
git push -u origin <your_branch_name>
```

**8.** Create a Pull Request!

**Review:** At this stage, we will review PR and merge your code into our codebase after approval by our team.

### Issues and Contributions

If you encounter any issues or have suggestions for improvements, feel free to submit an issue or propose a change.

## _____________________

Feel free to explore the codebase, run the program with different inputs, and contribute to enhancing the project.