const crypto = require('crypto')


function getHash(input, paddingLength) {
    const padding = "0".repeat(paddingLength);
    for(let postFix = 0; true; postFix++) {
        const target = input + postFix.toString(),
            encodedHash = crypto.createHash("sha256").update(target).digest("hex");
        if(encodedHash.indexOf(padding) === 0) {
            return btoa(JSON.stringify(Array(10).fill({
                postfix: postFix,
                hash: encodedHash,
            })))
        }
    }
}
