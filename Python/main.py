import hashlib
import base64
import json


def getHash(input, zeroCount):
    zeros = '0' * zeroCount
    postfix = 0
    while True:
        postfix += 1
        stri = input + str(postfix)
        encodedHash = hashlib.sha256(stri.encode()).hexdigest()
        if encodedHash.startswith(zeros):
            final = 10 * [{'postfix': postfix, 'hash': encodedHash}]
            return base64.b64encode(json.dumps(final).encode('utf-8'))
