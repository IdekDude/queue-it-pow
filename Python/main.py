import hashlib
import base64
import json 

def getHash(input, zeroCount):
    final = []
    zeros = "0" * zeroCount
    postfix = 0
    for x in range(10):
        while True:
            postfix  += 1
            stri = input + str(postfix)
            encodedHash = hashlib.sha256(stri.encode()).hexdigest()
            if encodedHash.startswith(zeros):
                final.append({ "postfix": postfix, "hash": encodedHash })
                break
            
    return base64.b64encode(json.dumps(final).encode("utf-8"))
