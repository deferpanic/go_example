# go_example
Running Go as a Unikernel

In this example we embed everything including static assets and that let's us remove the need for attaching a second volume like we do for the interpreted languages.

If have a more web facing projects vs an API project it might be wise to attach the 2nd volume linked to in an upcoming example.

To generate a new image:
```
 ~/go/src/github.com/cratonica/trayhost/make_icon.sh gopher.gif
```
