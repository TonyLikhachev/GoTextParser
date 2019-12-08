# GoTextParser

This programm helps you to create files in every directory you will prefer.
GoTextParser parses every url link in txt file and then on "out" directory you will find html-code of this URLs
in .txt file you have to write sites without "http(s)://www." Name of files will be the same as their URL without full domain name. 

HOW TO RUN THIS PROGRAM

This parser uses "flags" package, so first you have to dowload this on your PC.

Then you have to open your bash, cmd, etc. and write:
go run parser.go -src="a" -out="b"
Where a - directory with your txt example file, e.g. /home/user/in/example.txt
      b - directory, where you can get parsed files e.g. /home/user/out

the result in my example will be files:
duckduckgo.com
golang.org
google.com
radio.yandex.ru
wikipedia.org
with their HTMLs there

P.S. With questions and my errors please write on my e-mail antonlikhachev@protonmail.ch
You are welcome!



