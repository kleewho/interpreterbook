Output repl.gif

Require echo

Set Shell "bash"
Set FontSize 24
Set Width 1200
Set Height 800

Type "build/./interpreterbook" Sleep 1s Enter

Sleep 1500ms

Type@100ms "let a = 5;" Sleep 2s Enter
Type@100ms "if (b < a) { let d = 5; }" Enter
Sleep 3s
Type "Ctrl+C"
