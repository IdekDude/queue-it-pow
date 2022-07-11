import hashlib

def getHash(inputString: str, paddingLen: int, runs: int):
    results = []
    padding = "0" * paddingLen
    currentRuns = 0

    postfix = 0
    while True:
        postfix += 1
        target = (inputString + str(postfix)).encode()
        encoded = hashlib.sha256(target).hexdigest()

        if encoded.startswith(padding):
            results.append({"postfix": postfix, "hash": encoded})
            currentRuns += 1

        if currentRuns == runs:
            break
    
    return results
