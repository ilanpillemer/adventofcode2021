rawImgEnh = "10100001111111100000001100110101111100010010011011101101001100000101000010001111010001011010101000011100010110100010100001000101001000011101100000101111110101101001111001000101101100010000111001011001110010010000010111011000011000100001110100000001011111000101100111000011010110011011111100011001101001110111000001110001101101001000110010111001001110011000100010100100100110000111000000000001110100001111111000111110100110001001000011011010111000111100010100100010110101001011101011111010110111011001001000111010"

img = [
    "0101011011111110010010010101000111011101100110011011100101000110011000011000010011110101110010101001",
    "1001110110000000001001010001000000111001101011110000100100111010111111100011100101011010011010100000",
    "0111001101100101011110101100110101101001111101001000010000100001000101100111010111100011001101101000",
    "1100001100001010100010011110010010101011010101100110100100010011000110001010110000011011110000100000",
    "1000000100001011101100010100101010110011100110011101100110001111001010100110010010101000111000000011",
    "1101100011101011001111111001011101110110000010100011100111110001011101011111110110000101011100110010",
    "0010111000010101001100101110101101110100110100010001000000111110000001110110010011000010000100010001",
    "1011011000101010100110010101101000101110010111101101010011111000111100110011001011110111001111010011",
    "1111100010111100001100001100000000010111010001001110100000000101110100010101111111111100111110011000",
    "0010111100110110100101111101101010100111111001111010101011000101010100110011101110001000001011100011",
    "0010000010101001100111101110110101101010001011111010100001101010000110101101011011010100010000011011",
    "1001100001100011000111111011110110000100101011101011100101110000001101110000111111011010001100011100",
    "1100100001110110010100111111110000001100101011000101101001001011011101000100001011010011111010010101",
    "0010010011111100101100000010000101010010110110011000100001110110010011101100110000001100101001110110",
    "0000101111011011111011101100101111101110000010100111111111011011001110110101001000001111101110101001",
    "0001001010101010111111100010101111011001100000110011111001001011010011000101101100100000000001101011",
    "0001011101101111000110111011111011110000011010110001001001110111010000000001110111110000010110010100",
    "0110101010000010101101010000101110011110001000011010111110100011000111100100011111110001111001000011",
    "1101011101001110011100011110101110111100010010110000000100010100000000100100110000010110001110100100",
    "1101010010000000011001111110111111101111110011010011010000001011101001001100111000000000101111111010",
    "1011101110001010110111111111000010010111110101101100100101000001101001110100001110111000101000110010",
    "1011011110011000011000101000101101001101110011001111101000101000000010010111000010010011000001011100",
    "1011101010100011001010110001110010110000000010110111101001111010101110110000101101111110111101110000",
    "0111100010010010100111100110010100000110100110001010101000110000100000011000100101110100011001011110",
    "0011110001010000111001010010000110010010100111101011001101110001100100111111101011111110000001111000",
    "1010010011101011011100111011010111011111010010111110101101001011111111001110010001011100110010100111",
    "1111110001000010000100000010000001011011111100010000000100100110010011010010011110011000011011001010",
    "1000011101011111101110011010011100100100111010000010101100111001010011011101001001010100110110101111",
    "1001001111101100011010011111110101010011101011011100100001111000111100110011001000100010010010100100",
    "0001000011000111101011010101000000010011000011101100100011000110101010101110011100110010000011000011",
    "1010111111100111000111011010010011101111101110101001100010110011110001100101100010110111110110000000",
    "0100010000001100101011000111011010110000111101101111101101001010011000010101001111110000101110111110",
    "1001011001110010110100001100101001001100110001000111100000000001000100011101010010001001110110101100",
    "0001000101000100101011101010101001100110111111100010111111000010101110101011010101110111011011100111",
    "1011000010100000100111001111101001001101110000101110000101111101110101111110000010010100111001000001",
    "0000100100101000001111100000111000001101101001100001000100011100010011001001111100110101001010001111",
    "1011111000101100100110010110100110000000010001000010100001110101000110110111110011100110101010101111",
    "1011101111010001001111111111001011100101011100010110011111011100100111000011011000100111011010010110",
    "1101111110111010001100111010101011101011000011110010101110100010011011111000110011101001100111100000",
    "1010000011000101110101110101101011100001011000001100001111011110110100010000010011111001010000111011",
    "1110011010000111101111101010010010000111011101001111010111100111010100011101100000000101111101011001",
    "1101011101110000000111010110100101101111011100001001010100011000011001000011101111100111000110101100",
    "0100111010001000100100110001101000111111100010010110011110011000110110110111011110001100000100010111",
    "1011000000000110010011100100001011100110010110000011011100000110010101111111000100010000010011001001",
    "1100101110100010110001010011101100010001111101111000111010011011001011010100111110100011001111000100",
    "0010100000000011111011001110101010010110001101001101001111011010111010101110110001011000010111111101",
    "1110010010101000101110111110101011111100000110101101010001001010110011101001110111100011001111000000",
    "1001001010000111010110010101010110010011000000010010100000001100011110101001110000000000010001010101",
    "1000011001100101100011101001011111010101110101001111110000010111110000101001111001011011111001000100",
    "0110111001100000011100000101011000100111101001000100111011110100111101000101111110101001111100000110",
    "1000111011100011011111001110111011111011101001110100111111100100111110000100110001111110001101000011",
    "1110101010111110011111010011010111000111101010100101000000100010000100010100100100010001100100110011",
    "0111100001100110111001011111100010100101100100110111010110100011101010100001011110101000011011111001",
    "1100110001110101011111001011001111101111011010010011110010011011001110011100000101010111110001011101",
    "1110111101100001011010000100111001101100100000101110000101111110001100000101101100111011110101110101",
    "1010000101010101001101101100110011001111010101001101111100111110001010011000011111110111001100111111",
    "0100000101001111011100111000111011000010101010011011000110010111111110010100010010100101101010111010",
    "1001011110011110110011011001001101111001001100010110010010111010000111110011011111110101110010100001",
    "0110100001111000010110111001001001110101011010011011100001111111011110010111111110100000110100011001",
    "1011100101000000000101010000111111110111100111000000110010110101001111110101110010001111000110110100",
    "1011101001110001101000000110010011010011100101111000101001000101111100110111001101101010001010001011",
    "1101110010010101101101100011100100101010010001110010100010011111011110001111111101000101010111011000",
    "0111110010010011110100100000110110000111000110111101011110100111110010100101101110100110011100100100",
    "0101000010111101011101011101111010110100000011000100010011101110100110001010100110110111110110111001",
    "0110011101100111110110100001110001100001100111011011111111010000010011001000101010111001000101010001",
    "1011110001111011101010111011101111101000111011110010010011110111001010010110111001000000010100110100",
    "0001011100001000111010111010101000101100000010011100111010101001110110010101001001100101001110110101",
    "1001000010011110000111100010011111101010011010110111010100110100000111001000011000000110111100101001",
    "1100110000001001100100110110110010010100011110101100011010110100100011001101100100111100100100100110",
    "1011110111111100000110100110000100101000000111001101100110010101010011001010000010000001010101000010",
    "0000110111000100011100100101100010101011000110000111100100110010110110001100110000100011101011000010",
    "0010011101011101010100011110001110100101010001010000000110010000110010110110011001001111010100000001",
    "1101101101000010100111011110011001001111110010111011101111100000010100010100000111110001000010111000",
    "0100011011010010110011100101111110001001000010100001001110000101010110001111101001000100100111000000",
    "0100000011110000101000101011011101001001111010001111001111101101000101010110110001010101101011000111",
    "0011100111110100001100100101110001001100110100000010011110101110100010000111010110011110101111011100",
    "0011001010101011011110001001001001011001000110101101011101110101011000101111111010100000111100000000",
    "1101110101010010011001101101101110001000110010110000010100101101010111000110111100110011001010000100",
    "0010110101101110001110010010001100001001000000111101100011111000100010111101101110000010010100110110",
    "1001111110101001100011010100110001101111011111110010010011101100100100010000100101100001110001010110",
    "0010101000011110101000000000100101001100111100010111100010010100011101111001100000010010001011011010",
    "0111101100110111100101011001111010100000011011000111101111010010001101001111110101101011011010101101",
    "0100001100010011100101001011001011110100111001100111110010010101100100000100111000001110101101011010",
    "0001100010001001000000000011111110000001101101011011010010110111011100100100001100101110100100011111",
    "0111110001011000001111011101010010000100011001000011101011011111111110101000101111000000010010001010",
    "0100001011110110101110010111010000101110011000111000100000111001111000100110100110100001110100001010",
    "0110011100000000101011001101101100100111101010110110100101001111111000001111110101001001110000101101",
    "0010110010001101111110010011011100010101111001111001010001000001100011010011010010010111011001010000",
    "0010111011101001111100001000100100010010110011001011110010010110101100111111100011010101001011001110",
    "1110110010010001111101000010100010011100110001100010001000101110100100111111111001100000100001010001",
    "1001101000010011110011101010101101110100101110110000110101001001010100111001101010100100101011100111",
    "1010101110010101000001110101000010111110000110111011000010000000100111010100111001111010000011110000",
    "1111101111010100111110101000010000100000011110111101011101100100010100110110111111111011101010101011",
    "0000010110110111110001011000100011111111110111000011001111000100111101100111010111111101100011011100",
    "1010011001001100000000000111101000011011010010110011110100101001000100011101110100011010101101001110",
    "0111110101011111110110010010010011101110100111100101100001010010110110110111111111000000111100101101",
    "0100001110010111011010001010001000101111000100000001001110011110111100101111011010110001110111000010",
    "1100000011000001111101010100000010101010111011111100000101101001011111110111011111111001111001011101",
    "1110110000111001011111001001111001000000011011010110101010100100000011110011100111101101000000010101",
    "1001110111001001101111101101000011011101110000111111101000101010100111111100010000010011011110111011",
]