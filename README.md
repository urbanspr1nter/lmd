# lmd

A text-to-markdown list generator in `go`!

## What is it?
`lmd` is a tool written by me to:
1. Practice the `go` programming language with a small project
2. Solve a problem in that I need an easy way to manage my task lists and have some way to archive them for future reflections

It just takes any plain text files and saves it into a markdown list. See the [Input Format](#input-format) section in this document
for details how that is represented.

## Build
As of now, there is a prebuilt binary for macOS in the `bin` folder. If you want to build a custom version, you can
follow this section.


You will need:
* Most recent version of the `go` compiler.
* A `bin` folder within the project root. It is recommended that the binary is placed here.
* Computer in UNIX environment to run scripts. This means, macOS, Linux or Windows with WSL.

To build, just navigate to the project root and execute the following command:
```shell
go build -o bin/lmdsh
```

Now, `lmd` will be the binary that lives in the `bin` folder.

## Preparing for Usage
> ⚠️ The script on the project root makes assumptions that the binary is already in the `bin` folder.

The `lmdsh` script is very opinionated. You'll have to edit it yourself to make it work the way you want. 

`lmdsh` looks like this:

```shell
DATE=$(date '+%Y-%m-%d')
./bin/lmd $1 > output/$DATE.md
cp output/$DATE.md ~/code/work-log
cd ~/code/work-log
git add . && git commit -m "$DATE log" && git push
```

1. `lmd` is executed and is assumed to be in the `bin` folder.
2. A file with the filename format `$DATE.md` is output.
3. Once an output markdown file is generated, the output markdown file is copied to a `git` repository within `~/code/work-log`.
4. Go to the `git` repository and commit things.

Notice that you can modify any step in the script. Most likely you'll want to modify the part where the output file is copied to a `git` repo.
Feel free to hack around.

## Usage

Execute the `lmdsh` script with a path to the input text file as an argument:

```shell
./lmdsh /path/to/input.txt
```

## Input Format
Input files are just regular text files with the _assumption_ that the first line is the title of the list.
Every line following the first line is a list item.

So something like this:
```text
my list
wake up
brush teeth
read
breakfast
```

Turns into this:

```text
# my list
* [ ] wake up
* [ ] brush teeth
* [ ] read
* [ ] breakfast
```

## Contact

Reach out to me if you have issues. This was written as a personal project and I just wanted to share so that others may use it
or get some inspiration on their own. 😄

Website: [https://roger.lol](https://roger.lol)

Email: urbanspr1nter@gmail.com