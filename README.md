# SerialToTCPBridgeSimple
A simple serial UART to TCP connection, raw data bridge. Written in Go.

You need to set:
- COM Port & Baud rate
- Destination IP Address & port

##### Not fault tolerant
Needs data integrity checking like CRC32 and retry capability for a real-world application. Perfect for bench testing though.
