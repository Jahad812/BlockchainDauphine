# EX 1 - Function Encoding
First I did it with abi.encodeWithSelector()
The result was 164 bytes.

Then I used abi.encodePacked, and I got 69 bytes.

abi.encodePacked does not align or pad data. Elements are concatenated tightly, with no extra zeros added.


# EX 2 & 3

Please refer to the source codes for these exercises 