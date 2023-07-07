def brute_force(message):
    for key in range(1, 27):
        decryption = decrypt_message(message, key)
        print(f"Decryption with key {key}: {decryption}")

def decrypt_message(message, key):
    decryption = ""
    for char in message:
        if 'a' <= char <= 'z':
            decryption += shift_char(char, key, 'a', 26)
        elif 'A' <= char <= 'Z':
            decryption += shift_char(char, key, 'A', 26)
        else:
            decryption += char
    return decryption

def shift_char(char, key, base_char, alphabet_size):
    return chr((ord(char) - ord(base_char) + key) % alphabet_size + ord(base_char))

target_message = "Ugew gnwj zwjw Oslkgf"
brute_force(target_message)
