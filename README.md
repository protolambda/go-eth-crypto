# go-eth-crypto

Fork of [`go-ethereum/crypto/*.go`](https://github.com/ethereum/go-ethereum/tree/master/crypto), isolated for external usage.

Extracting the crypto utils, with git-history, for reference (warning, affects go-ethereum repo itself):
```sh
cd go-ethereum
git filter-branch -f --prune-empty --index-filter   'git rm --cached -r -q -- . ; git reset -q $GIT_COMMIT -- crypto/crypto.go crypto/crypto_test.go crypto/signature_cgo.go crypto/signature_nocgo.go crypto/signature_test.go' -- --all
cd ..
mkdir go-eth-crypto && cd go-eth-crypto
git init
git pull ../go-ethereum master 
mv crypto/* ./
git add .
git commit -S -m "post-filter: move files"
rm -rf ./crypto
```
Then one small refactor to remove depencencies from go-ethereum common/math/hexutils/more.

## License

LGPL v3, see [`LICENSE`](./LICENSE) file.
