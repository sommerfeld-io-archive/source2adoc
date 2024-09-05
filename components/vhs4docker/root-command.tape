## This tape file represents the settings for the
## link:https://hub.docker.com/r/sommerfeldio/vhs4docker[sommerfeldio/vhs4docker] command recording
## tool. The tool is based on link:https://github.com/charmbracelet/vhs[VHS].
##
## The GIF image is overwritten regularly because the deployment pipeline invokes this tape with
## every workflow run.
##
## * Best Width for Antora Default UI is 792px
## * Best Width for GitHub README is 830px

Output components/vhs4docker/root-command.gif

Set FontSize 14
Set Width 830
Set Height 460
Set LineHeight 1.0
Set Padding 14
Set Margin 2
Set MarginFill "#2E343D"
Set WindowBar Rings

# Type `docker run -v "$(pwd):$(pwd)" -w "$(pwd)" sommerfeldio/source2adoc:latest -s testdata/common -o target/docs`
# Sleep 500ms
# Enter
# Sleep 6s

Type `pwd`
Sleep 500ms
Enter
Sleep 2s

Type `ls`
Sleep 500ms
Enter
Sleep 2s

Type `docker run -v "$(pwd):$(pwd)" -w "$(pwd)" alpine:latest pwd`
Sleep 500ms
Enter
Sleep 2s

Type `docker run -v "$(pwd):$(pwd)" -w "$(pwd)" alpine:latest ls`
Sleep 500ms
Enter
Sleep 6s
