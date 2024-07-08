### How to use
1. Search for `INSERT` to find replacements that need to be made
2. Import to API as repository
   1. create [this](https://www.digitalocean.com/community/tutorials/how-to-use-a-private-go-module-in-your-own-project#providing-private-module-credentials-for-https) .netrc file
   2. run `go env -w GOPRIVATE=github.com/COMPANYNAME` -> COMPANYNAME is case-sensitive & make sure the go import module path is case-sensitive too
   3. after you upload this to your own repository make s ure to change all imports to the correct module path
      1. global replace `"github.com/Infinitare/types-template`
      2. and replace in go.mod