# eolutil

I didn't want to futz around trying to get dos2unix to compile on macOS so I wrote a *very* simple EOL utility.

```
Usage of ./eolutil:
  -debug
    	once more, with feeling
  -f string
    	the file on which to act
  -fs string
    	glob pattern for multiple files; note that the pattern should be double quoted or escaped to prevent shell from globbing first
  -help
    	show this message
  -i	modify the file in place instead of echoing to STDOUT
  -rn
    	use \r\n instead of \n
```

It might be more work, but instead of using `-i` you should redirect the correct output to a new file and then rename it. It's more work but fewer args to remember/support and harder to screw up. To that end, I chose to make `-i` very dangerous.
