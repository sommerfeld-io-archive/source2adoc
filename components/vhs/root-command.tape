## This tape file represents the settings for the link:https://github.com/charmbracelet/vhs[VHS]
## command recording.
##
## The GIF image is overwritten regularly because the deployment pipeline invokes this tape with
## every workflow run.
##
## * Best Width for Antora Default UI is 792px
## * Best Width for GitHub README is 830px

Output root-command.gif

Set FontSize 16
Set Width 820
Set Height 260
Set LineHeight 1.0
Set Padding 14
Set Margin 2
Set MarginFill "#2E343D"
Set WindowBar Rings

Type `docker run -v "$(pwd):$(pwd)" -w "$(pwd)" sommerfeldio/source2adoc:latest --source-dir src --output-dir docs`
Sleep 500ms
Enter
Sleep 6s
