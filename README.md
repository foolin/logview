# logview

Logview is tool for formatting the color of text output to a terminal. Only support *unix.

# Usage

1.Download [Logview](https://github.com/foolin/logview/releases "Logview release"), or source code and build.

2.Run *unix command, such as: ./logview xxx(shell command like cat/tail)


# Example

Shell command:
`tail -f xxx.log`

![tail -f tigo.log](example/tail.png "Shell command: tail -f tigo.log")



Use logview like this:
`./logview tail -f xxx.log`

![./logview tail -f tigo.log](example/logview.png "Logview command: ./logview tail -f tigo.log")