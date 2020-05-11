TCP Server (S) listen for a connections on a port
TCP Client (C) connects to a TCP server on a port 
C reads STDIN for any input. On hit enter (‘\n’) C sends input it got to a server
S reads input (split by ‘\n’) and checks if it’s an int, returns input multiplied by 2
If it’s not an integer, return uppercased input string.
C shows response from S to STDOUT and waits for another input from STDIN.
C exits by input `exit`
S exits by ctrl+C
