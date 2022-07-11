const crypto = require('crypto')


function getHash(input, paddingLength, runs) {
    let final = [];
    const padding = "0".repeat(paddingLength);
    let currentRuns = 0;
    for (let postFix = 0; true; postFix++) {
        const target = input + postFix.toString(),
            encodedHash = crypto.createHash("sha256").update(target).digest("hex");
        if (encodedHash.indexOf(padding) === 0) {
            final.push({
                postfix: postFix,
                hash: encodedHash,
            });
            currentRuns++;
        }
        if (currentRuns === runs) {
            return final;
        }
    }
}
