
; COPY PASTE THIS INTO A NEW AHK SCRIPT

#NoEnv  ; Recommended for performance and compatibility with future AutoHotkey releases.
SendMode Input  ; Recommended for new scripts due to its superior speed and reliability.
SetWorkingDir %A_ScriptDir%  ; Ensures a consistent starting directory.

#MaxThreadsPerHotkey 3

; increase these numbers by a little bit if you're being gg ez'd by the anticheats

ClickWait = 150 ; time to wait from hold to actual click
HoldDelay = 70 ;time from original keypress to the actual press
Timeout = 83 ; this is the timeout in milliseconds, don't mess with it unless hotkeys are broken (make number larger if it's not working)

SetTimer, Loop1, %tick%
SetTimer, Loop2, %tick%
SetTimer, Loop3, %tick%

Loop1:
{
     KeyIsDown := GetKeyState("q")
     if KeyIsDown {
          SendInput, 3
          Sleep, %ClickWait%
          Click
          Sleep, %Timeout%
     }
}
return

Loop2:
{
     KeyIsDown := GetKeyState("f")
     if KeyIsDown {
          SendInput, 7
          Sleep, %ClickWait%
          Click
          Sleep, %Timeout%
     }
}
return
Loop3:
{
     KeyIsDown := GetKeyState("v")
     if KeyIsDown {
          SendInput, 5
          Sleep, %ClickWait%
          Click
          Sleep, %Timeout%
     }
}
return

c::
SendInput, 8
Click
return
v::
return
q::
return
f::
return

^e::
Suspend, Toggle
return

b::
ExitApp
Return