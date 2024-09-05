Output components/vhs4docker/root-command.gif

Set FontSize 16
Set Width 1200
Set Height 360
Set Margin 2
Set MarginFill "#2E343D"
Set Padding 18

Set WindowBar Rings

Type "docker run -v $(pwd):$(pwd) -w $(pwd) sommerfeldio/source2adoc:latest -s testdata -o target/docs"
Sleep 500ms
Enter
Sleep 6s
